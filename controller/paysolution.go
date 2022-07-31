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
	"time"

	"github.com/gin-gonic/gin"
)

func (ct *controller) ConfirmPayment(c *gin.Context, submit *models.SubmitPayment) {

	h := c.MustGet("headers").(models.Header)
	user_id := jwt.ExtractClaims(h.Token, "uuid")

	services.NewService().UpdateOrderStatus(submit.Payment_code, user_id, config.GetOrderStatus().APPROVED)

	pay, err := services.NewService().GetPaymentDetail(submit.Payment_code, user_id)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{})
		return
	}

	ref_no, err := services.NewService().NewInitPaysolution(submit.Payment_code)
	if err != nil {

		c.HTML(http.StatusOK, "error.html", gin.H{})
		return
	}

	services.NewService().UpdateOrderStatus(submit.Payment_code, user_id, config.GetOrderStatus().PENDING_PAYMENT)

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

func PaysolutionInquiry(ref_no string) {
	lg := &logs.LogExternalParams{}
	lg.Url = "PaySolution: Inquiry"
	lg.Request = ref_no
	services.NewService().EnquipryNextStep(ref_no, config.GetOrderStatus().VALIDATE_PAYMENT)

	defer func(begin time.Time) {
		lg.Begin = begin
		logs.ExternalLogs(lg).WriteExternalLogs()
	}(time.Now())

	inquiry, err := services.NewService().InquiryPaysolution(ref_no)
	if err != nil {
		lg.Error = err.Error()
		return
	}

	services.NewService().EnquipryNextStep(ref_no, config.GetOrderStatus().PENDING_PAYMENT)
	if inquiry.Status == "COMPLETE" {
		v, err := services.NewService().GetOrderPrice(ref_no)
		if err != nil {
			lg.Error = err.Error()
			return
		}
		if v == inquiry.Total {
			services.NewService().EnquipryNextStep(ref_no, config.GetOrderStatus().PAYMENT_COMPLETED)
			lg.Response = fmt.Sprintf("COMPLETE: %v == %v", inquiry.Total, v)
			return
		}
		lg.Response = fmt.Sprintf("INCOMPLETE: %v != %v", inquiry.Total, v)
	} else {
		lg.Response = util.SigleLine(fmt.Sprintf("INCOMPLETE: %+v", inquiry))
		lg.Error = inquiry.Status
	}
}
