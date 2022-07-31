package routers

import (
	"net/http"
	"piennews/controller"
	"piennews/helper/util"
	"piennews/models"

	"github.com/gin-gonic/gin"
)

func Account(rg *gin.RouterGroup) {

	rg.GET("/", func(c *gin.Context) {
		controller.NewController().Account(c)
	})

	rg.POST("/:code/logo", func(c *gin.Context) {

		company := models.CompanyModel{}
		if err := c.ShouldBindUri(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		controller.NewController().UploadCompanyLogo(c, company.Code)

	})

	rg.POST("/company", func(c *gin.Context) {
		company := models.CompanyModel{}
		if err := c.ShouldBindJSON(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		controller.NewController().UpdateCompany(c, &company)
	})

	rg.POST("/company/del/:code", func(c *gin.Context) {
		company := models.CompanyModel{}
		if err := c.ShouldBindUri(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		controller.NewController().DeleteCompany(c, &company)
	})

	rg.POST("/company/init", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": util.GetUUID(),
		})
	})

	rg.POST("/company/new", func(c *gin.Context) {
		company := models.CompanyModel{}
		if err := c.ShouldBindJSON(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		controller.NewController().NewCompany(c, &company)
	})

}
