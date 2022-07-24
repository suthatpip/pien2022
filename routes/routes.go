package routers

import (
	"net/http"
	"piennews/middleware"

	"github.com/gin-gonic/gin"
)

func Gateway(router *gin.Engine) {

	Home(router.Group("/"))
	Login(router.Group("/auth"))
	router.Use(middleware.AuthRequired())
	{

		Template(router.Group("/template"))
		Payment(router.Group("/payment"))
		Company(router.Group("/company"))
		Dashboard(router.Group("/dashboard"))
		Account(router.Group("/account"))
		Product(router.Group("/product"))
		Paysolution(router.Group("/paysolution"))

	}

	router.GET("/error", func(c *gin.Context) {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "sdfsdfsd",
		})
	})

}
