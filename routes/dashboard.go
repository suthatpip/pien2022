package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(rg *gin.RouterGroup) {

	rg.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "dashboard.html", nil)
	})
}
