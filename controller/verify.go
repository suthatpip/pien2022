package controller

import (
	"fmt"
	"net/http"
	"piennews/helper/config"
	"piennews/helper/logs"
	"piennews/helper/util"
	"piennews/models"
	"piennews/services/company"
	"strings"
	"time"

	"github.com/bojanz/currency"
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
	uuid := "81107e82-8425-454c-b630-8b5979c3b8ef"
	company_id := "ABCDEFG"
	start, _ := goment.New(v.Start_Date, "DD/MM/YYYY")
	end, _ := goment.New(v.End_Date, "DD/MM/YYYY")
	days := end.Diff(start, "days") + 1

	due_date, _ := goment.New(v.Start_Date, "DD/MM/YYYY")
	due_date = due_date.Subtract(1, "days")

	amount, _ := currency.NewAmount(fmt.Sprintf("%v", config.Price*float64(days)), "THB")
	vat, _ := currency.NewAmount(fmt.Sprintf("%v", config.Price*float64(0.07)*float64(days)), "THB")
	discount, _ := currency.NewAmount(fmt.Sprintf("0"), "THB")
	total, _ := currency.NewAmount(fmt.Sprintf("%v", config.Price*float64(days)), "THB")

	locale := currency.NewLocale("th")
	formatter := currency.NewFormatter(locale)
	formatter.MaxDigits = 2
	formatter.MinDigits = 2
	formatter.CurrencyDisplay = currency.DisplayNone
	company := company.NewService().GetCompany(company_id)

	s := models.SummaryPaymentModel{
		Company_Detail: &models.SummaryCompanyModel{
			Name:      company.Name,
			Address:   company.Address,
			Telephone: company.Telephone,
			Logo:      company.Logo,
		},
		Order_Detail: &models.SummaryOrderModel{
			Order_Id:    getOrderNo(),
			Payment_due: util.DateTH(due_date.Format("DD/MM/YYYY")),
			Account:     uuid,
		},
		Products_Detail: &models.SummaryProductModel{
			No:         1,
			Product:    v.File_Name,
			Start_Date: util.DateTH(start.Format("DD/MM/YYYY")),
			End_Date:   util.DateTH(end.Format("DD/MM/YYYY")),
			Sum_Price:  formatter.Format(amount),
		},
		Sub_Total: formatter.Format(amount),
		Vat:       formatter.Format(vat),
		Discount:  formatter.Format(discount),
		Total:     formatter.Format(total),
	}
	c.JSON(http.StatusOK, s)
}

func getOrderNo() string {
	d, _ := goment.New(time.Now().Format("02-01-2006"), "DD-MM-YYYY")
	return fmt.Sprintf("P%v-%v-%v ", d.Format("YY"), util.RandInt(1000, 9999), strings.ToUpper(util.RandSeq(8)))
}
