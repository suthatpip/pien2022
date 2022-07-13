package controller

import (
	"fmt"
	"net/http"
	"piennews/helper/logs"
	"piennews/services"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

func (ct *controller) GetTemplate(c *gin.Context, code string) {
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
