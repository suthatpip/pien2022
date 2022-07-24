package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"piennews/controller/sidebar"
	"piennews/helper/jwt"
	"piennews/helper/logs"
	"piennews/helper/util"
	"piennews/models"
	"piennews/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ct *controller) Product(c *gin.Context) {

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")

	customer, exist := services.NewService().GetCustomerWithUUID(uuid)
	name, profile := sidebar.GetUserSidebar(customer, exist)
	c.HTML(http.StatusOK, "product.html", gin.H{
		"customer": gin.H{
			"name":    template.HTML(name),
			"profile": profile,
		},
	})

}

func (ct *controller) CustomeFile(c *gin.Context, product *models.ProductModel) {
	logbody := ""
	logerror := ""

	defer func(begin time.Time) {
		logs.InternalLogs(&logs.LogInternalParams{
			Begin:   begin,
			Context: c,
			Body:    logbody,
			Error:   logerror,
		}).WriteInternalLogs()
	}(time.Now())

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")
	if err := services.NewService().NewProduct(&models.ProductModel{
		Product_Code:   product.Product_Code,
		Product_Name:   product.Product_Name,
		Product_Size:   "-",
		Product_Detail: product.Product_Detail,
		Product_Type:   "template",
		Template_code:  product.Template_code,
	}, uuid); err != nil {
		logerror = err.Error()
		c.Status(http.StatusServiceUnavailable)
		return
	}

	c.Status(http.StatusOK)

}

func (ct *controller) UploadFile(c *gin.Context) {
	logbody := ""
	logerror := ""

	defer func(begin time.Time) {
		logs.InternalLogs(&logs.LogInternalParams{
			Begin:   begin,
			Context: c,
			Body:    logbody,
			Error:   logerror,
		}).WriteInternalLogs()
	}(time.Now())

	form, err := c.MultipartForm()

	if err != nil {
		logerror = err.Error()
		c.Status(http.StatusBadRequest)
		return
	}
	h := c.MustGet("headers").(models.Header)
	user_id := jwt.ExtractClaims(h.Token, "uuid")
	files := form.File["file"]
	path := "/assets/upload/" + user_id
	os.MkdirAll("."+path, os.ModePerm)
	for _, file := range files {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		full_path := path + "/" + newFileName
		logbody = fmt.Sprintf("%v -> %v", file.Filename, full_path)
		if err := c.SaveUploadedFile(file, "."+path+"/"+newFileName); err != nil {
			c.Status(http.StatusFailedDependency)
			return
		}
		if err := services.NewService().NewProduct(&models.ProductModel{
			Product_Code:   util.GetUUID(),
			Product_Name:   file.Filename,
			Product_Size:   fmt.Sprintf("%v", file.Size),
			Product_Detail: full_path,
			Product_Type:   "file",
		}, user_id); err != nil {
			logerror = err.Error()
			c.Status(http.StatusServiceUnavailable)
			return
		}
	}
	c.Status(http.StatusOK)
}

func (ct *controller) GetProduct(c *gin.Context) {
	logbody := ""
	logerror := ""

	defer func(begin time.Time) {
		logs.InternalLogs(&logs.LogInternalParams{
			Begin:   begin,
			Context: c,
			Body:    logbody,
			Error:   logerror,
		}).WriteInternalLogs()
	}(time.Now())

	h := c.MustGet("headers").(models.Header)
	user_id := jwt.ExtractClaims(h.Token, "uuid")

	files, err := services.NewService().GetProduct(user_id)
	if err != nil {
		logerror = err.Error()
		c.Status(http.StatusServiceUnavailable)
		return
	}

	c.JSON(http.StatusOK, files)
}

func (ct *controller) DeleteProduct(c *gin.Context, products *models.ProductsModel) {
	logbody := ""
	logerror := ""

	defer func(begin time.Time) {
		logs.InternalLogs(&logs.LogInternalParams{
			Begin:   begin,
			Context: c,
			Body:    logbody,
			Error:   logerror,
		}).WriteInternalLogs()
	}(time.Now())

	h := c.MustGet("headers").(models.Header)
	uuid := jwt.ExtractClaims(h.Token, "uuid")
	for _, v := range products.Products {
		fmt.Printf("%v %v\n", &v, uuid)
		services.NewService().DelProduct(&v, uuid)
	}

	c.Status(http.StatusOK)

}
