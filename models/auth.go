package models

type Provider struct {
	Name string `uri:"provider" binding:"required"`
}

type Login struct {
	User  string `json:"user" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type ConfirmCode struct {
	Passcode string `uri:"passcode" binding:"required"`
	Code     string `uri:"code" binding:"required"`
}
