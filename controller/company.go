package controller

import (
	"net/http"
	"piennews/helper/jwt"
	"piennews/helper/logs"
	"piennews/models"
	"piennews/services"

	"time"

	"github.com/gin-gonic/gin"
)

func (ct *controller) CompanyList(c *gin.Context) {
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

	company, err := services.NewService().GetCompanyList(uuid)
	if err != nil {

		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, company)
}

func (ct *controller) CompanyNew(c *gin.Context, com *models.CompanyModel) {
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

	com.UUID = uuid

	is_success := services.NewService().SaveCompany(com)
	if is_success {
		c.Status(http.StatusOK)
		return
	}
	c.Status(http.StatusBadRequest)

}

func (ct *controller) CompanyNewLogo(c *gin.Context, com *models.CompanyModel) {
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
	com.UUID = uuid

	is_success := services.NewService().SaveCompanyLogo(com)
	if is_success {
		c.Status(http.StatusOK)
		return
	}
	c.Status(http.StatusBadRequest)

}
