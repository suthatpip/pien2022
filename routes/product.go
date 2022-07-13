package routers

import (
	"net/http"
	"piennews/controller"
	"piennews/models"

	"github.com/gin-gonic/gin"
)

func Product(rg *gin.RouterGroup) {

	rg.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "product.html", nil)
	})

	rg.GET("/list", func(c *gin.Context) {
		controller.NewController().GetProduct(c)
	})

	rg.POST("/new", func(c *gin.Context) {
		controller.NewController().UploadFile(c)
	})

	rg.POST("/delete", func(c *gin.Context) {

		products := models.ProductsModel{}
		if err := c.ShouldBindJSON(&products); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		controller.NewController().DeleteProduct(c, &products)
	})

}
