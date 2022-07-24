package controller

import (
	"fmt"
	"net/http"
	"piennews/helper/config"
	"piennews/helper/jwt"
	"piennews/helper/logs"
	"piennews/helper/util"
	"piennews/models"
	"piennews/services"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (ct *controller) Auth(c *gin.Context, provider string) {
	logbody := ""
	logerror := ""

	var name, account, avatarURL string

	defer func(begin time.Time) {
		logs.InternalLogs(&logs.LogInternalParams{
			Begin:   begin,
			Context: c,
			Body:    logbody,
			Error:   logerror,
		}).WriteInternalLogs()
	}(time.Now())

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

	c.JSON(http.StatusOK, gin.H{
		"passcode": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	})
}

func (ct *controller) Confirm(c *gin.Context, passcode string, code string) {
	time.Sleep(10 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"isConfirmed": false,
		"passcode":    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	})
}
