package services

import (
	"database/sql"
	"piennews/helper/database"
	"piennews/helper/util/timeago"
	"piennews/models"
	"strconv"
	"time"

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

func (sv *service) NewCompany(com *models.CompanyModel) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare(`INSERT INTO company 
	(code, name, address, telephone, customer_uuid, about)  
	VALUES(?, ?, ?, ?, ?, ?); `)

	if err != nil {
		return err
	}
	_, err = statement.Exec(com.Code, com.Name, com.Address, com.Telephone, com.UUID, com.About)
	if err != nil {
		return err
	}
	return nil
}

func (sv *service) UpdateCompany(com *models.CompanyModel) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare(`UPDATE company SET name=?, address=?, telephone=?, about=? WHERE code=? AND customer_uuid=?;`)

	if err != nil {
		return err
	}
	_, err = statement.Exec(com.Name, com.Address, com.Telephone, com.About, com.Code, com.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (sv *service) UpdateCompanyLogo(com *models.CompanyModel) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare(`UPDATE company SET logo =? WHERE code=? AND customer_uuid=?;`)

	if err != nil {
		return err
	}
	_, err = statement.Exec(com.Logo, com.Code, com.UUID)
	if err != nil {

		return err
	}
	return nil
}

func (sv *service) GetMyCompany(uuid string) (*[]models.CompanyModel, error) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return nil, err
	}

	defer db.Close()

	sql := `SELECT code, IFNULL(name,'') as   name, IFNULL(about,'') as about , IFNULL(address,'') as address, IFNULL(telephone,'') as telephone, IFNULL(logo,'') as logo, TIMESTAMPDIFF(second ,NOW(),create_date) as create_date FROM company WHERE customer_uuid= ?; `

	list, err := db.Query(sql, uuid)
	if err != nil {
		return nil, err
	}
	company := models.CompanyModel{}
	companys := []models.CompanyModel{}
	for list.Next() {

		err := list.Scan(&company.Code, &company.Name, &company.About, &company.Address, &company.Telephone, &company.Logo, &company.Create_Date)
		if err != nil {
			panic(err.Error())
		}

		if s64, err := strconv.ParseInt(company.Create_Date, 10, 64); err == nil {
			t := time.Now().Add(time.Duration(s64) * time.Second)
			company.Create_Date = timeago.Thailand.Format(t)

		}

		companys = append(companys, company)

	}
	return &companys, nil
}

func (sv *service) DeleteMyCompany(uuid, code string) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()
	statement, err := db.Prepare(`DELETE FROM company WHERE customer_uuid= ? AND code =?;`)

	if err != nil {
		return err
	}
	_, err = statement.Exec(uuid, code)
	if err != nil {
		return err
	}

	return nil
}
