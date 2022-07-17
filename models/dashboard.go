package models

type DashboardModel struct {
	Summary *DashboardSummaryModel `json:"summary"`
	Orders  *[]DashboardOrderModel `json:"orders"`
}

type DashboardSummaryModel struct {
	PENDING_PAYMENT string
	ON_PROCESS      string
	PUBLISH         string
}

type DashboardOrderModel struct {
	Product_Name         string
	Order_No             string
	Order_Status         string
	Order_Status_Message string
	Order_Status_Level   string
	Order_Create_Date    string
	Order_Total          string
	Start_Date           string
	End_Date             string
	Payment_Due_Date     string
	Payment_Date         string
	Payment_Code         string
	Days                 string
}
