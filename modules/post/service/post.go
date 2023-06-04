package service

import (
	"context"
	"newsfeed/common/utils"
	"newsfeed/modules/post/models"
	"newsfeed/modules/post/repository"
)

type PostServiceInterface interface {
	GetPost(id int, ctx context.Context) (*models.RespPost, error)
	CreatePost(post models.Post, ctx context.Context) (uint, error)
	AllPost(ctx context.Context) ([]models.RespPost, []map[string]interface{}, error)
	UpdatePost(id int, post models.Post, ctx context.Context) (uint, error)
}

type PostService struct {
	repository repository.PostRepositoryInterface
}

func NewPostService(repository repository.PostRepositoryInterface) PostServiceInterface {
	return &PostService{repository: repository}
}

func (u PostService) GetPost(id int, ctx context.Context) (*models.RespPost, error) {

	resp, err := u.repository.FindByID(id, ctx)
	if err != nil {
		return nil, err
	}
	post := models.RespPost{}
	utils.CopyStructToStruct(resp, &post)
	utils.CopyStructToStruct(resp.Edges.Attachments, &post.Attachments)
	utils.CopyStructToStruct(resp.Edges.Comments, &post.Comments)
	utils.CopyStructToStruct(resp.Edges.Reacts, &post.Reacts)

	return &post, nil
}

func (u PostService) CreatePost(post models.Post, ctx context.Context) (uint, error) {

	id, err := u.repository.CreatePost(post, ctx)
	//if err != nil {
	return id, err
	//}

}

func (u PostService) AllPost(ctx context.Context) ([]models.RespPost, []map[string]interface{}, error) {

	resp, res, err := u.repository.AllPost(ctx)
	posts := []models.RespPost{}
	for _, post := range resp {
		tmpPost := &models.RespPost{}
		utils.CopyStructToStruct(post, &tmpPost)
		utils.CopyStructToStruct(post.Edges.Attachments, &tmpPost.Attachments)
		utils.CopyStructToStruct(post.Edges.Comments, &tmpPost.Comments)
		utils.CopyStructToStruct(post.Edges.Reacts, &tmpPost.Reacts)
		posts = append(posts, *tmpPost)
	}
	return posts, res, err
	//}

}

func (u PostService) UpdatePost(id int, post models.Post, ctx context.Context) (uint, error) {

	_, err := u.repository.UpdatePost(id, post, ctx)
	//if err != nil {
	return uint(id), err
	//}

}
