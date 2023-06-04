package route

import (
	"newsfeed/common/logger"
	"newsfeed/db"
	"newsfeed/modules/reaction/controller"
	"newsfeed/modules/reaction/repository"
	"newsfeed/modules/reaction/service"

	"github.com/gin-gonic/gin"
)

func ReactSetup(api *gin.RouterGroup) {
	db := db.NewEntDb()
	raventClient := logger.NewRavenClient()
	logger := logger.NewLogger(raventClient)
	repo := repository.NewReactionRepository(db, logger)
	service := service.NewReactionService(repo)
	ReactionController := controller.NewReactionController(service)

	Reaction := api.Group("/react")

	Reaction.GET("/", ReactionController.GetReaction)
	Reaction.POST("/create", ReactionController.CreatReact)
	Reaction.GET("/list", ReactionController.AllReaction)
	Reaction.POST("/update/:react_id", ReactionController.UpdateReaction)

}
