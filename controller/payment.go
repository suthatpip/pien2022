package controller

import (
	"fmt"
	"net/http"
	"piennews/helper/config"
	"piennews/helper/logs"
	"piennews/helper/util"
	"piennews/models"
	"piennews/services"

	"strings"
	"time"

	"github.com/bojanz/currency"
	"github.com/gin-gonic/gin"
	"github.com/nleeper/goment"
)

func (ct *controller) InitPayment(c *gin.Context, v *models.InitPaymentModel) {
	req := ""
	res := ""
	errorMessage := ""
	statusCode := http.StatusOK
	defer func(begin time.Time) {
		logs.NewLogs(&logs.LogParams{
			Begin:       begin,
			Context:     c,
			Request:     req,
			Response:    res,
			Status:      statusCode,
			Source:      config.GetENV().Owner,
			Destination: "internal",
			Error:       errorMessage,
		}).Write()
	}(time.Now())

	payment_code, err := setPayment(v)
	if err != nil {
		errorMessage = err.Error()
		c.Status(http.StatusBadRequest)
		return
	}

	summary, err := services.NewService().GetPayment(payment_code)
	if err != nil {
		errorMessage = err.Error()
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, summary)

}

func setPayment(v *models.InitPaymentModel) (string, error) {
	locale := currency.NewLocale("th")
	formatter := currency.NewFormatter(locale)
	formatter.MaxDigits = 2
	formatter.MinDigits = 2
	formatter.CurrencyDisplay = currency.DisplayNone

	payment_code := getPaymentCode()
	company_code := v.Company_Code
	customer_uuid := ""
	start, _ := goment.New(v.Start_Date, "DD-MM-YYYY")
	end, _ := goment.New(v.End_Date, "DD-MM-YYYY")
	days := end.Diff(start, "days") + 1

	due_date, _ := goment.New(v.Start_Date, "DD-MM-YYYY")
	due_date = due_date.Subtract(1, "days")

	sub_total_baht, _ := currency.NewAmount(fmt.Sprintf("%v", config.Price*float64(days)), "THB")
	vat, _ := currency.NewAmount(fmt.Sprintf("%v", config.Price*float64(0.07)*float64(days)), "THB")
	total_baht, _ := currency.NewAmount(fmt.Sprintf("%v", config.Price*float64(days)), "THB")

	p := &models.AddPaymentModel{
		Product:          v.Product_Name,
		Start_Date:       start.Subtract(543, "year").Format("YYYY-MM-DD"),
		End_Date:         end.Subtract(543, "year").Format("YYYY-MM-DD"),
		Days:             fmt.Sprintf("%v", days),
		Sub_Total_Baht:   formatter.Format(sub_total_baht), //   fmt.Sprintf("%v", sub_total_baht),
		VAT:              formatter.Format(vat),
		Total_Baht:       formatter.Format(total_baht),
		Customer_UUID:    customer_uuid,
		Company_Code:     company_code,
		Order_No:         getOrderNo(),
		Payment_Due_Date: due_date.Subtract(543, "year").Format("YYYY-MM-DD"),
		Tax_Invoice_No:   "Unknow",
		Payment_code:     payment_code,
	}

	err := services.NewService().AddPayment(p)
	if err != nil {
		return "", err
	}

	return payment_code, nil
}

func getOrderNo() string {
	d, _ := goment.New(time.Now().Format("02-01-2006"), "DD-MM-YYYY")
	return fmt.Sprintf("A%v-%v-%v ", d.Format("YY"), util.RandInt(1000, 9999), strings.ToUpper(util.RandSeq(4)))
}

func getPaymentCode() string {
	// d, _ := goment.New(time.Now().Format("02-01-2006"), "DD-MM-YYYY")
	// return fmt.Sprintf("A%v-%v-%v ", d.Format("YY"), util.RandInt(1000, 9999), strings.ToUpper(util.RandSeq(4)))

	return strings.ToUpper(fmt.Sprintf("%v", util.RandSeq(10)))
}
