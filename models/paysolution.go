package models

type PaysolutionModel struct {
	Id             string
	Ref_no         string
	Merchant_id    string
	Customere_mail string
	Product_detail string
	Total          string
	Card_type      string
	Access_token   string
}

type PaysolutionCallback struct {
	Ref_no         string `form:"refno" binding:"required"`
	Merchant_id    string `form:"merchantid" binding:"required"`
	Customere_mail string `form:"email"`
	Product_detail string `form:"productdetail" binding:"required"`
	Total          string `form:"total" binding:"required"`
	Card_type      string `form:"cardtype" binding:"required"`
}

type InquiryModel struct {
	ReferenceNo        string  `json:"ReferenceNo,omitemtry"`
	OrderNo            string  `json:"OrderNo,omitemtry"`
	MerchantID         int64   `json:"MerchantID,omitemtry"`
	ProductDetail      string  `json:"ProductDetail,omitemtry"`
	Total              float64 `json:"Total,omitemtry"`
	CardType           string  `json:"CardType,omitemtry"`
	CustomerEmail      string  `json:"CustomerEmail,omitemtry"`
	CurrencyCode       string  `json:"CurrencyCode,omitemtry"`
	Status             string  `json:"Status,omitemtry"`
	StatusName         string  `json:"StatusName,omitemtry"`
	PostBackUrl        string  `json:"PostBackUrl,omitemtry"`
	PostBackParameters string  `json:"PostBackParameters,omitemtry"`
	PostBackMethod     string  `json:"PostBackMethod,omitemtry"`
	PostBackCompleted  bool    `json:"PostBackCompleted,omitemtry"`
	OrderDateTime      string  `json:"OrderDateTime,omitemtry"`
	Installment        string  `json:"installment,omitemtry"`
}

type PaysolutionInquiry []InquiryModel
