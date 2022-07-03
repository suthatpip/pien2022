package routers

import (
	"net/http"
	"piennews/controller"
	"piennews/models"

	"github.com/gin-gonic/gin"
)

func Payment(rg *gin.RouterGroup) {
	rg.POST("/init", func(c *gin.Context) {
		v := &models.InitPaymentModel{}
		if err := c.ShouldBindJSON(&v); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"code":    http.StatusBadRequest,
				"pienweb": "/",
				"message": err.Error(),
			})
			return
		}
		controller.NewController().InitPayment(c, v)
	})
}
