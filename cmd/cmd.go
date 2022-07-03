package cmd

import (
	"fmt"
	"net/http"
	"piennews/helper/jwt"
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

	router.GET("/jwt-dev", func(c *gin.Context) {
		maxAge := 86400 * 30
		jwtcookie, _ := jwt.Generate("e5c1d98e-7f97-4284-9f2d-17022583020e")
		fmt.Printf("cookie = %v\n", jwtcookie)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "auth",
			Value:  jwtcookie,
			MaxAge: maxAge,
			Path:   "/",
		})
	})

	routers.Gateway(router)
	router.Run(":8080")
}
