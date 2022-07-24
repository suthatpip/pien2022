package routers

import (
	"piennews/controller"

	"github.com/gin-gonic/gin"
)

func Account(rg *gin.RouterGroup) {

	rg.GET("/", func(c *gin.Context) {

		controller.NewController().Account(c)
	})
}
