package middleware

import (
	"net/http"
	"piennews/helper/jwt"
	"piennews/models"
	"piennews/services"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		v, err := c.Request.Cookie("auth")
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/auth")
			return
		}

		if err := jwt.Validate(v.Value); err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/auth")
			return
		}

		uuid := jwt.ExtractClaims(v.Value, "uuid")
		_, exist := services.NewService().GetCustomerWithUUID(uuid)
		if !exist {
			//c.AbortWithStatus(http.StatusUnauthorized)
			c.Redirect(http.StatusTemporaryRedirect, "/auth")
			return
		}

		c.Set("headers", models.Header{
			Token: v.Value,
		})

		c.Next()
	}
}
