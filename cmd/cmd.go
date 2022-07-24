package cmd

import (
	"fmt"
	"net/http"
	"piennews/helper/jwt"
	routers "piennews/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.HTML(http.StatusBadRequest, "404.html", gin.H{
			"code":    http.StatusInternalServerError,
			"pienweb": "/",
			"message": "",
		})
	}))

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusBadRequest, "404.html", gin.H{
			"code":    http.StatusBadRequest,
			"pienweb": "/",
			"message": "",
		})
	})
	// router.Use(middleware.ErrorHandler())
	// router.Use(middleware.AuthHandler())

	router.Static("/assets", "./assets")

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.LoadHTMLGlob("pages/*/*")

	router.GET("/jwt-dev", func(c *gin.Context) {
		maxAge := 86400 * 30
		jwtcookie, _ := jwt.Generate("65eeda46-fedb-4c39-8b93-44cb12d6b3ef")

		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "auth",
			Value:  jwtcookie,
			MaxAge: maxAge,
			Path:   "/",
		})
	})

	router.GET("/passcode", func(c *gin.Context) {
		c.HTML(http.StatusOK, "passcode.html", gin.H{
			"title": "passcode",
		})

	})

	routers.Gateway(router)
	router.Run(":8080")
}
