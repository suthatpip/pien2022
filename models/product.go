package models

type ProductsModel struct {
	Products []ProductModel `json:"products"`
}

type ProductModel struct {
	Product_Code        string `json:"product_code"`
	Product_Name        string `json:"product_name"`
	Product_Size        string `json:"product_size"`
	Product_Detail      string `json:"product_detail"`
	Product_Type        string `json:"product_type"`
	Product_Create_Date string `json:"product_create_date"`
	Product_Connect     string `json:"product_connect"`
	Template_code       string `json:"template_code"`
}
