package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Account(rg *gin.RouterGroup) {

	rg.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "account.html", nil)
	})
}
