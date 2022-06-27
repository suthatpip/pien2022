package middleware

import (
	"piennews/helper/apiErrors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if err := c.Errors.Last(); err != nil {
			if parseError := apiErrors.ParseError(err.Err); parseError != nil {
				c.JSON(parseError.Status, gin.H{
					"error_message": parseError.Message,
					"error_code":    parseError.Code,
				})
				return
			}
		}
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {

			if err := recover(); err != nil {
				serverError := apiErrors.ThrowError(apiErrors.ServerError)
				c.JSON(serverError.Status, gin.H{
					"error_message": serverError.Message,
					"error_code":    serverError.Code,
				})
				return
			}
		}()
		c.Next()
	}
}
