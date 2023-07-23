package service

import (
	"context"
	"newsfeed/common/logger"
	"newsfeed/ent"
	fileUploader "newsfeed/modules/attachment/file_uploader"
	"newsfeed/modules/attachment/models"
	"newsfeed/modules/attachment/pb"

	"github.com/gin-gonic/gin"
)

type AttachmentUploaderServiceInterface interface {
	UploadAttachments(attachments models.Attachments) (*ent.Attachment, error)
	GetSingleAttachment(context *gin.Context, attachmentPath string)
	Delete(attachmentPath string) error
}

type AttachmentUploaderService struct {
	uploader   fileUploader.FileUploaderInterface
	gRPCClient pb.AttachmentServiceClient
}

func NewAttachmentService(gRPCClient pb.AttachmentServiceClient, uploader fileUploader.FileUploaderInterface) AttachmentUploaderServiceInterface {
	service := &AttachmentUploaderService{
		uploader:   uploader,
		gRPCClient: gRPCClient,
	}
	return service
}

func (aus AttachmentUploaderService) Delete(attachmentPath string) error {
	return nil
}

func (aus AttachmentUploaderService) GetSingleAttachment(context *gin.Context, attachmentPath string) {
	
}

func (aus AttachmentUploaderService) UploadAttachments(attachments models.Attachments) (*ent.Attachment, error) {
	requestAttachments := &pb.RequestAttachments{}
	for _, attachment := range attachments.Attachment {
		tmpAttachment := pb.RequestAttachment{Name: attachment.Name, Path: attachment.Path}
		requestAttachments.Attachments = append(requestAttachments.Attachments, &tmpAttachment)
	}

	_, err := aus.gRPCClient.CreateMultiple(context.Background(), requestAttachments)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return nil, nil
}
