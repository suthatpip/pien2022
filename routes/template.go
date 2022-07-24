package routers

import (
	"net/http"
	"piennews/controller"
	"piennews/models"

	"github.com/gin-gonic/gin"
)

type templateCode struct {
	Code string `uri:"code" binding:"required"`
}

func Template(rg *gin.RouterGroup) {

	rg.POST("/:code", func(c *gin.Context) {
		var t templateCode
		if err := c.ShouldBindUri(&t); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"code":    http.StatusBadRequest,
				"pienweb": "/",
				"message": err.Error(),
			})
			return
		}
		controller.NewController().GetTemplate(c, t.Code)
	})

	rg.POST("/new", func(c *gin.Context) {
		fmodel := models.ProductModel{}
		if err := c.ShouldBindJSON(&fmodel); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		controller.NewController().CustomeFile(c, &fmodel)
	})

	rg.GET("/a", func(c *gin.Context) {

		controller.NewController().Template(c, "A")
	})
	rg.GET("/b", func(c *gin.Context) {

		controller.NewController().Template(c, "B")
	})
	rg.GET("/c", func(c *gin.Context) {

		controller.NewController().Template(c, "C")
	})

}
