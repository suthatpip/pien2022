package controller

import (
	"fmt"
	"net/http"
	"piennews/helper/config"
	"piennews/helper/jwt"
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
	logbody := ""
	logerror := ""

	defer func(begin time.Time) {
		logs.InternalLogs(&logs.LogInternalParams{
			Begin:   begin,
			Context: c,
			Body:    logbody,
			Error:   logerror,
		}).WriteInternalLogs()
	}(time.Now())

	payment_code, err := setPayment(v)
	if err != nil {
		logerror = err.Error()
		c.Status(http.StatusBadRequest)
		return
	}
	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")
	summary, err := services.NewService().GetPaymentDetail(payment_code, uuid)
	if err != nil {
		logerror = err.Error()
		c.Status(http.StatusServiceUnavailable)
		return
	}

	c.JSON(http.StatusOK, summary)
}

func (ct *controller) DeleteInitPayment(c *gin.Context, del *models.DeleteInitPayment) {
	logbody := ""
	logerror := ""

	defer func(begin time.Time) {
		logs.InternalLogs(&logs.LogInternalParams{
			Begin:   begin,
			Context: c,
			Body:    logbody,
			Error:   logerror,
		}).WriteInternalLogs()
	}(time.Now())

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	services.NewService().DeletePayment(del, uuid)

	c.Status(http.StatusOK)
}

func (ct *controller) DeleteInitAllPayment(c *gin.Context, del *models.DeleteInitPayment) {
	logbody := ""
	logerror := ""

	defer func(begin time.Time) {
		logs.InternalLogs(&logs.LogInternalParams{
			Begin:   begin,
			Context: c,
			Body:    logbody,
			Error:   logerror,
		}).WriteInternalLogs()
	}(time.Now())

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	services.NewService().DeleteProductAndPayment(del, uuid)

	c.Status(http.StatusOK)
}

func setPayment(v *models.InitPaymentModel) (string, error) {
	locale := currency.NewLocale("th")
	formatter := currency.NewFormatter(locale)
	formatter.MaxDigits = 2
	formatter.MinDigits = 2
	formatter.CurrencyDisplay = currency.DisplayNone

	payment_code := util.GetUUID()
	company_code := v.Company_Code
	customer_uuid := v.UUID
	start, _ := goment.New(v.Start_Date, "DD-MM-YYYY")
	end, _ := goment.New(v.End_Date, "DD-MM-YYYY")
	days := end.Diff(start, "days") + 1

	due_date, _ := goment.New(v.Start_Date, "DD-MM-YYYY")
	due_date = due_date.Subtract(1, "days")

	num_product := len(*v.Products)

	sub_total_baht, _ := currency.NewAmount(fmt.Sprintf("%v", float64(num_product)*config.Price*float64(days)), "THB")
	vat, _ := currency.NewAmount(fmt.Sprintf("%v", config.Price*float64(0)*float64(days)), "THB")
	total_baht, _ := currency.NewAmount(fmt.Sprintf("%v", float64(num_product)*config.Price*float64(days)), "THB")

	p := &models.AddPaymentModel{
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
	err = services.NewService().SubmitProduct(*v.Products, payment_code)
	if err != nil {
		return "", err
	}

	return payment_code, nil
}

func getOrderNo() string {
	d, _ := goment.New(time.Now().Format("02-01-2006"), "DD-MM-YYYY")
	return fmt.Sprintf("A%v-%v-%v ", d.Format("YY"), util.RandInt(1000, 9999), strings.ToUpper(util.RandSeq(4)))
}
