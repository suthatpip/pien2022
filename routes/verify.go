package routers

import (
	"net/http"

	"piennews/controller"
	"piennews/models"

	"github.com/gin-gonic/gin"
)

func Verify(tmpl *gin.RouterGroup) {
	tmpl.POST("/", func(c *gin.Context) {
		v := &models.VerifyModel{}

		if err := c.ShouldBindJSON(&v); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"code":    http.StatusBadRequest,
				"pienweb": "/",
				"message": err.Error(),
			})
			return
		}
		controller.NewController().Verify(c, v)
	})
}
