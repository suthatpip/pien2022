package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"piennews/controller/sidebar"
	"piennews/helper/jwt"
	"piennews/helper/util"
	"piennews/models"
	"piennews/services"
	"reflect"

	"github.com/gin-gonic/gin"
)

func (ct *controller) Template(c *gin.Context, code string) {
	// logbody := ""
	// logerror := ""

	// defer func(begin time.Time) {
	// 	logs.InternalLogs(&logs.LogInternalParams{
	// 		Begin:   begin,
	// 		Context: c,
	// 		Body:    logbody,
	// 		Error:   logerror,
	// 	}).WriteInternalLogs()
	// }(time.Now())

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	customer, exist := services.NewService().GetCustomerWithUUID(uuid)

	name, profile := sidebar.GetUserSidebar(customer, exist)

	c.HTML(http.StatusOK, "template.html", gin.H{
		"template":    code,
		"document_no": util.GetUUID(),
		"customer": gin.H{
			"name":    template.HTML(name),
			"profile": profile,
		},
	})

}

func (ct *controller) GetTemplate(c *gin.Context, code string) {
	// logbody := ""
	// logerror := ""

	// defer func(begin time.Time) {
	// 	logs.InternalLogs(&logs.LogInternalParams{
	// 		Begin:   begin,
	// 		Context: c,
	// 		Body:    logbody,
	// 		Error:   logerror,
	// 	}).WriteInternalLogs()
	// }(time.Now())

	template, found := services.NewService().GetTemplate(code)
	if !found {
		c.Status(http.StatusNotFound)
		return
	} else {
		s := reflect.ValueOf(template)
		c.JSON(http.StatusOK, gin.H{
			"title":    fmt.Sprintf("%s", s.Index(0)),
			"subtitle": fmt.Sprintf("%s", s.Index(1)),
			"detail":   fmt.Sprintf("%s", s.Index(2)),
		})
	}

}
