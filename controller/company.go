package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"piennews/helper/apiErrors"
	"piennews/helper/jwt"
	"piennews/helper/util"
	"piennews/models"
	"piennews/services"

	"github.com/gin-gonic/gin"
)

func (ct *controller) CompanyList(c *gin.Context) {

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	company, err := services.NewService().GetCompanyList(uuid)
	if err != nil {
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))
		return
	}

	c.JSON(http.StatusOK, company)
}

// func (ct *controller) CompanyNew(c *gin.Context, com *models.CompanyModel) {

// 	h := c.MustGet("headers").(models.Header)
// 	uuid := jwt.ExtractClaims(h.Token, "uuid")

// 	com.UUID = uuid

// 	err := services.NewService().NewCompany(com)
// 	if err != nil {
// 		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "OK",
// 	})

// }

func (ct *controller) CompanyNewLogo(c *gin.Context, com *models.CompanyModel) {

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")
	com.UUID = uuid

	err := services.NewService().UpdateCompanyLogo(com)
	if err != nil {
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}

func (ct *controller) CompanyUpdate(c *gin.Context, com *models.CompanyModel) {
	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	com.UUID = uuid

	err := services.NewService().UpdateCompany(com)
	if err != nil {
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}

//
func (ct *controller) UploadCompanyLogo(c *gin.Context, company_code string) {
	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")
	form, err := c.MultipartForm()
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	files := form.File["file"]
	path := "/assets/upload/" + uuid + "/company"
	os.MkdirAll("."+path, os.ModePerm)
	for _, file := range files {
		extension := filepath.Ext(file.Filename)
		newFileName := util.GetUUID() + extension
		full_path := path + "/" + newFileName

		if err := c.SaveUploadedFile(file, "."+path+"/"+newFileName); err != nil {
			c.Status(http.StatusBadRequest)
			return
		} else {
			v := &models.CompanyModel{
				Logo: full_path,
				Code: company_code,
			}
			ct.CompanyNewLogo(c, v)
		}
	}

	c.Status(http.StatusOK)

}

func (ct *controller) UpdateCompany(c *gin.Context, company *models.CompanyModel) {
	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	company.UUID = uuid
	err := services.NewService().UpdateCompany(company)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}

func (ct *controller) DeleteCompany(c *gin.Context, company *models.CompanyModel) {
	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	company.UUID = uuid
	err := services.NewService().DeleteMyCompany(uuid, company.Code)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}

func (ct *controller) NewCompany(c *gin.Context, company *models.CompanyModel) {
	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	company.UUID = uuid
	err := services.NewService().NewCompany(company)

	if err != nil {
		c.Error(apiErrors.ThrowError(apiErrors.ServiceUnavailable, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

}
