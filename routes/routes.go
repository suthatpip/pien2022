package routers

import (
	"net/http"
	"piennews/middleware"

	"github.com/gin-gonic/gin"
)

func Gateway(router *gin.Engine) {

	Home(router.Group("/"))
	Login(router.Group("/auth"))

	template := router.Group("/template")
	template.Use(middleware.AuthRequired())
	{
		Template(template)
	}

	payment := router.Group("/payment")
	payment.Use(middleware.AuthRequired())
	{
		Payment(payment)
	}

	company := router.Group("/company")
	company.Use(middleware.AuthRequired())
	{
		Company(company)
	}

	dashboard := router.Group("/dashboard")
	dashboard.Use(middleware.AuthRequired())
	{
		Dashboard(dashboard)
	}

	account := router.Group("/account")
	account.Use(middleware.AuthRequired())
	{
		Account(account)
	}

	product := router.Group("/product")
	product.Use(middleware.AuthRequired())
	{
		Product(product)
	}

	paysolution := router.Group("/paysolution")
	paysolution.Use(middleware.AuthRequired())
	{
		Paysolution(product)
	}

	router.GET("/error", func(c *gin.Context) {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "sdfsdfsd",
		})
	})

}
