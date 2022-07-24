package controller

import (
	"piennews/models"

	"github.com/gin-gonic/gin"
)

type controllerInterface interface {
	Template(c *gin.Context, template string)
	GetTemplate(c *gin.Context, code string)
	InitPayment(c *gin.Context, v *models.InitPaymentModel)
	DeleteInitPayment(c *gin.Context, del *models.DeleteInitPayment)
	DeleteInitAllPayment(c *gin.Context, del *models.DeleteInitPayment)
	CompanyList(c *gin.Context)
	CompanyNew(c *gin.Context, com *models.CompanyModel)
	CompanyNewLogo(c *gin.Context, com *models.CompanyModel)
	UploadFile(c *gin.Context)
	Product(c *gin.Context)
	GetProduct(c *gin.Context)
	CustomeFile(c *gin.Context, file *models.ProductModel)
	DeleteProduct(c *gin.Context, products *models.ProductsModel)

	ConfirmPayment(c *gin.Context, pay *models.SubmitPayment)

	Dashboard(c *gin.Context, status ...string)
	DashboardOrderDetail(c *gin.Context, payment_code string)

	Auth(c *gin.Context, provider string)
	Email(c *gin.Context, name string, email string)
	Account(c *gin.Context)
	Confirm(c *gin.Context, passcode string, code string)
}

type controller struct {
}

func NewController() controllerInterface {
	return &controller{}
}
