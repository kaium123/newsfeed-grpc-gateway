package controller

import (
	"net/http"
	"newsfeed/common/logger"
	"newsfeed/errors"
	"newsfeed/modules/reaction/models"
	"newsfeed/modules/reaction/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ReactionController struct {
	errors.GinError
	service service.ReactionServiceInterface
}

func NewReactionController(service service.ReactionServiceInterface) *ReactionController {
	return &ReactionController{service: service}
}

func (c ReactionController) GetReaction(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	react_type := ginContext.Query("react_type")
	postID := ginContext.Query("post_id")
	pID, err := strconv.Atoi(postID)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	postType := ginContext.Query("post_type")

	Reaction, err := c.service.GetReaction(pID, postType, react_type, ctx)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}
	ginContext.JSON(http.StatusCreated, gin.H{"message": "react", "Reaction": Reaction})

}

func (c ReactionController) CreatReact(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	Reaction := models.React{}

	if err := ginContext.ShouldBindBodyWith(&Reaction, binding.JSON); err != nil {
		logger.LogError("JSON body binding error ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalidJSONBody"})
		return
	}
	logger.LogError(Reaction)

	ReactionID, err := c.service.CreateReact(Reaction, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "ReacCreated", "id": ReactionID})
}

func (c ReactionController) AllReaction(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	postID := ginContext.Query("post_id")
	pID, err := strconv.Atoi(postID)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	postType := ginContext.Query("post_type")

	reacts, err := c.service.AllReaction(pID, postType, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "reacts", "reacts": reacts})
}

func (c ReactionController) UpdateReaction(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	react := models.React{}
	react_id := ginContext.Param("react_id")
	logger.LogError(react_id)
	id, err := strconv.Atoi(react_id)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	if err := ginContext.ShouldBindBodyWith(&react, binding.JSON); err != nil {
		logger.LogError("JSON body binding error ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalidJSONBody"})
		return
	}

	ReactionID, err := c.service.UpdateReaction(id, react, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "reactsUpdated", "id": ReactionID})
}
