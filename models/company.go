package models

type CompanyModel struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	Logo        string `json:"logo"`
	Code        string `json:"code" uri:"code" binding:"required"`
	ID          string `json:"id,omitempty"`
	UUID        string `json:"customer_uuid,omitempty"`
	About       string `json:"about"`
	Create_Date string `json:"create_date"`
}
