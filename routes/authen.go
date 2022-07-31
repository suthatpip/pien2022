package routers

import (
	"net/http"
	"piennews/controller"
	"piennews/helper/config"
	"piennews/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/line"
)

func Login(rg *gin.RouterGroup) {

	url := config.GetENV().URL

	key := "pien"        // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New("896145657250-ldimd033j5un98sr7o5r9uukn4flb8rk.apps.googleusercontent.com", "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", url+"/auth/google/callback", "email", "profile"),
		facebook.New("936226287272219", "e92340c81631c29a1855e13a501395e3", url+"/auth/facebook/callback"),
		line.New("1656602577", "86d81b7bc1d7529a239f494c72e9f480", url+"/auth/line/callback", "profile", "openid", "email"),
	)

	rg.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "login",
		})

	})

	rg.GET("/:provider", func(c *gin.Context) {

		var p models.Provider
		if err := c.ShouldBindUri(&p); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"code":    http.StatusBadRequest,
				"pienweb": "/",
				"message": err.Error(),
			})
			return
		}

		q := c.Request.URL.Query()
		q.Add("provider", p.Name)
		c.Request.URL.RawQuery = q.Encode()
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})

	rg.GET("/:provider/callback", func(c *gin.Context) {

		var p models.Provider

		if err := c.ShouldBindUri(&p); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"code":    http.StatusBadRequest,
				"pienweb": "/",
				"message": err.Error(),
			})
			return
		}

		controller.NewController().Auth(c, p.Name)

	})

	rg.GET("/passcode/:passcode", func(c *gin.Context) {

		var p models.Confirm

		if err := c.ShouldBindUri(&p); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"code":    http.StatusBadRequest,
				"pienweb": "/",
				"message": err.Error(),
			})
			return
		}

		controller.NewController().Confirm(c, p.Passcode)

	})

	rg.GET("/ready/:passcode/:confirm", func(c *gin.Context) {

		login := &models.ConfirmCode{}

		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
			return
		}

		controller.NewController().LoginReady(c, login.ConfirmCode)

	})

	rg.POST("/newpasscode", func(c *gin.Context) {

		login := &models.Login{}

		if err := c.ShouldBindJSON(&login); err != nil {

			c.JSON(500, gin.H{"status": err.Error()})
			return
		}

		controller.NewController().Email(c, login.User, login.Email)

	})

	rg.POST("/passcode/:passcode/:code", func(c *gin.Context) {

		login := &models.ConfirmCode{}

		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
			return
		}

		controller.NewController().ConfirmCode(c, login.Passcode, login.Code)

	})

	rg.GET("/logout", func(c *gin.Context) {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "auth",
			MaxAge: -1,
			Path:   "/",
		})
		c.Redirect(http.StatusFound, "/")
	})

}
