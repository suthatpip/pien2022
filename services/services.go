package services

import "piennews/models"

type serviceInterface interface {
	GetCompany(company_id string) models.CompanyModel
	GetCompanyList(uuid string) ([]models.CompanyModel, error)
	GetCustomer(uuid string) models.CustomerModel
	AddPayment(payment *models.AddPaymentModel) error
	GetPayment(payment_code string) (*models.SummaryPaymentModel, error)
	GetTemplate(code string) ([]interface{}, bool)
	SaveCompany(com *models.CompanyModel) bool
	SaveCompanyLogo(com *models.CompanyModel) bool
}

type service struct {
}

func NewService() serviceInterface {
	return &service{}
}
