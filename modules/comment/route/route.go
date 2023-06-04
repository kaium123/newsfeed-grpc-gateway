package route

import (
	"newsfeed/common/logger"
	"newsfeed/db"
	"newsfeed/modules/comment/controller"
	"newsfeed/modules/comment/repository"
	"newsfeed/modules/comment/service"

	"github.com/gin-gonic/gin"
)

func CommentSetup(api *gin.RouterGroup) {
	db := db.NewEntDb()
	raventClient := logger.NewRavenClient()
	logger := logger.NewLogger(raventClient)
	repo := repository.NewCommentRepository(db, logger)
	service := service.NewCommentService(repo)
	commentController := controller.NewCommentController(service)

	comment := api.Group("/comment")

	comment.GET("/:comment_id", commentController.GetComment)
	comment.POST("/create", commentController.Comment)
	comment.POST("/update/:comment_id", commentController.UpdateComment)
	comment.GET("/list", commentController.AllComment)
}
