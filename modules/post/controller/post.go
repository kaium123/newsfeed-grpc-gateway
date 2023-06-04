package controller

import (
	"net/http"
	"newsfeed/common/logger"
	"newsfeed/errors"
	"newsfeed/modules/post/models"
	"newsfeed/modules/post/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PostController struct {
	errors.GinError
	service service.PostServiceInterface
}

func NewPostController(service service.PostServiceInterface) *PostController {
	return &PostController{service: service}
}

func (c PostController) GetPost(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()

	post_id := ginContext.Param("post_id")
	logger.LogError(post_id)
	id, err := strconv.Atoi(post_id)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	post, err := c.service.GetPost(id, ctx)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}
	ginContext.JSON(http.StatusCreated, gin.H{"message": "post","post":post})

}

func (c PostController) CreatePost(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	post := models.Post{}

	if err := ginContext.ShouldBindBodyWith(&post, binding.JSON); err != nil {
		logger.LogError("JSON body binding error ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalidJSONBody"})
		return
	}

	PostID, err := c.service.CreatePost(post, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "PostCreated", "id": PostID})
}

func (c PostController) AllPost(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()

	posts, _, err := c.service.AllPost(ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "posts", "posts": posts})
}

func (c PostController) UpdatePost(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	post := models.Post{}

	post_id := ginContext.Param("post_id")
	logger.LogError(post_id)
	id, err := strconv.Atoi(post_id)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	if err := ginContext.ShouldBindBodyWith(&post, binding.JSON); err != nil {
		logger.LogError("JSON body binding error ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalidJSONBody"})
		return
	}

	PostID, err := c.service.UpdatePost(id, post, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "PostUpdated", "id": PostID})
}
