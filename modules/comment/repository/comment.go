package repository

import (
	"context"
	"net/http"
	"newsfeed/common/logger"
	"newsfeed/ent"
	"newsfeed/ent/comment"
	"newsfeed/ent/post"
	"newsfeed/errors"
	"newsfeed/modules/comment/models"
)

type CommentRepositoryInterface interface {
	FindByID(id int, ctx context.Context) (*ent.Comment, error)
	Comment(Comment models.Comment, ctx context.Context) (uint, error)
	AllComment(postID int, ctx context.Context) ([]*ent.Comment, error)
	UpdateComment(id int, Comment models.Comment, ctx context.Context) (uint, error)
}

type CommentRepository struct {
	Db     *ent.Client
	logger logger.LoggerInterface
}

func NewCommentRepository(db *ent.Client, logger logger.LoggerInterface) CommentRepositoryInterface {
	return &CommentRepository{Db: db, logger: logger}
}

func (r *CommentRepository) FindByID(id int, ctx context.Context) (*ent.Comment, error) {

	resp, err := r.Db.Comment.Query().Where(comment.ID(id)).
		WithAttachments().
		WithReacts().
		First(ctx)
	if err != nil {
		return nil, &errors.ApplicationError{TranslationKey: "failedFindingComment", HttpCode: http.StatusInternalServerError}
	}

	return resp, nil
}

func (r *CommentRepository) Comment(comment models.Comment, ctx context.Context) (uint, error) {
	logger.LogError(comment.CommentAttachments)
	resp, err := r.Db.Comment.Create().SetContent(comment.Content).
		AddAttachmentIDs(comment.CommentAttachments...).AddPostIDs(int(comment.PostID)).Save(ctx)
	if err != nil {
		return 0, &errors.ApplicationError{TranslationKey: "failedComment", HttpCode: http.StatusInternalServerError}
	}

	return uint(resp.ID), nil
}

func (r *CommentRepository) AllComment(postID int, ctx context.Context) ([]*ent.Comment, error) {

	resp, err := r.Db.Post.Query().Where(post.ID(postID)).WithComments().
		First(ctx)
	if err != nil {
		logger.LogError(err)
		return nil, &errors.ApplicationError{TranslationKey: "failedFindingComment", HttpCode: http.StatusInternalServerError}
	}

	logger.LogError(resp.Edges.Comments)

	comments := []*ent.Comment{}

	for _, comment := range resp.Edges.Comments {
		comments = append(comments, comment)
	}
	return comments, nil

}

func (r *CommentRepository) UpdateComment(id int, comment models.Comment, ctx context.Context) (uint, error) {
	_, err := r.Db.Comment.UpdateOneID(id).SetContent(comment.Content).
		AddAttachmentIDs(comment.CommentAttachments...).AddPostIDs(int(comment.PostID)).Save(ctx)
	if err != nil {
		logger.LogError(err)
		return 0, &errors.ApplicationError{TranslationKey: "failedComment", HttpCode: http.StatusInternalServerError}
	}
	return uint(id), nil
}
