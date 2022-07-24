package models

type CustomerModel struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type NewCustomerModel struct {
	UUID     string
	Name     string
	Account  string
	Image    string
	Provider string
}
