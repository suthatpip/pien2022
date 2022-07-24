package routers

import (
	"net/http"
	"piennews/controller"
	"piennews/models"

	"github.com/gin-gonic/gin"
)

func Dashboard(rg *gin.RouterGroup) {

	rg.GET("/", func(c *gin.Context) {

		controller.NewController().Dashboard(c)
	})

	rg.GET("/:code", func(c *gin.Context) {
		payment := models.QueryPayment{}
		if err := c.ShouldBindUri(&payment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		controller.NewController().Dashboard(c, payment.Payment_code)
	})

	rg.POST("/order/detail/:code", func(c *gin.Context) {

		payment := models.QueryPayment{}
		if err := c.ShouldBindUri(&payment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		controller.NewController().DashboardOrderDetail(c, payment.Payment_code)
	})

}
