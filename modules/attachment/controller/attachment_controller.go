package controller

import (
	"net/http"

	"newsfeed/common/logger"
	"newsfeed/errors"

	"newsfeed/modules/attachment/models"
	attachmentService "newsfeed/modules/attachment/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AttachmentController struct {
	errors.GinError
	service attachmentService.AttachmentUploaderServiceInterface
}

func NewAttachmentController(service attachmentService.AttachmentUploaderServiceInterface) *AttachmentController {
	return &AttachmentController{service: service}
}

func (c AttachmentController) UploadAttachments(context *gin.Context) {
	attachments := models.Attachments{}
	if err := context.ShouldBindBodyWith(&attachments, binding.JSON); err != nil {
		logger.LogError("JSON body binding error ", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalidJSONBody"})
		return
	}

	_, upErr := c.service.UploadAttachments(attachments)
	if upErr != nil {
		context.AbortWithStatusJSON(c.GetStatusCode(upErr), gin.H{"error": c.ErrorTraverse(upErr)})
		logger.LogInfo(upErr)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "attachments created"})
}

func (c AttachmentController) GetSingleFile(context *gin.Context) {
	path := context.Query("attachment_path")
	c.service.GetSingleAttachment(context, path)
}

func (c AttachmentController) Delete(context *gin.Context) {
	path := context.Query("attachment_path")
	err := c.service.Delete(path)
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
