package route

import (
	"newsfeed/common/logger"
	"newsfeed/db"
	"newsfeed/modules/post/controller"
	"newsfeed/modules/post/repository"
	"newsfeed/modules/post/service"

	"github.com/gin-gonic/gin"
)

func PostSetup(api *gin.RouterGroup) {
	db := db.NewEntDb()
	raventClient := logger.NewRavenClient()
	logger := logger.NewLogger(raventClient)
	repo := repository.NewPostRepository(db, logger)
	service := service.NewPostService(repo)
	postController := controller.NewPostController(service)

	post := api.Group("/post")

	post.GET("/:post_id", postController.GetPost)
	post.POST("/create", postController.CreatePost)
	post.POST("/update/:post_id", postController.UpdatePost)
	post.GET("/list", postController.AllPost)
}
