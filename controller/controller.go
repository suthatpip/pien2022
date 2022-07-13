package controller

import (
	"piennews/models"

	"github.com/gin-gonic/gin"
)

type controllerInterface interface {
	GetTemplate(c *gin.Context, code string)
	InitPayment(c *gin.Context, v *models.InitPaymentModel)
	DeleteInitPayment(c *gin.Context, del *models.DeleteInitPayment)
	DeleteInitAllPayment(c *gin.Context, del *models.DeleteInitPayment)
	CompanyList(c *gin.Context)
	CompanyNew(c *gin.Context, com *models.CompanyModel)
	CompanyNewLogo(c *gin.Context, com *models.CompanyModel)
	UploadFile(c *gin.Context)
	GetProduct(c *gin.Context)
	CustomeFile(c *gin.Context, file *models.ProductModel)
	DeleteProduct(c *gin.Context, products *models.ProductsModel)

	ConfirmPayment(c *gin.Context, pay *models.SubmitPayment)
}

type controller struct {
}

func NewController() controllerInterface {
	return &controller{}
}
