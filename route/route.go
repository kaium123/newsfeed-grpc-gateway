// Application
//
// Application description
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /api
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package route

import (
	fileModuleRoute "newsfeed/modules/attachment/route"
	commentModuleRoute "newsfeed/modules/comment/route"
	postModuleRoute "newsfeed/modules/post/route"

	"net/http"
	reactModuleRoute "newsfeed/modules/reaction/route"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Setup() *gin.Engine {
	gin.SetMode(viper.GetString("GIN_MODE"))

	r := gin.New()
	setupCors(r)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api := r.Group("/api")
	postModuleRoute.PostSetup(api)
	reactModuleRoute.ReactSetup(api)
	commentModuleRoute.CommentSetup(api)
	fileModuleRoute.AttachmentSetup(api)
	return r
}

func setupCors(r *gin.Engine) {
	allowConf := viper.GetString("CORS_ALLOW_ORIGINS")
	if allowConf == "" {
		r.Use(cors.Default())
		return
	}
	allowSites := strings.Split(allowConf, ",")
	config := cors.DefaultConfig()
	config.AllowOrigins = allowSites
}
