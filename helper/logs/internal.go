package logs

import (
	"fmt"
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
	Body    interface{}
	Error   string
}

type InternalLogs interface {
	WriteInternalLogs()
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
		Body:    params.Body,
		Error:   params.Error,
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
		"time":        app.Begin.Format("2006-01-02 15:04:05.000"),
		"uuid":        user_id,
		"status":      app.Context.Writer.Status(),
		"level":       util.IfThenElse(app.Error != "", "ERROR", "INFO"),
		"path":        app.Context.FullPath(),
		"params":      fmt.Sprintf("%+v", app.Context.Params),
		"body":        app.Body,
		"err":         util.SigleLine(app.Error),
		"elapsedtime": time.Since(app.Begin).String(),
	}).Info(config.GetENV().OWNER)
}

func (s *logInternalFormatter) Format(entry *log.Entry) ([]byte, error) {

	msg := fmt.Sprintf(`%v %v %v uuid=%v path="%v" params="%v" body="%v" err="%+v" %v`+"\n",
		entry.Data["time"],
		entry.Data["level"],
		entry.Data["status"],
		entry.Data["uuid"],
		entry.Data["path"],
		entry.Data["params"],
		entry.Data["body"],
		entry.Data["err"],
		entry.Data["elapsedtime"],
	)
	return []byte(msg), nil
}
