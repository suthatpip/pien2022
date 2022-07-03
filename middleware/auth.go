package middleware

import (
	"net/http"
	"piennews/helper/jwt"
	"piennews/models"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		v, err := c.Request.Cookie("auth")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err := jwt.Validate(v.Value); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("headers", models.Header{
			Token: v.Value,
		})

		c.Next()
	}
}
