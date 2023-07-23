package service

import (
	"context"
	"newsfeed/common/logger"
	"newsfeed/common/utils"
	"newsfeed/modules/comment/models"
	"newsfeed/modules/comment/repository"
)

type CommentServiceInterface interface {
	GetComment(id int, ctx context.Context) (*models.RespComment, error)
	Comment(Comment models.Comment, ctx context.Context) (uint, error)
	AllComment(postID int, ctx context.Context) ([]models.RespComment, error)
	UpdateComment(id int, Comment models.Comment, ctx context.Context) (uint, error)
}

type CommentService struct {
	repository repository.CommentRepositoryInterface
}

func NewCommentService(repository repository.CommentRepositoryInterface) CommentServiceInterface {
	return &CommentService{repository: repository}
}

func (u CommentService) GetComment(id int, ctx context.Context) (*models.RespComment, error) {

	resp, err := u.repository.FindByID(id, ctx)
	if err != nil {
		return nil, err
	}
	logger.LogError(resp.Edges.Attachments)
	comment := models.RespComment{}
	utils.CopyStructToStruct(resp, &comment)
	utils.CopyStructToStruct(resp.Edges.Attachments, &comment.Attachments)
	utils.CopyStructToStruct(resp.Edges.Reacts, &comment.React)

	return &comment, nil
}

func (u CommentService) Comment(Comment models.Comment, ctx context.Context) (uint, error) {

	id, err := u.repository.Comment(Comment, ctx)
	//if err != nil {
	return id, err
	//}

}

func (u CommentService) AllComment(postID int, ctx context.Context) ([]models.RespComment, error) {

	resp, err := u.repository.AllComment(postID, ctx)
	comments := []models.RespComment{}
	for _, comment := range resp {
		tmpComment := &models.RespComment{}
		utils.CopyStructToStruct(comment, &tmpComment)
		utils.CopyStructToStruct(comment.Edges.Attachments, &tmpComment.Attachments)
		utils.CopyStructToStruct(comment.Edges.Reacts, &tmpComment.React)
		comments = append(comments, *tmpComment)
	}

	return comments, err
	//}

}

func (u CommentService) UpdateComment(id int, Comment models.Comment, ctx context.Context) (uint, error) {

	_, err := u.repository.UpdateComment(id, Comment, ctx)
	//if err != nil {
	return uint(id), err
	//}

}
