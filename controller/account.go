package controller

import (
	"html/template"
	"net/http"
	"piennews/controller/sidebar"
	"piennews/helper/jwt"
	"piennews/models"
	"piennews/services"

	"github.com/gin-gonic/gin"
)

func (ct *controller) Account(c *gin.Context) {

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	customer, exist := services.NewService().GetCustomerWithUUID(uuid)
	name, profile := sidebar.GetUserSidebar(customer, exist)
	c.HTML(http.StatusOK, "account.html", gin.H{
		"customer": gin.H{
			"name":    template.HTML(name),
			"profile": profile,
		},
	})

}
