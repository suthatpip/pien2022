package logs

import (
	"fmt"
	"os"

	"piennews/helper/config"
	"piennews/helper/util"
	"time"

	log "github.com/sirupsen/logrus"
)

var externalLog = log.New()

type LogExternalParams struct {
	Begin    time.Time
	Body     interface{}
	Url      string
	Request  interface{}
	Response interface{}
	Error    interface{}
	Status   string
	Level    string
}

type ExternalLogs interface {
	WriteExternalLogs()
}

type logExternalFormatter struct{}

var loggerExternal = log.New()

func init() {

	externalLog.Level = log.DebugLevel
	externalLog.SetOutput(os.Stdout)
	externalLog.SetFormatter(new(logExternalFormatter))
}

func NewExternalLogs(params *LogExternalParams) ExternalLogs {
	return &LogExternalParams{
		Begin: params.Begin,
		Body:  params.Body,
		Error: params.Error,
	}
}

func (app *LogExternalParams) WriteExternalLogs() {

	externalLog.WithFields(log.Fields{
		"time":        app.Begin.Format("2006-01-02 15:04:05.000"),
		"status":      app.Status,
		"level":       util.IfThenElse(app.Error != "", "ERROR", "INFO"),
		"url":         app.Url,
		"req":         fmt.Sprintf("%+v", app.Request),
		"res":         fmt.Sprintf("%+v", app.Response),
		"err":         util.SigleLine(fmt.Sprintf("%+v", app.Error)),
		"elapsedtime": time.Since(app.Begin).String(),
	}).Info(config.GetENV().OWNER)
}

func (s *logExternalFormatter) Format(entry *log.Entry) ([]byte, error) {

	msg := fmt.Sprintf(`%v %v url="%v" req="%v" res="%v" err="%+v" %v`+"\n",
		entry.Data["time"],
		entry.Data["level"],
		entry.Data["url"],
		entry.Data["req"],
		entry.Data["res"],
		entry.Data["err"],
		entry.Data["elapsedtime"],
	)
	return []byte(msg), nil
}
