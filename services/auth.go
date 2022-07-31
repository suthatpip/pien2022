package services

import (
	"database/sql"
	"fmt"
	"piennews/helper/database"

	_ "github.com/go-sql-driver/mysql"
)

func (sv *service) NewPasscode(passcode, code, confirm_code, uuid string) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare(`INSERT INTO passcode
	(passcode, code, confirm_code, uuid) VALUES(?, ?, ?, ?);`)

	if err != nil {
		return err
	}
	_, err = statement.Exec(passcode, code, confirm_code, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (sv *service) VerifyCode(passcode, code string) (string, string, error) {

	retry := 999
	validate := 0
	confirm_code := ""
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return "ERROR", "", err
	}

	defer db.Close()

	tx, err := db.Begin()

	sql := `UPDATE passcode set retry=retry+1 WHERE passcode=?;`

	_, err = tx.Exec(sql, passcode)
	if err != nil {
		tx.Rollback()
		return "ERROR", "", err
	}
	sql = fmt.Sprintf(`SELECT validate, retry, confirm_code FROM 
	(SELECT count(*) as validate, 1 as link  FROM passcode 
	WHERE passcode='%v' AND code=%v) as a
	JOIN 
	(SELECT 1 as link, retry, confirm_code  FROM passcode 
	WHERE passcode='%v' ) as b
	on a.link=b.link;`, passcode, code, passcode)

	err = tx.QueryRow(sql).Scan(&validate, &retry, &confirm_code)
	if err != nil {
		tx.Rollback()
		return "ERROR", "", err
	}

	err = tx.Commit()
	if err != nil {
		return "ERROR", "", err
	}

	if validate == 1 && retry <= 3 {
		return "VALID", confirm_code, nil
	} else if validate == 1 && retry > 3 {
		return "BLOCK", "", nil
	} else if validate == 0 && retry > 3 {
		return "BLOCK", "", nil
	} else {
		return "INVALID", "", nil
	}

}

func (sv *service) WelcomeHome(cnfcode string) (string, error) {
	uuid := ""
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return "", err
	}

	defer db.Close()

	tx, err := db.Begin()

	sql := fmt.Sprintf(`SELECT uuid FROM passcode WHERE confirm_code='%v';`, cnfcode)

	err = tx.QueryRow(sql).Scan(&uuid)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return uuid, err

}

func deletePasscode(passcode string) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()
	sql := `DELETE FROM passcode WHERE passcode = ?; `

	statement, err := db.Prepare(sql)

	if err != nil {
		return err
	}
	_, err = statement.Exec(passcode)
	if err != nil {
		return err
	}

	return nil
}
