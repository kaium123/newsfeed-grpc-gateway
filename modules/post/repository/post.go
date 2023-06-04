package repository

import (
	"context"
	"net/http"
	"newsfeed/common/logger"
	"newsfeed/ent"
	"newsfeed/ent/post"
	"newsfeed/errors"
	"newsfeed/modules/post/models"
)

type PostRepositoryInterface interface {
	FindByID(id int, ctx context.Context) (*ent.Post, error)
	CreatePost(post models.Post, ctx context.Context) (uint, error)
	AllPost(ctx context.Context) ([]*ent.Post, []map[string]interface{}, error)
	UpdatePost(id int, post models.Post, ctx context.Context) (uint, error)
}

type PostRepository struct {
	Db     *ent.Client
	logger logger.LoggerInterface
}

func NewPostRepository(db *ent.Client, logger logger.LoggerInterface) PostRepositoryInterface {
	return &PostRepository{Db: db, logger: logger}
}

func (r *PostRepository) FindByID(id int, ctx context.Context) (*ent.Post, error) {
	//post := models.Post{}
	resp, err := r.Db.Post.Query().
		Where(post.ID(id)).
		WithAttachments().
		WithComments().
		WithReacts().
		First(ctx)
	if err != nil {
		return nil, &errors.ApplicationError{TranslationKey: "failedFindingPost", HttpCode: http.StatusInternalServerError}
	}

	return resp, nil
}

func (r *PostRepository) CreatePost(post models.Post, ctx context.Context) (uint, error) {
	resp, err := r.Db.Post.Create().
		SetContent(post.Content).
		AddAttachmentIDs(post.Attachments...).
		AddAuthorIDs(int(post.AuthorID)).Save(ctx)
	if err != nil {
		return 0, errors.ApplicationError{TranslationKey: "failedCreatePost", HttpCode: http.StatusInternalServerError}
	}

	return uint(resp.ID), nil
}

func (r *PostRepository) AllPost(ctx context.Context) ([]*ent.Post, []map[string]interface{}, error) {

	resp, err := r.Db.Post.Query().
		WithAttachments().
		WithComments().
		WithReacts().
		All(ctx)
	if err != nil {
		return nil, nil, &errors.ApplicationError{TranslationKey: "failedFindingPost", HttpCode: http.StatusInternalServerError}
	}

	return resp, nil, nil

}

func (r *PostRepository) UpdatePost(id int, post models.Post, ctx context.Context) (uint, error) {
	_, err := r.Db.Post.UpdateOneID(id).
		SetContent(post.Content).
		AddAttachmentIDs(post.Attachments...).
		AddAuthorIDs(int(post.AuthorID)).Save(ctx)
	if err != nil {
		return 0, errors.ApplicationError{TranslationKey: "failedCreatePost", HttpCode: http.StatusInternalServerError}
	}

	return post.ID, nil
}
