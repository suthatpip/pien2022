package services

import (
	"database/sql"
	"fmt"
	"piennews/helper/config"
	"piennews/helper/database"
	"piennews/models"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func (sv *service) Dashboard(uuid string) (*models.DashboardModel, error) {

	dashboard := models.DashboardModel{
		Summary: &models.DashboardSummaryModel{},
		Orders:  &[]models.DashboardOrderModel{},
	}
	sumary, err := getSummary(uuid)
	if err == nil {
		dashboard.Summary = sumary
	}
	orders, err := getOrders(uuid)
	if err == nil {
		dashboard.Orders = orders
	}

	return &dashboard, nil

}

func getSummary(uuid string) (*models.DashboardSummaryModel, error) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())
	if err != nil {
		return nil, err
	}
	defer db.Close()
	sql := `SELECT os.status_code, count(*) 
	FROM  orders o inner join order_status os on o.status = os.status_code 
	where customer_uuid = ? group by os.status_code;`

	list, err := db.Query(sql, uuid)
	if err != nil {
		return nil, err
	}
	summary := models.DashboardSummaryModel{
		PENDING_PAYMENT: "0",
		ON_PROCESS:      "0",
		PUBLISH:         "0",
	}

	for list.Next() {
		var status_code, count string
		err := list.Scan(
			&status_code,
			&count,
		)
		if err != nil {
			panic(err.Error())
		}

		switch status := status_code; status {
		case config.GetOrderStatus().PENDING_PAYMENT:
			summary.PENDING_PAYMENT = count
		case config.GetOrderStatus().ON_PROCESS:
			summary.ON_PROCESS = count
		case config.GetOrderStatus().PUBLISH:
			summary.PUBLISH = count
		}
	}
	return &summary, nil
}

func getOrders(uuid string) (*[]models.DashboardOrderModel, error) {
	db, err := sql.Open("mysql", database.Connect().ConnectionString())
	if err != nil {
		return nil, err
	}
	defer db.Close()
	sql := `SELECT CONCAT(o.order_no, '/', LPAD(o.order_seq, 8, 0) ) as order_no,   
	GROUP_CONCAT(p2.product_name  SEPARATOR ', '),
	o.status ,  
	DATE_FORMAT(o.start_date , "%d/%m/%y") as start_date,   
	DATE_FORMAT(o.end_date , "%d/%m/%y") as end_date,  
	DATE_FORMAT(o.create_date , "%d/%m/%y") as create_date, 
	IFNULL(o.payment_date,'') as payment_date ,  
	DATE_FORMAT(o.payment_due_date , "%d/%m/%y") as payment_due_date, 
	o.total_baht, o.days ,o.payment_code 
	FROM  orders o inner join order_status os on o.status = os.status_code 
	inner join company c on o.company_code = c.code 
	inner join paysolution p on o.payment_code =p.payment_code 
	inner join order_product op on o.payment_code  = op.payment_code 
	inner join product p2 on op.product_code = p2.product_code 
	where o.customer_uuid  = ?  
	group by o.order_no, o.status,
	o.start_date, o.end_date, o.create_date, o.payment_date, o.payment_due_date, o.total_baht, o.days, o.payment_code;`

	list, err := db.Query(sql, uuid)
	if err != nil {
		return nil, err
	}
	order := models.DashboardOrderModel{}
	orders := []models.DashboardOrderModel{}
	for list.Next() {

		err := list.Scan(
			&order.Order_No,
			&order.Product_Name,
			&order.Order_Status,
			&order.Start_Date,
			&order.End_Date,
			&order.Order_Create_Date,
			&order.Payment_Date,
			&order.Payment_Due_Date,
			&order.Order_Total,
			&order.Days,
			&order.Payment_Code,
		)

		if err != nil {
			panic(err.Error())
		}

		if s, err := strconv.ParseFloat(order.Order_Total, 64); err == nil {
			order.Order_Total = fmt.Sprintf("%v", s)
		}

		order.Order_Status_Message, order.Order_Status_Level = statusMessage(order.Order_Status)
		orders = append(orders, order)

	}
	return &orders, nil
}

func statusMessage(status string) (message string, level string) {
	switch s := status; s {
	case config.GetOrderStatus().INITIAL_ORDER:
		return fmt.Sprintf("สร้างออร์เดอร์"), "BUILD"
	case config.GetOrderStatus().APPROVED:
		return fmt.Sprintf("ยืนยันออร์เดอร์"), "BUILD"
	case config.GetOrderStatus().PENDING_PAYMENT:
		return fmt.Sprintf("รอการชำระเงิน"), "WAITPAYMENT"
	case config.GetOrderStatus().PAYMENT_COMPLETED:
		return fmt.Sprintf("ชำระเงินเรียบร้อย"), "PAYMENTCOMPLETE"
	case config.GetOrderStatus().ON_PROCESS:
		return fmt.Sprintf("กำลังประมวลผล"), "PROCESS"
	case config.GetOrderStatus().PUBLISH:
		return fmt.Sprintf("ประกาศ"), "ONLINE"
	case config.GetOrderStatus().FAILED:
		return fmt.Sprintf("เกิดความผิดพลาด"), "ERROR"
	case config.GetOrderStatus().CANCELED:
		return fmt.Sprintf("ยกเลิก"), "ERROR"
	case config.GetOrderStatus().VALIDATE_PAYMENT:
		return fmt.Sprintf("ตรวจสอบการชำระเงิน"), "PROCESS"
	case config.GetOrderStatus().COMPLETE:
		return fmt.Sprintf("เรียบร้อย"), "COMPLETE"
	}
	return "", ""
}
