package models

type InitPaymentModel struct {
	Start_Date   string              `json:"start_date"`
	End_Date     string              `json:"end_date"`
	Product_Name string              `json:"product_name"`
	Company_Code string              `json:"company_code"`
	UUID         string              `json:"uuid"`
	Products     *[]InitProductModel `json:"products"`
}

type InitProductModel struct {
	Code string `json:"code"`
	Data string `json:"data"`
}

type AddPaymentModel struct {
	Start_Date       string
	End_Date         string
	Days             string
	Sub_Total_Baht   string
	VAT              string
	Total_Baht       string
	Customer_UUID    string
	Company_Code     string
	Order_No         string
	Payment_Due_Date string
	Tax_Invoice_No   string
	Payment_code     string
}

type SummaryPaymentModel struct {
	Customer_Name      string                 `json:"customer_name"`
	Publish_Start_Date string                 `json:"publish_start_date"`
	Publish_End_Date   string                 `json:"publish_end_date"`
	Company_Detail     *SummaryCompanyModel   `json:"company"`
	Order_Detail       *SummaryOrderModel     `json:"order"`
	Products_Detail    *[]SummaryProductModel `json:"products"`
	Sub_Total          string                 `json:"sub_total"`
	VAT                string                 `json:"vat"`
	Total              string                 `json:"total"`
	Create_Date        string                 `json:"create_date,omitemtry"`
}

type SummaryCompanyModel struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Logo      string `json:"logo"`
	Code      string `json:"code,omitemtry"`
}

type SummaryOrderModel struct {
	Order_No         string `json:"order_no,omitemtry"`
	Payment_Due_Date string `json:"payment_due_date,omitemtry"`
	Tax_invoice_No   string `json:"tax_invoice_no,omitemtry"`
	Payment_Code     string `json:"payment_code,omitemtry"`
	Payment_Date     string `json:"payment_date,omitemtry"`
}

type SummaryProductModel struct {
	No           int    `json:"no"`
	Product      string `json:"product"`
	Start_Date   string `json:"start_date"`
	End_Date     string `json:"end_date"`
	Days         string `json:"days"`
	Product_Baht string `json:"product_baht"`
	Detail       string `json:"detail,omitemtry"`
	Type         string `json:"type,omitemtry"`
	Size         string `json:"size,omitemtry"`
}

type DeleteInitPayment struct {
	Payment_code string `json:"payment_code"`
}

type SubmitPayment struct {
	Payment_code string `uri:"code" binding:"required"`
}

type QueryPayment struct {
	Payment_code string `uri:"code" binding:"required"`
}
