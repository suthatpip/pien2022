package routers

import (
	"net/http"
	"piennews/controller"
	"piennews/helper/jwt"
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

		h := c.MustGet("headers").(models.Header)
		v.UUID = jwt.ExtractClaims(h.Token, "uuid")

		controller.NewController().InitPayment(c, v)
	})

	rg.POST("/delete", func(c *gin.Context) {

		del := &models.DeleteInitPayment{}
		if err := c.ShouldBindJSON(&del); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"code":    http.StatusBadRequest,
				"pienweb": "/",
				"message": err.Error(),
			})
			return
		}

		controller.NewController().DeleteInitPayment(c, del)
	})

	rg.POST("/delete/all", func(c *gin.Context) {

		del := &models.DeleteInitPayment{}
		if err := c.ShouldBindJSON(&del); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"code":    http.StatusBadRequest,
				"pienweb": "/",
				"message": err.Error(),
			})
			return
		}

		controller.NewController().DeleteInitAllPayment(c, del)
	})

}
