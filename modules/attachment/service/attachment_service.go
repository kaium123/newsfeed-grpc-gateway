package service

import (
	"context"
	"mime/multipart"
	"net/http"
	"newsfeed/common/utils"
	"newsfeed/ent"
	fileUploader "newsfeed/modules/attachment/file_uploader"
	"newsfeed/modules/attachment/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

type AttachmentUploaderServiceInterface interface {
	UploadSingleAttachment(fileHeader *multipart.FileHeader, name string) (*ent.Attachment, error)
	GetSingleAttachment(context *gin.Context, attachmentPath string)
	DeleteSingleAttachment(attachmentPath string) error
}

type AttachmentUploaderService struct {
	uploader             fileUploader.FileUploaderInterface
	attachmentRepository repository.AttachmentRepositoryInterface
}

func NewAttachmentService(attachmentRepository repository.AttachmentRepositoryInterface, uploader fileUploader.FileUploaderInterface) AttachmentUploaderServiceInterface {
	service := &AttachmentUploaderService{
		uploader:             uploader,
		attachmentRepository: attachmentRepository,
	}
	return service
}

func (aus AttachmentUploaderService) DeleteSingleAttachment(attachmentPath string) error {
	return aus.uploader.DeleteSingleFile(attachmentPath)
}

func (aus AttachmentUploaderService) GetSingleAttachment(context *gin.Context, attachmentPath string) {
	attachment, err := aus.attachmentRepository.GetByPath(context, attachmentPath)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "fileNotFound", "status": "error"})
		return
	}
	aus.uploader.GetSingleFile(context, attachmentPath, attachment.Name)
}

func (aus AttachmentUploaderService) UploadSingleAttachment(fileHeader *multipart.FileHeader, name string) (*ent.Attachment, error) {
	parentDirectory := "attachments"
	attachmentPath := parentDirectory + "/" + utils.GetFileHashName(fileHeader.Filename)

	err := aus.uploader.UploadSingleFile(parentDirectory, attachmentPath, fileHeader)
	if err != nil {
		return nil, err
	}

	pathSegregation := strings.Split(attachmentPath, "/")
	fileName := pathSegregation[1]

	if strings.Compare(name, "") == 0 {
		nameWithExtension := strings.Split(fileName, ".")
		nameWithoutExtension := nameWithExtension[0]
		name = nameWithoutExtension
	}

	attachment, err := aus.attachmentRepository.Store(context.Background(), attachmentPath, name)

	// Attachment Path modification
	attachment.Path = "/attachment/single/" + fileName
	if strings.Compare(attachment.Name, "") == 0 {
		attachment.Name = fileName
	}

	return attachment, nil
}
