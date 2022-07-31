package controller

import (
	"fmt"
	"net/http"
	"piennews/helper/apiErrors"
	"piennews/helper/config"
	"piennews/helper/jwt"
	"piennews/helper/util"
	"piennews/models"
	"piennews/services"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (ct *controller) Auth(c *gin.Context, provider string) {

	var name, account, avatarURL string

	u, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"code":    http.StatusBadRequest,
			"pienweb": "/",
			"message": err.Error(),
		})
		return
	}

	q := c.Request.URL.Query()
	q.Add("auth", u.Email)
	c.Request.URL.RawQuery = q.Encode()

	switch p := provider; p {
	case "google", "facebook":
		name = fmt.Sprintf("%v %v", u.FirstName, u.LastName)
		account = u.Email
		avatarURL = u.AvatarURL
	case "line":
		name = u.NickName
		account = u.UserID
		avatarURL = u.AvatarURL

	default:

	}

	newuser := models.NewCustomerModel{
		UUID:     util.GetUUID(),
		Name:     name,
		Account:  account,
		Image:    avatarURL,
		Provider: provider,
	}

	err = services.NewService().Customer(&newuser)

	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"code":    http.StatusBadRequest,
			"pienweb": "/",
			"message": err.Error(),
		})
		return
	}
	customer, exist := services.NewService().GetCustomerWithAccount(account)
	if !exist {

		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"code":    http.StatusBadRequest,
			"pienweb": "/",
			"message": "Not Found Customer",
		})
		return
	}

	token, err := jwt.Generate(customer.UUID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"code":    http.StatusBadRequest,
			"pienweb": "/",
			"message": err.Error(),
		})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "auth",
		Value:  token,
		MaxAge: int(config.Token_expire),
		Path:   "/",
	})

	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")

}

func (ct *controller) Email(c *gin.Context, name string, email string) {
	new_passcode := util.GetUUID()
	code := fmt.Sprintf("%v", util.RandInt(1000, 9999))

	// send passcode to mail

	err := services.NewService().SendMail(email, config.GetENV().URL, new_passcode, code)
	if err != nil {
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))

		return
	}

	uuid := util.GetUUID()
	newuser := models.NewCustomerModel{
		UUID:     uuid,
		Name:     name,
		Account:  email,
		Provider: "email",
	}

	err = services.NewService().Customer(&newuser)
	if err != nil {
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))

		return
	}

	err = services.NewService().NewPasscode(new_passcode, code, new_passcode, uuid)
	if err != nil {
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"passcode": new_passcode,
	})
}

func (ct *controller) Confirm(c *gin.Context, passcode string) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":    "login",
		"passcode": passcode,
	})
}

func (ct *controller) ConfirmCode(c *gin.Context, passcode string, code string) {

	result, confirm, err := services.NewService().VerifyCode(passcode, code)
	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"result":   "error",
			"passcode": passcode,
		})
	}

	if result == "VALID" {
		c.JSON(http.StatusOK, gin.H{
			"result":   result,
			"passcode": passcode,
			"confirm":  confirm,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":   result,
			"passcode": passcode,
		})
	}
}

func (ct *controller) LoginReady(c *gin.Context, cnfcode string) {
	uuid, err := services.NewService().WelcomeHome(cnfcode)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"code":    http.StatusBadRequest,
			"pienweb": "/",
			"message": err.Error(),
		})
		return
	}
	auth, _ := jwt.Generate(uuid)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "auth",
		Value:  auth,
		MaxAge: int(config.Token_expire),
		Path:   "/",
	})

	c.Redirect(http.StatusTemporaryRedirect, "/dashboard")

}
