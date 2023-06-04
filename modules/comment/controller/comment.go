package controller

import (
	"net/http"
	"newsfeed/common/logger"
	"newsfeed/errors"
	"newsfeed/modules/comment/models"
	"newsfeed/modules/comment/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CommentController struct {
	errors.GinError
	service service.CommentServiceInterface
}

func NewCommentController(service service.CommentServiceInterface) *CommentController {
	return &CommentController{service: service}
}

func (c CommentController) GetComment(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()

	Comment_id := ginContext.Param("comment_id")
	logger.LogError(Comment_id)
	id, err := strconv.Atoi(Comment_id)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	Comment, err := c.service.GetComment(id, ctx)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}
	ginContext.JSON(http.StatusCreated, gin.H{"message": "Comment", "Comment": Comment})

}

func (c CommentController) Comment(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	Comment := models.Comment{}

	if err := ginContext.ShouldBindBodyWith(&Comment, binding.JSON); err != nil {
		logger.LogError("JSON body binding error ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalidJSONBody"})
		return
	}

	CommentID, err := c.service.Comment(Comment, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "CommentCreated", "id": CommentID})
}

func (c CommentController) AllComment(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	Comment_id := ginContext.Query("post_id")
	logger.LogError(Comment_id)
	postID, err := strconv.Atoi(Comment_id)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	Comments, err := c.service.AllComment(postID, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "Comments", "Comments": Comments})
}

func (c CommentController) UpdateComment(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	Comment := models.Comment{}

	Comment_id := ginContext.Param("comment_id")
	logger.LogError(Comment_id)
	id, err := strconv.Atoi(Comment_id)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	if err := ginContext.ShouldBindBodyWith(&Comment, binding.JSON); err != nil {
		logger.LogError("JSON body binding error ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalidJSONBody"})
		return
	}

	CommentID, err := c.service.UpdateComment(id, Comment, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "CommentUpdated", "id": CommentID})
}
