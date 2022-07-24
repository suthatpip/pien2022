package controller

import (
	"html/template"
	"net/http"
	"piennews/controller/sidebar"
	"piennews/helper/jwt"
	"piennews/helper/logs"
	"piennews/models"
	"piennews/services"
	"time"

	"github.com/gin-gonic/gin"
)

func (ct *controller) Dashboard(c *gin.Context, status ...string) {
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

	summary, err := services.NewService().Dashboard(uuid, status)
	if err != nil {
		logerror = err.Error()
		c.Status(http.StatusServiceUnavailable)
		return
	}

	customer, exist := services.NewService().GetCustomerWithUUID(uuid)
	name, profile := sidebar.GetUserSidebar(customer, exist)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"pending_payment_total": summary.Summary.PENDING_PAYMENT,
		"on_process_total":      summary.Summary.ON_PROCESS,
		"publish_total":         summary.Summary.PUBLISH,
		"order_total":           summary.Summary.ALL,
		"orders":                summary.Orders,
		"customer": gin.H{
			"name":    template.HTML(name),
			"profile": profile,
		},
	})

}

func (ct *controller) DashboardOrderDetail(c *gin.Context, payment_code string) {
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

	summary, err := services.NewService().GetOrderDetail(payment_code, uuid)
	if err != nil {
		logerror = err.Error()
		c.Status(http.StatusServiceUnavailable)
		return
	}
	c.JSON(http.StatusOK, summary)
}
