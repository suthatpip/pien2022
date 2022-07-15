package services

import (
	"database/sql"
	"piennews/helper/database"
	"piennews/helper/util"
	"piennews/models"

	_ "github.com/go-sql-driver/mysql"
)

func (s *service) AddPayment(payment *models.AddPaymentModel) error {

	db, err := sql.Open("mysql", database.Connect().ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()
	statement, err := db.Prepare(`INSERT INTO orders 
	(payment_code, company_code, customer_uuid, order_no, payment_due_date, tax_invoice_no, start_date, end_date, days, sub_total_baht, vat, total_baht)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`)

	if err != nil {
		return err
	}
	_, err = statement.Exec(
		payment.Payment_code,
		payment.Company_Code,
		payment.Customer_UUID,
		payment.Order_No,
		payment.Payment_Due_Date,
		payment.Tax_Invoice_No,
		payment.Start_Date,
		payment.End_Date,
		payment.Days,
		payment.Sub_Total_Baht,
		payment.VAT,
		payment.Total_Baht,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetPaymentDetail(pay_code string, uuid string) (*models.SummaryPaymentModel, error) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return nil, err
	}
	defer db.Close()

	sql := `SELECT  
	p.product_name, 
	CONCAT(cus.first_name, ' ', cus.last_name) as name, 
	com.name as company, 
	com.code, 
	com.address, 
	com.telephone, 
	IFNULL(com.logo,''), 
	o.payment_code, 
	o.order_no, 
	DATE_FORMAT(o.payment_due_date,'%d/%m/%Y'), 
	o.tax_invoice_no, 
	DATE_FORMAT(o.start_date,'%d/%m/%Y'), 
	DATE_FORMAT(o.end_date,'%d/%m/%Y'), 	 
	o.days, 
	o.sub_total_baht, 
	o.vat, 
	o.total_baht 
	FROM orders o 
	inner join company com on o.company_code = com.code 
	inner join customer cus on com.customer_uuid =cus.uuid 
	inner join order_product pp on pp.payment_code= o.payment_code 
	inner join product p on p.product_code =pp.product_code 
	where o.payment_code = ? AND o.customer_uuid=?;`

	list, err := db.Query(sql, pay_code, uuid)
	if err != nil {
		return nil, err
	}
	products := []models.SummaryProductModel{}

	var product_name, customer_name, name, code, address, telephone, logo, payment_code, order_no, payment_due, tax_invoice_no, start_date, end_date, days, sub_total, vat, total string
	no := 0
	for list.Next() {
		err := list.Scan(
			&product_name,
			&customer_name,
			&name,
			&code,
			&address,
			&telephone,
			&logo,
			&payment_code,
			&order_no,
			&payment_due,
			&tax_invoice_no,
			&start_date,
			&end_date,
			&days,
			&sub_total,
			&vat,
			&total,
		)
		no++
		product := models.SummaryProductModel{
			No:           no,
			Product:      product_name,
			Start_Date:   util.DateTH(start_date),
			End_Date:     util.DateTH(end_date),
			Days:         days,
			Product_Baht: "49 บาท",
		}
		products = append(products, product)
		if err != nil {
			panic(err.Error())
		}
	}

	payment := models.SummaryPaymentModel{
		Customer_Name:      customer_name,
		Publish_Start_Date: start_date,
		Publish_End_Date:   end_date,
		Company_Detail: &models.SummaryCompanyModel{
			Name:      name,
			Address:   address,
			Telephone: telephone,
			Logo:      logo,
			Code:      code,
		},
		Order_Detail: &models.SummaryOrderModel{
			Order_No:       order_no,
			Payment_Due:    payment_due,
			Tax_invoice_No: tax_invoice_no,
			Payment_code:   payment_code,
		},
		Products_Detail: &products,
		Sub_Total:       sub_total,
		VAT:             vat,
		Total:           total,
	}

	return &payment, nil
}

func (s *service) DeletePayment(p *models.DeleteInitPayment, uuid string) error {

	db, err := sql.Open("mysql", database.Connect().ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()
	statement, err := db.Prepare(`DELETE FROM orders 
	WHERE payment_code =? and customer_uuid=? and status=0;`)

	if err != nil {
		return err
	}
	_, err = statement.Exec(p.Payment_code, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteProductAndPayment(p *models.DeleteInitPayment, uuid string) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()

	tx, err := db.Begin()

	sql := `DELETE FROM product 
	WHERE product_code IN ( 
		SELECT pp.product_code  
		FROM orders o INNER JOIN order_product pp ON o.payment_code = pp.payment_code 
		WHERE o.payment_code =? AND o.customer_uuid =? AND o.status =0 
	) AND product_type ='template';`

	_, err = tx.Exec(sql, p.Payment_code, uuid)
	if err != nil {
		tx.Rollback()
		return err
	}
	sql = `DELETE FROM orders WHERE payment_code =? and customer_uuid=? and status=0;`

	_, err = tx.Exec(sql, p.Payment_code, uuid)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateOrderStatus(pay_code string, uuid string, status string) error {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()
	statement, err := db.Prepare(`UPDATE orders SET update_date=now(), status=? WHERE payment_code=? AND customer_uuid=?;`)

	if err != nil {
		return err
	}
	_, err = statement.Exec(status, pay_code, uuid)
	if err != nil {
		return err
	}

	return nil
}
