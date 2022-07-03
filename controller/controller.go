package controller

import (
	"piennews/models"

	"github.com/gin-gonic/gin"
)

type controllerInterface interface {
	GetTemplate(c *gin.Context, code string)
	InitPayment(c *gin.Context, v *models.InitPaymentModel)
	CompanyList(c *gin.Context)
	CompanyNew(c *gin.Context, com *models.CompanyModel)
	CompanyNewLogo(c *gin.Context, com *models.CompanyModel)
}

type controller struct {
}

func NewController() controllerInterface {
	return &controller{}
}
