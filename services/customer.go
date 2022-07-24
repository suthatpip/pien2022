package services

import (
	"database/sql"
	"piennews/helper/database"
	"piennews/models"

	_ "github.com/go-sql-driver/mysql"
)

func (sv *service) GetCustomerWithAccount(account string) (*models.CustomerModel, bool) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return nil, false
	}

	defer db.Close()
	cus := models.CustomerModel{}

	err = db.QueryRow(`SELECT uuid, name, image FROM customer WHERE account=?`, account).Scan(
		&cus.UUID,
		&cus.Name,
		&cus.Image,
	)
	if err != nil {
		return nil, false
	}

	return &cus, true
}

func (sv *service) GetCustomerWithUUID(uuid string) (*models.CustomerModel, bool) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return nil, false
	}

	defer db.Close()
	cus := models.CustomerModel{}

	err = db.QueryRow(`SELECT uuid, name, image FROM customer WHERE uuid=?`, uuid).Scan(
		&cus.UUID,
		&cus.Name,
		&cus.Image,
	)
	if err != nil {
		return nil, false
	}

	return &cus, true
}

func (sv *service) Customer(u *models.NewCustomerModel) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()

	sql := `INSERT INTO customer (uuid, name, account, image, provider) 
	  VALUES(?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE last_update=now(); `

	statement, err := db.Prepare(sql)

	if err != nil {
		return err
	}
	_, err = statement.Exec(u.UUID, u.Name, u.Account, u.Image, u.Provider)
	if err != nil {
		return err
	}

	return nil
}
