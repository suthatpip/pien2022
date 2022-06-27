package controller

import (
	"fmt"
	"net/http"
	"piennews/helper/config"
	"piennews/helper/logs"
	"piennews/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nleeper/goment"
)

func (ct *controller) Verify(c *gin.Context, v *models.VerifyModel) {
	req := ""
	res := ""
	message := ""
	statusCode := http.StatusOK
	defer func(begin time.Time) {
		logs.NewLogs(&logs.LogParams{
			Begin:       begin,
			Context:     c,
			Request:     req,
			Response:    res,
			Status:      statusCode,
			Source:      config.Get().Owner,
			Destination: "internal",
			Error:       message,
		}).Write()
	}(time.Now())

	start, _ := goment.New(v.Start_Date, "DD/MM/YYYY")
	end, _ := goment.New(v.End_Date, "DD/MM/YYYY")

	loop := end.Diff(start, "days") + 1
	fmt.Println(loop)
	ps := []models.SummaryProductModel{}
	for i := 0; i < loop; i++ {
		booking, _ := goment.New(v.Start_Date, "DD/MM/YYYY")
		p := models.SummaryProductModel{
			No:      i + 1,
			Product: v.File_Name,
			Date:    booking.Add(i, "d").Format("DD/MM/YYYY"),
			Price:   "49",
		}
		ps = append(ps, p)
	}

	s := models.SummaryPaymentModel{
		Customer_Detail: &models.SummaryCustomerModel{
			Name:      "x",
			Address:   "x",
			Telephone: "x",
			Email:     "x",
		},
		Order_Detail: &models.SummaryOrderModel{
			Order_Id:    "x",
			Payment_due: "x",
			Account:     "x",
			Email:       "x",
		},
		Products_Detail: &ps,
		Sub_Total:       "x",
		Vat:             "x",
		Discount:        "x",
		Total:           "x",
	}

	c.JSON(http.StatusOK, s)

}
