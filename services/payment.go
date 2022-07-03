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

	statement, err := db.Prepare(`INSERT INTO payment
	(payment_code, company_code, customer_uuid, order_no, payment_due_date, tax_invoice_no, file_name, start_date, end_date, days, sub_total_baht, vat, total_baht)
	VALUES(?, ?, ?, ?, ?  , ?, ?, ?, ?, ?, ?, ?, ?);`)

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
		payment.Product,
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

func (s *service) GetPayment(payment_code string) (*models.SummaryPaymentModel, error) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())

	if err != nil {
		return nil, err
	}
	defer db.Close()

	sql := `SELECT  
	CONCAT(cus.first_name, ' ', cus.last_name) as name, 
	com.name as company, 
	com.code, 
	com.address, 
	com.telephone, 
	IFNULL(com.logo,''), 
	pay.payment_code, 
	pay.order_no, 
	DATE_FORMAT(pay.payment_due_date,'%d/%m/%Y'), 
	pay.tax_invoice_no, 
	pay.file_name,  
	DATE_FORMAT(pay.start_date,'%d/%m/%Y'), 
	DATE_FORMAT(pay.end_date,'%d/%m/%Y'), 	 
	pay.days, 
	pay.sub_total_baht, 
	pay.vat, 
	pay.total_baht 
	FROM payment pay inner join company com on pay.company_code = com.code inner join customer cus on com.customer_uuid =cus.uuid 
	where payment_code = ? `

	list, err := db.Query(sql, payment_code)
	if err != nil {
		return nil, err
	}
	payment := models.SummaryPaymentModel{}
	company := models.SummaryCompanyModel{}
	order := models.SummaryOrderModel{}
	product := models.SummaryProductModel{
		No: 1,
	}

	for list.Next() {
		err := list.Scan(
			&payment.Customer_Name,
			&company.Name,
			&company.Code,
			&company.Address,
			&company.Telephone,
			&company.Logo,

			&order.Payment_code,
			&order.Order_No,
			&order.Payment_Due,
			&order.Tax_invoice_No,

			&product.Product,
			&product.Start_Date,
			&product.End_Date,
			&product.Days,

			&payment.Sub_Total,
			&payment.VAT,
			&payment.Total,
		)
		if err != nil {
			panic(err.Error())
		}
		product.Product_Baht = payment.Sub_Total
		payment.Company_Detail = &company

		order.Payment_Due = util.DateTH(order.Payment_Due)
		payment.Order_Detail = &order

		product.Start_Date = util.DateTH(product.Start_Date)
		product.End_Date = util.DateTH(product.End_Date)
		payment.Products_Detail = &product
	}

	return &payment, nil
}
