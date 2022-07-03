package services

import (
	"database/sql"
	"piennews/helper/database"
	"piennews/models"

	_ "github.com/go-sql-driver/mysql"
)

func (sv *service) GetCustomer(uuid string) models.CustomerModel {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return models.CustomerModel{}
	}

	defer db.Close()
	cus := models.CustomerModel{}

	err = db.QueryRow(`SELECT cus.uuid FROM customer WHERE uuid=?`, uuid).Scan(&cus.UUID)
	if err != nil {
		return models.CustomerModel{}
	}
	return cus
}
