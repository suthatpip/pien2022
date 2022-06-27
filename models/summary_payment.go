package models

type SummaryPaymentModel struct {
	Customer_Detail *SummaryCustomerModel  `json:"customer"`
	Order_Detail    *SummaryOrderModel     `json:"order"`
	Products_Detail *[]SummaryProductModel `json:"product"`
	Sub_Total       string                 `json:"sub_total"`
	Vat             string                 `json:"vat"`
	Discount        string                 `json:"discount"`
	Total           string                 `json:"total"`
}

type SummaryCustomerModel struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
}

type SummaryOrderModel struct {
	Order_Id    string `json:"order_no"`
	Payment_due string `json:"payment_due"`
	Account     string `json:"account"`
	Email       string `json:"email"`
}

type SummaryProductModel struct {
	No      int    `json:"no"`
	Product string `json:"product"`
	Date    string `json:"date"`
	Price   string `json:"price"`
}
