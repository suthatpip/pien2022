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
	NewCompany(c *gin.Context, company *models.CompanyModel)
	CompanyNewLogo(c *gin.Context, com *models.CompanyModel)
	UpdateCompany(c *gin.Context, company *models.CompanyModel)

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

	ConfirmCode(c *gin.Context, passcode string, code string)
	Confirm(c *gin.Context, passcode string)
	LoginReady(c *gin.Context, cnfcode string)

	Account(c *gin.Context)
	UploadCompanyLogo(c *gin.Context, company_code string)
	DeleteCompany(c *gin.Context, company *models.CompanyModel)
}

type controller struct {
}

func NewController() controllerInterface {
	return &controller{}
}
