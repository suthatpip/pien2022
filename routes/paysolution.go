package routers

import (
	"net/http"
	"piennews/controller"
	"piennews/models"

	"github.com/gin-gonic/gin"
)

func Paysolution(rg *gin.RouterGroup) {

	rg.POST("/callback", func(c *gin.Context) {
		// var callback PaysolutionCallback
		// if err := c.Bind(&callback); err != nil {
		// 	fmt.Printf("%v\n", err.Error())
		// 	c.String(http.StatusBadRequest, err.Error())
		// 	return
		// }

		// paysolutionModel := &service.PaysolutionModel{}
		// paysolutionModel.Ref_no = callback.Ref_no
		// paysolutionModel.Merchant_id = callback.Merchant_id
		// paysolutionModel.Customere_mail = callback.Customere_mail
		// paysolutionModel.Product_detail = callback.Product_detail
		// paysolutionModel.Total = callback.Total
		// paysolutionModel.Card_type = callback.Card_type
		// err := service.PaysolutionControl().PaysolutionCallback(paysolutionModel)

		// if err != nil {
		// 	c.String(http.StatusBadRequest, "Error")
		// 	return
		// }
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
