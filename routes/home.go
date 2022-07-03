package routers

import (
	"net/http"
	"piennews/helper/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/line"
)

func Home(rg *gin.RouterGroup) {
	// //message := rg.Group("/message")
	key := "pien"        // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	url := config.GetENV().Environment

	goth.UseProviders(
		// google.New("48856069706-ujhkqk66uge342bbiqrhcok3akg47chp.apps.googleusercontent.com", "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", url+"/auth/google/callback", "email", "profile"),
		google.New("254948135185-4ci1pkhondtg99eknek03fi3rf113b31.apps.googleusercontent.com", "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", url+"/auth/google/callback", "email", "profile"),
		facebook.New("936226287272219", "e92340c81631c29a1855e13a501395e3", url+"/auth/facebook/callback"),
		line.New("1656602577", "86d81b7bc1d7529a239f494c72e9f480", url+"/auth/line/callback", "profile", "openid", "email"),
	)

	rg.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "support",
		})

		// highlight4, _ := service.PageControl().Highlight(10, service.DropPosition.POSITION_LETTER, service.DropPosition.POSITION_HEAD_LINE)
		// highlight2, _ := service.PageControl().Highlight(1, service.DropPosition.POSITION_SEMI)
		// highlight1, _ := service.PageControl().Highlight(1, service.DropPosition.POSITION_SINGLE)

		// advertised, next, _ := service.PageControl().Advertise("")
		// v, err := c.Request.Cookie("auth")

		// u := goth.User{}
		// if err != nil {
		// 	http.SetCookie(c.Writer, &http.Cookie{
		// 		Name:  "ids_cv",
		// 		Value: next,
		// 		Path:  "/",
		// 	})
		// 	c.HTML(http.StatusOK, "index.html", gin.H{
		// 		"title":         "Pien News",
		// 		"profileName":   "Guest",
		// 		"profileAvatar": "/assets/img/user/guest32.png",
		// 		"isAuth":        false,
		// 		"isAdmin":       false,
		// 		"highlight4":    highlight4,
		// 		"highlight2":    highlight2,
		// 		"highlight1":    highlight1,
		// 		"advertised":    advertised,
		// 		"newsdate":      fmt.Sprintf("ประจำวันที่ %v", util.DateTH(time.Now().Format("01-02-2006"))),
		// 	})
		// 	return
		// } else {
		// 	if middleware.VerifyJWT(v.Value) {

		// 		auth_user, err := service.UserControl().GetProfile(v.Value)

		// 		if err != nil {
		// 			http.SetCookie(c.Writer, &http.Cookie{
		// 				Name:  "ids_cv",
		// 				Value: next,
		// 				Path:  "/",
		// 			})
		// 			c.HTML(http.StatusOK, "index.html", gin.H{
		// 				"title":         "Pien News",
		// 				"profileName":   "Guest",
		// 				"profileAvatar": "/assets/img/user/guest32.png",
		// 				"isAuth":        false,
		// 				"isAdmin":       false,
		// 				"highlight4":    highlight4,
		// 				"highlight2":    highlight2,
		// 				"highlight1":    highlight1,
		// 				"advertised":    advertised,
		// 				"newsdate":      fmt.Sprintf("ประจำวันที่ %v", util.DateTH(time.Now().Format("01-02-2006"))),
		// 			})

		// 			return
		// 		} else {
		// 			u = goth.User{
		// 				Name:      auth_user.Name,
		// 				AvatarURL: auth_user.Avatar,
		// 			}
		// 			http.SetCookie(c.Writer, &http.Cookie{
		// 				Name:  "ids_cv",
		// 				Value: next,
		// 				Path:  "/",
		// 			})

		// 			isadmin, _ := strconv.ParseBool(auth_user.IsAdmin)
		// 			c.HTML(http.StatusOK, "index.html", gin.H{
		// 				"title":         "Pien News",
		// 				"profileName":   u.Name,
		// 				"profileAvatar": u.AvatarURL,
		// 				"isAuth":        true,
		// 				"isAdmin":       isadmin,
		// 				"highlight4":    highlight4,
		// 				"highlight2":    highlight2,
		// 				"highlight1":    highlight1,
		// 				"advertised":    advertised,
		// 				"newsdate":      fmt.Sprintf("ประจำวันที่ %v", util.DateTH(time.Now().Format("01-02-2006"))),
		// 			})
		// 		}
		// 	}
		// }
	})

	rg.GET("/support", func(c *gin.Context) {
		c.HTML(http.StatusOK, "support.html", gin.H{
			"title": "support",
		})
	})

}
