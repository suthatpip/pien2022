package models

type SummaryPaymentModel struct {
	Company_Detail  *SummaryCompanyModel `json:"customer"`
	Order_Detail    *SummaryOrderModel   `json:"order"`
	Products_Detail *SummaryProductModel `json:"product"`
	Sub_Total       string               `json:"sub_total"`
	Vat             string               `json:"vat"`
	Discount        string               `json:"discount"`
	Total           string               `json:"total"`
}

type SummaryCompanyModel struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Logo      string `json:"logo"`
}

type SummaryOrderModel struct {
	Order_Id    string `json:"order_no"`
	Payment_due string `json:"payment_due"`
	Account     string `json:"account"`
}

type SummaryProductModel struct {
	No         int    `json:"no"`
	Product    string `json:"product"`
	Start_Date string `json:"start_date"`
	End_Date   string `json:"end_date"`
	Sum_Price  string `json:"sum_price"`
}
