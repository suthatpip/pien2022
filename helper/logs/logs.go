package logs

import (
	"fmt"
	"os"

	"piennews/helper/config"
	"piennews/helper/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type LogParams struct {
	Begin       time.Time
	Context     *gin.Context
	Header      string
	Path        string
	Request     interface{}
	Response    interface{}
	Status      int
	Error       string
	Source      string
	Destination string
	Ssid        string
}

type AffLogs interface {
	Write()
}

type logFormatter struct{}

var logger = log.New()

func init() {

	logger.Level = log.DebugLevel
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(new(logFormatter))
}

func NewLogs(params *LogParams) AffLogs {
	return &LogParams{
		Begin:       params.Begin,
		Context:     params.Context,
		Path:        params.Path,
		Header:      params.Header,
		Request:     params.Request,
		Response:    params.Response,
		Status:      params.Status,
		Error:       params.Error,
		Source:      params.Source,
		Destination: params.Destination,
		Ssid:        params.Ssid,
	}
}

func (app *LogParams) Write() {

	logger.WithFields(log.Fields{
		"time":        app.Begin.Format("2006-01-02 15:04:05.000"),
		"source":      app.Source,
		"destination": app.Destination,
		"ssid":        app.Ssid,
		"status":      util.IfThenElse(app.Error != "", "ERROR", "SUCCESS"),
		"path":        app.Path,
		"statuscode":  app.Status,
		"hreq":        app.Header,
		"breq":        util.TruncateText(util.ToString(app.Request)),
		"bres":        util.TruncateText(util.ToString(app.Response)),
		"err":         util.SigleLine(app.Error),
		"elapsedtime": time.Since(app.Begin).String(),
	}).Info(config.GetENV().Owner)
}

func (s *logFormatter) Format(entry *log.Entry) ([]byte, error) {

	msg := fmt.Sprintf("time=\"%v\" level=%v source=\"%v\" destination=\"%v\" ssid=\"%v\" status=\"%v\" path=\"%v\" statuscode=%+v hreq=\"%v\" breq=\"%v\" bres=\"%+v\" err=\"%+v\" elapsedtime=\"%v\"\n",
		entry.Data["time"],
		strings.ToUpper(entry.Level.String()),
		entry.Data["source"],
		entry.Data["destination"],
		entry.Data["ssid"],
		entry.Data["status"],
		entry.Data["path"],
		entry.Data["statuscode"],
		entry.Data["hreq"],
		entry.Data["breq"],
		entry.Data["bres"],
		entry.Data["err"],
		entry.Data["elapsedtime"],
	)
	return []byte(msg), nil
}
