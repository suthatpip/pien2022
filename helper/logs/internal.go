package logs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"piennews/helper/config"
	"piennews/helper/jwt"
	"piennews/helper/util"
	"piennews/models"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var internalLog = log.New()

type LogInternalParams struct {
	Begin   time.Time
	Context *gin.Context

	Error string
}

type InternalLogs interface {
	WriteInternalLogs()
	WriteInternalLogs2()
}

type logInternalFormatter struct{}

func init() {

	internalLog.Level = log.DebugLevel
	internalLog.SetOutput(os.Stdout)
	internalLog.SetFormatter(new(logInternalFormatter))
}

func NewInternalLogs(params *LogInternalParams) InternalLogs {
	return &LogInternalParams{
		Begin:   params.Begin,
		Context: params.Context,

		Error: params.Error,
	}
}

func (app *LogInternalParams) WriteInternalLogs() {
	v, exits := app.Context.Get("headers")
	user_id := ""
	if exits {
		user_id = jwt.ExtractClaims(v.(models.Header).Token, "uuid")
	} else {
		user_id = app.Context.Query("auth")
	}

	internalLog.WithFields(log.Fields{
		"time":   app.Begin.Format("2006-01-02 15:04:05.000"),
		"uuid":   user_id,
		"status": app.Context.Writer.Status(),
		"level":  util.IfThenElse(app.Error != "", "ERROR", "INFO"),
		"path":   app.Context.FullPath(),
		"params": fmt.Sprintf("%+v", app.Context.Params),

		"err":         util.SigleLine(app.Error),
		"elapsedtime": time.Since(app.Begin).String(),
	}).Info(config.GetENV().OWNER)
}

func (app *LogInternalParams) WriteInternalLogs2() {
	req := app.Context.GetHeader("x-request-id")

	internalLog.WithFields(log.Fields{
		"time": app.Begin.Format("2006-01-02 15:04:05.000"),
		"req":  req,
		"err":  util.SigleLine(app.Error),
	}).Info(config.GetENV().OWNER)

}

func headerToString(c *gin.Context) string {
	return fmt.Sprintf("%+v", c.Request.Header)
}

func (s *logInternalFormatter) Format(entry *log.Entry) ([]byte, error) {
	msg := fmt.Sprintf("%v %v ERROR	%v\n",
		entry.Data["time"],
		entry.Data["req"],
		entry.Data["err"],
	)
	return []byte(msg), nil
}

// ///////////////////////////////////////////////////////////
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func body(c *gin.Context) string {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	c.Next()
	fmt.Printf("%+v\n", bodyBytes)
	return string(bodyBytes)

}
