package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(rg *gin.RouterGroup) {

	rg.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
	})

	rg.GET("/support", func(c *gin.Context) {
		c.HTML(http.StatusOK, "support.html", gin.H{
			"title": "support",
		})
	})

}
