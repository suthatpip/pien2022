package controller

import (
	"html/template"
	"net/http"
	"piennews/controller/sidebar"
	"piennews/helper/apiErrors"
	"piennews/helper/jwt"
	"piennews/models"
	"piennews/services"

	"github.com/gin-gonic/gin"
)

func (ct *controller) Dashboard(c *gin.Context, status ...string) {

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	summary, err := services.NewService().Dashboard(uuid, status)
	if err != nil {
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))
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

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	summary, err := services.NewService().GetOrderDetail(payment_code, uuid)
	if err != nil {

		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))
		return
	}
	c.JSON(http.StatusOK, summary)
}
