package services

import (
	"database/sql"
	"fmt"
	"piennews/helper/database"
	"piennews/helper/util"
	"piennews/helper/util/timeago"
	"piennews/models"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func (sv *service) NewProduct(f *models.ProductModel, uuid string) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()

	statement, err := db.Prepare(`INSERT INTO product 
	(product_code, product_name, product_detail, product_size, product_type, uuid) 
	VALUES(?, ?, ?, ?, ?, ?);`)

	if err != nil {
		return err
	}
	_, err = statement.Exec(
		f.Product_Code,
		f.Product_Name,
		f.Product_Detail,
		f.Product_Size,
		f.Product_Type,
		uuid,
	)
	if err != nil {
		return err
	}

	return nil
}

func (sv *service) GetProduct(uuid string) ([]models.ProductModel, error) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return nil, err
	}

	defer db.Close()

	sql := `SELECT distinct  
	p.product_code,  
	CASE when COALESCE(pp.product_code, 0)>0 THEN 'connect' ELSE 'disconnect' end as relation,
   product_name, 
   product_detail, 
   product_size, 
   TIMESTAMPDIFF(second ,NOW(),p.create_date) as create_date
   FROM product p left join order_product pp on p.product_code = pp.product_code 
   WHERE p.product_type='file' 
   AND p.uuid=?
   ORDER BY p.create_date DESC;`

	list, err := db.Query(sql, uuid)
	if err != nil {
		return nil, err
	}
	f := models.ProductModel{}
	fs := []models.ProductModel{}
	for list.Next() {
		second := ""
		size := ""
		err := list.Scan(&f.Product_Code, &f.Product_Connect, &f.Product_Name, &f.Product_Detail, &size, &second)
		if err != nil {
			panic(err.Error())
		}

		if s64, err := strconv.ParseInt(second, 10, 64); err == nil {
			t := time.Now().Add(time.Duration(s64) * time.Second)
			f.Product_Create_Date = timeago.Thailand.Format(t)

		}

		if size64, err := strconv.ParseFloat(size, 64); err == nil {
			f.Product_Size = util.HumanFileSize(size64)
		}
		fs = append(fs, f)
	}
	return fs, nil
}

func (sv *service) SubmitProduct(f []models.InitProductModel, payment_code string) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()
	sql := "INSERT INTO order_product (product_code, payment_code) VALUES "
	for _, element := range f {
		sql += `('` + element.Code + `', '` + payment_code + `'),`
	}
	sql = fmt.Sprintf("%v;", strings.TrimSuffix(sql, ","))

	statement, err := db.Prepare(sql)

	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (sv *service) DelProduct(p *models.ProductModel, uuid string) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()
	sql := `DELETE FROM product
	WHERE product_code not in (
	SELECT product_code 
		FROM order_product  
	) and product_code= ? and uuid= ?; `

	statement, err := db.Prepare(sql)

	if err != nil {
		return err
	}
	_, err = statement.Exec(p.Product_Code, uuid)
	if err != nil {
		return err
	}

	return nil
}
