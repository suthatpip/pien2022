package services

import "piennews/models"

type serviceInterface interface {
	GetCompany(company_id string) models.CompanyModel
	GetCompanyList(uuid string) ([]models.CompanyModel, error)
	GetCustomer(uuid string) models.CustomerModel
	AddPayment(payment *models.AddPaymentModel) error
	GetPaymentDetail(pay_code string, uuid string) (*models.SummaryPaymentModel, error)
	DeletePayment(p *models.DeleteInitPayment, uuid string) error
	UpdateOrderStatus(pay_code string, uuid string, status string) error

	DeleteProductAndPayment(p *models.DeleteInitPayment, uuid string) error
	GetTemplate(code string) ([]interface{}, bool)
	SaveCompany(com *models.CompanyModel) bool
	SaveCompanyLogo(com *models.CompanyModel) bool
	NewProduct(f *models.ProductModel, uuid string) error
	GetProduct(uuid string) ([]models.ProductModel, error)
	SubmitProduct(f []models.InitProductModel, uuid string) error
	DelProduct(p *models.ProductModel, uuid string) error
	NewInitPaysolution(payment_code string) (int64, error)
	InquiryPaysolution(ref_no string) (*models.InquiryModel, error)
	EnquipryNextStep(ref_no string, status string) error
	GetOrderPrice(refno string) (float64, error)

	Dashboard(uuid string) (*models.DashboardModel, error)
	GetOrderDetail(pay_code string, uuid string) (*models.SummaryPaymentModel, error)
}

type service struct {
}

func NewService() serviceInterface {
	return &service{}
}
