package routers

import (
	"piennews/middleware"

	"github.com/gin-gonic/gin"
)

func Gateway(router *gin.Engine) {

	Home(router.Group("/"))

	router.Use(middleware.AuthRequired())
	{
		Template(router.Group("/template"))
		Payment(router.Group("/payment"))
		Company(router.Group("/company"))
		Upload(router.Group("/upload"))
		Dashboard(router.Group("/dashboard"))
		Account(router.Group("/account"))
	}

	// m := router.Group("/inbox")
	// grouproutes.Inbox(m)

	// u := router.Group("/user")
	// grouproutes.User(u)

	// a := router.Group("/auth")
	// grouproutes.Auth(a)

	// t := router.Group("/template")
	// grouproutes.Template(t)

	// pay := router.Group("/payment")
	// grouproutes.Payment(pay)

	// ps := router.Group("/paysolution")
	// grouproutes.Paysolution(ps)

	// page := router.Group("/page")
	// grouproutes.Page(page)

	// pt := router.Group("/posts")
	// grouproutes.Posts(pt)

	// nw := router.Group("/news")
	// grouproutes.News(nw)

	// adm := router.Group("/admin")
	// grouproutes.Admin(adm)

	// web := router.Group("/web")
	// grouproutes.WebScraper(web)

}
