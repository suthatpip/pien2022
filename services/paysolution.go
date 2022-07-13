package services

import (
	"database/sql"
	"piennews/helper/database"
	"piennews/helper/util"

	_ "github.com/go-sql-driver/mysql"
)

func (sv *service) NewInitPaysolution(payment_code string) (int64, error) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return 0, err
	}

	defer db.Close()
	ref_no := util.GetUniqNumber()
	statement, err := db.Prepare(`INSERT INTO paysolution 
	(payment_code, paysolution_ref_no) VALUES(?, ?);`)

	if err != nil {
		return 0, err
	}
	_, err = statement.Exec(payment_code, ref_no)
	if err != nil {
		return 0, err
	}

	return 0, nil
}
