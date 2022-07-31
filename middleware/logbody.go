package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"piennews/helper/util"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogMiddleware(t time.Time) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		req := util.GetUUID()
		c.Request.Header.Add("x-request-id", req)

		c.Next()

		if accept(c.ContentType()) {
			fmt.Printf("%v %v %v	uri: %v, method: %v, req-body: %v, resp code: %v , res-body: %v elapsedtime: %v %v\n",
				t.Format("2006-01-02 15:04:05.000"),
				req,
				util.IfThenElse(c.Writer.Status() == 200, "INFO", "ERROR"),
				c.Request.RequestURI,
				c.Request.Method,
				util.SigleLine(string(bodyBytes)),
				c.Writer.Status(),
				blw.body.String(),
				time.Since(t).String(),
				time.Now().Format("2006-01-02 15:04:05.000"),
			)
		}

	}

}

func accept(contentType string) bool {
	switch contentType {
	case "application/json":
		return true
	}
	return false
}
