package routers

import (
	"net/http"
	"os"
	"path/filepath"
	"piennews/controller"
	"piennews/helper/util"
	"piennews/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Company(rg *gin.RouterGroup) {

	rg.POST("/list", func(c *gin.Context) {
		controller.NewController().CompanyList(c)
	})

	rg.POST("/init", func(c *gin.Context) {
		code := util.GetUUID()
		c.JSON(http.StatusOK, gin.H{"code": code})
	})

	rg.POST("/new", func(c *gin.Context) {

		v := &models.CompanyModel{}
		if err := c.ShouldBindJSON(&v); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		controller.NewController().CompanyNew(c, v)
	})

	rg.POST("/logo", func(c *gin.Context) {
		// fmt.Println("/logo/:code")
		// var company models.CompanyModel
		// if err := c.ShouldBindUri(&company); err != nil {
		// 	c.Status(http.StatusBadRequest)
		// 	return
		// }

		form, err := c.MultipartForm()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		v, err := c.Request.Cookie("code")

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		} else {
			files := form.File["file"]
			path := "/assets/upload/company/" + v.Value
			os.MkdirAll("."+path, os.ModePerm)
			for _, file := range files {
				extension := filepath.Ext(file.Filename)
				newFileName := uuid.New().String() + extension
				full_path := path + "/" + newFileName

				if err := c.SaveUploadedFile(file, "."+path+"/"+newFileName); err != nil {
					c.Status(http.StatusBadRequest)
					return
				} else {
					v := &models.CompanyModel{
						Logo: full_path,
						Code: v.Value,
					}
					controller.NewController().CompanyNewLogo(c, v)
				}
			}
		}

		c.Status(http.StatusOK)
	})

}
