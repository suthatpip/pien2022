package services

import (
	"database/sql"
	"fmt"
	"piennews/helper/database"
	"piennews/models"

	_ "github.com/go-sql-driver/mysql"
)

func (sv *service) GetCompanyList(uuid string) ([]models.CompanyModel, error) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return nil, err
	}

	defer db.Close()

	sql := `SELECT code, IFNULL(name,'') as name , IFNULL(address,'') as address , IFNULL(telephone,'') as telephone , IFNULL(logo,'') as logo 
	FROM company WHERE customer_uuid=?;`

	list, err := db.Query(sql, uuid)
	if err != nil {
		return nil, err
	}
	com := models.CompanyModel{}
	coms := []models.CompanyModel{}
	for list.Next() {

		err := list.Scan(&com.ID,
			&com.Name,
			&com.Address,
			&com.Telephone,
			&com.Logo)
		if err != nil {
			panic(err.Error())
		}

		coms = append(coms, com)
	}
	return coms, nil

}

func (sv *service) GetCompany(company_id string) models.CompanyModel {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		fmt.Printf("%v", err.Error())
		return models.CompanyModel{}
	}

	defer db.Close()
	com := models.CompanyModel{}

	err = db.QueryRow(`SELECT id, IFNULL(name,'') as name , IFNULL(address,'') as address , IFNULL(telephone,'') as telephone , IFNULL(logo,'') as logo FROM company WHERE id=?`, company_id).Scan(
		&com.ID,
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

func (sv *service) SaveCompany(com *models.CompanyModel) bool {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return false
	}

	defer db.Close()

	statement, err := db.Prepare(`UPDATE company SET 
	name=?, address=?, telephone=? 
	WHERE code =? AND customer_uuid=? `)

	if err != nil {
		return false
	}
	_, err = statement.Exec(com.Name, com.Address, com.Telephone, com.Code, com.UUID)
	if err != nil {
		return false
	}
	return true
}
func (sv *service) SaveCompanyLogo(com *models.CompanyModel) bool {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return false
	}

	defer db.Close()

	statement, err := db.Prepare(`REPLACE INTO company
	(code, logo, customer_uuid) 
	VALUES(?, ?, ?); `)

	if err != nil {
		return false
	}
	_, err = statement.Exec(com.Code, com.Logo, com.UUID)
	if err != nil {
		return false
	}
	return true
}
