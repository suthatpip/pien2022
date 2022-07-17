package routers

import (
	"fmt"
	"net/http"
	"piennews/controller"
	"piennews/models"

	"github.com/gin-gonic/gin"
)

func Paysolution(rg *gin.RouterGroup) {

	rg.POST("/callback", func(c *gin.Context) {
		paysolution := models.PaysolutionCallback{}
		if err := c.Bind(&paysolution); err != nil {
			fmt.Printf("%v\n", err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		controller.PaysolutionInquiry(paysolution.Ref_no)

		c.String(http.StatusOK, "OK")
	})

	rg.GET("/:code", func(c *gin.Context) {

		payment := models.SubmitPayment{}
		if err := c.ShouldBindUri(&payment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		controller.NewController().ConfirmPayment(c, &payment)

	})

}
