package routers

import (
	"net/http"
	"piennews/controller"

	"github.com/gin-gonic/gin"
)

type templateCode struct {
	Code string `uri:"code" binding:"required"`
}

func Template(rg *gin.RouterGroup) {

	rg.POST("/:code", func(c *gin.Context) {
		var t templateCode
		if err := c.ShouldBindUri(&t); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"code":    http.StatusBadRequest,
				"pienweb": "/",
				"message": err.Error(),
			})
			return
		}
		controller.NewController().GetTemplate(c, t.Code)
	})
	rg.GET("/a", func(c *gin.Context) {

		c.HTML(http.StatusOK, "document.html", gin.H{
			"template": "A",
		})
	})
	rg.GET("/b", func(c *gin.Context) {

		c.HTML(http.StatusOK, "document.html", gin.H{
			"template": "B",
		})
	})
	rg.GET("/c", func(c *gin.Context) {

		c.HTML(http.StatusOK, "document.html", gin.H{
			"template": "C",
		})
	})

	// tmpl.Use(middleware.AuthorizeJWT())
	// {
	// 	tmpl.GET("/page/:type", func(c *gin.Context) {
	// 		var t TemplateType

	// 		if err := c.ShouldBindUri(&t); err != nil {
	// 			c.HTML(http.StatusBadRequest, "notification.html", gin.H{
	// 				"code":    http.StatusBadRequest,
	// 				"pienweb": "/",
	// 				"message": err.Error(),
	// 			})
	// 			return
	// 		}

	// 		v, err := c.Request.Cookie("auth")
	// 		u := goth.User{}

	// 		if err != nil {
	// 			c.HTML(http.StatusBadRequest, "notification.html", gin.H{
	// 				"code":    http.StatusBadRequest,
	// 				"pienweb": "/",
	// 				"message": err.Error(),
	// 			})
	// 			return
	// 		} else {
	// 			auth_user, err := service.UserControl().GetProfile(v.Value)

	// 			if err != nil {
	// 				c.HTML(http.StatusNotFound, "notification.html", gin.H{
	// 					"code":    http.StatusUnauthorized,
	// 					"pienweb": "/",
	// 					"message": err.Error(),
	// 				})
	// 			} else {
	// 				u = goth.User{
	// 					Name:      auth_user.Name,
	// 					AvatarURL: auth_user.Avatar,
	// 				}
	// 				isAdmin, _ := strconv.ParseBool(auth_user.IsAdmin)
	// 				c.HTML(http.StatusOK, "template.html", gin.H{
	// 					"draftTemplate": t.Type,
	// 					"template":      template.HTML(document.Template(t.Type)),
	// 					"title":         "Pien News",
	// 					"profileName":   u.Name,
	// 					"profileAvatar": u.AvatarURL,
	// 					"companyName":   "Pien corp.",
	// 					"isAuth":        true,
	// 					"isAdmin":       isAdmin,
	// 				})

	// 			}
	// 		}
	// 	})

	// 	tmpl.GET("/:type", func(c *gin.Context) {
	// 		var t TemplateType
	// 		if err := c.ShouldBindUri(&t); err != nil {

	// 			c.HTML(http.StatusBadRequest, "notification.html", gin.H{
	// 				"code":    http.StatusBadRequest,
	// 				"pienweb": "/",
	// 				"message": err.Error(),
	// 			})
	// 			return
	// 		}
	// 		file := fmt.Sprintf("%s.html", t.Type)
	// 		c.HTML(http.StatusOK, file, gin.H{})
	// 	})

	// 	tmpl.POST("/", func(c *gin.Context) {
	// 		var tc TemplateTypeCode

	// 		v, err := c.Request.Cookie("auth")

	// 		if err != nil {
	// 			c.HTML(http.StatusBadRequest, "notification.html", gin.H{
	// 				"code":    http.StatusBadRequest,
	// 				"pienweb": "/",
	// 				"message": err.Error(),
	// 			})
	// 			return
	// 		}

	// 		u, err := service.UserControl().GetProfile(v.Value)
	// 		if err != nil {
	// 			c.HTML(http.StatusBadRequest, "notification.html", gin.H{
	// 				"code":    http.StatusBadRequest,
	// 				"pienweb": "/",
	// 				"message": err.Error(),
	// 			})
	// 			return
	// 		}

	// 		if err := c.ShouldBindJSON(&tc); err != nil {
	// 			c.HTML(http.StatusBadRequest, "notification.html", gin.H{
	// 				"code":    http.StatusBadRequest,
	// 				"pienweb": "/",
	// 				"message": err.Error(),
	// 			})
	// 			return
	// 		}

	// 		prep := &service.PreProductModel{}
	// 		prep.UserID = u.UserID
	// 		prep.Product_type = service.ProductType.CUSTOM
	// 		prep.Product_html = document.Detail(tc.Code)
	// 		actualValue, actualFound := document.TemplateTitle(tc.Code)
	// 		if actualFound {
	// 			s := reflect.ValueOf(actualValue)
	// 			prep.Product_title = fmt.Sprintf("%s", s.Index(0))
	// 		}

	// 		service.PreProductControl().AddNew(prep)

	// 		c.JSON(http.StatusOK, gin.H{"template": document.Detail(tc.Code)})
	// 	})

	// }

	// tmpl.Use(middleware.AuthorizeJWTPOST())
	// {
	// 	tmpl.POST("/upload", func(c *gin.Context) {
	// 		form, err := c.MultipartForm()
	// 		if err != nil {

	// 			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	// 			return
	// 		}
	// 		v, err := c.Request.Cookie("auth")
	// 		if err != nil {

	// 			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	// 			return
	// 		} else {

	// 			auth_user, err := service.UserControl().GetProfile(v.Value)
	// 			if err != nil {
	// 				c.String(http.StatusBadRequest, fmt.Sprintf("StatusUnauthorized: %s", err.Error()))
	// 				return
	// 			} else {

	// 				files := form.File["file"]
	// 				path := "/assets/img/user/" + auth_user.UUID + "/uploads"
	// 				//path := "./uploads/" + auth_user.UserID
	// 				os.MkdirAll("."+path, os.ModePerm)
	// 				for _, file := range files {
	// 					extension := filepath.Ext(file.Filename)
	// 					newFileName := uuid.New().String() + extension

	// 					prep := &service.PreProductModel{}
	// 					prep.UserID = auth_user.UserID
	// 					prep.Product_type = service.ProductType.UPLOAD
	// 					prep.Product_file = path + "/" + newFileName
	// 					prep.Product_title = file.Filename

	// 					service.PreProductControl().AddNew(prep)

	// 					if err := c.SaveUploadedFile(file, "."+path+"/"+newFileName); err != nil {
	// 						c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	// 						return
	// 					} else {
	// 						c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields.", len(files)))
	// 						return
	// 					}
	// 				}
	// 			}
	// 		}
	// 	})
	// }
}
