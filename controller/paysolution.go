package controller

import (
	"fmt"
	"net/http"
	"piennews/helper/jwt"
	"piennews/helper/logs"
	"piennews/models"
	"piennews/services"
	"time"

	"github.com/gin-gonic/gin"
)

func (ct *controller) ConfirmPayment(c *gin.Context, submit *models.SubmitPayment) {
	logbody := ""
	logerror := ""

	defer func(begin time.Time) {
		logs.NewLogs(&logs.LogParams{
			Begin:   begin,
			Context: c,
			Body:    logbody,
			Error:   logerror,
		}).Write()
	}(time.Now())

	h := c.MustGet("headers").(models.Header)
	user_id := jwt.ExtractClaims(h.Token, "uuid")

	pay, err := services.NewService().GetPaymentDetail(submit.Payment_code, user_id)
	if err != nil {
		logerror = err.Error()
		c.HTML(http.StatusOK, "error.html", gin.H{})
		return
	}

	ref_no, err := services.NewService().NewInitPaysolution(submit.Payment_code)
	if err != nil {
		logerror = err.Error()
		c.HTML(http.StatusOK, "error.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "paysolution.html", gin.H{
		"paysolution_url":           "https://www.thaiepay.com/epaylink/payment.aspx",
		"paysolution_refno":         fmt.Sprintf("%v", ref_no),
		"paysolution_merchantid":    "39233015",
		"paysolution_customeremail": "admin@test.com",
		"paysolution_cc":            "00",
		"paysolution_productdetail": fmt.Sprintf("ประกาศหนังสือพิมพ์ (%v-%v)", pay.Publish_Start_Date, pay.Publish_End_Date),
		"paysolution_total":         pay.Total,
	})

}
