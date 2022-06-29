package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BodyToString(c *gin.Context) string {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	body := SigleLine(reqBody)

	return fmt.Sprintf("%+v", body)
}
