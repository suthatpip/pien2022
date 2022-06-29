package company

import (
	"database/sql"
	"fmt"
	"piennews/helper/database"
	"piennews/models"

	_ "github.com/go-sql-driver/mysql"
)

type serviceInterface interface {
	GetCompany(company_id string) models.CompanyModel
}

type service struct {
}

func NewService() serviceInterface {
	return &service{}
}

func (sv *service) GetCompany(company_id string) models.CompanyModel {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		fmt.Printf("%v", err.Error())
		return models.CompanyModel{}
	}

	defer db.Close()
	com := models.CompanyModel{}

	err = db.QueryRow(`SELECT IFNULL(name,'') as name , IFNULL(address,'') as address , IFNULL(telephone,'') as telephone , IFNULL(logo,'') as logo FROM company WHERE company_id=?`, company_id).Scan(
		&com.Name,
		&com.Address,
		&com.Telephone,
		&com.Logo,
	)
	if err != nil {

		return models.CompanyModel{}
	}
	return com
}
