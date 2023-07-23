package controller

import (
	"net/http"

	"newsfeed/common/logger"
	"newsfeed/errors"

	attachmentService "newsfeed/modules/attachment/service"

	"github.com/gin-gonic/gin"
)

type AttachmentController struct {
	errors.GinError
	service attachmentService.AttachmentUploaderServiceInterface
}

func NewAttachmentController(service attachmentService.AttachmentUploaderServiceInterface) *AttachmentController {
	return &AttachmentController{service: service}
}

func (c AttachmentController) UploadSingleFile(context *gin.Context) {
	file, err := context.FormFile("attachment")
	if err != nil {
		context.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		logger.LogInfo(err)
		return
	}

	fileName := context.Query("name")

	attachmentPath, upErr := c.service.UploadSingleAttachment(file, fileName)

	if upErr != nil {
		context.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		logger.LogInfo(upErr)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Comment", "attachment": attachmentPath})
}

func (c AttachmentController) GetSingleFile(context *gin.Context) {
	path := context.Query("attachment_path")
	c.service.GetSingleAttachment(context, path)
}

func (c AttachmentController) DeleteSingleFile(context *gin.Context) {
	path := context.Query("attachment_path")
	err := c.service.DeleteSingleAttachment(path)
	if err != nil {
		context.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		logger.LogInfo(err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Comment"})
}

func (c AttachmentController) GetSingleAttachmentFile(context *gin.Context) {
	path := "attachments/" + context.Param("path")
	c.service.GetSingleAttachment(context, path)
}
