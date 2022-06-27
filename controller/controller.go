package controller

import (
	"piennews/models"

	"github.com/gin-gonic/gin"
)

type controllerInterface interface {
	GetTemplate(c *gin.Context, code string)
	Verify(c *gin.Context, v *models.VerifyModel)
}

type controller struct {
}

func NewController() controllerInterface {
	return &controller{}
}
