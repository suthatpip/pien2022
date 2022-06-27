package cmd

import (
	"net/http"
	routers "piennews/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()
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
	routers.Gateway(router)
	router.Run(":8080")
}
