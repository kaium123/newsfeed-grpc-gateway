package service

import (
	"context"
	"newsfeed/common/logger"
	"newsfeed/common/utils"
	"newsfeed/modules/reaction/models"
	"newsfeed/modules/reaction/repository"
)

type ReactionServiceInterface interface {
	GetReaction(postID int, postType string, react_type string, ctx context.Context) ([]*models.RespReaction, error)
	CreateReact(Reaction models.React, ctx context.Context) (uint, error)
	AllReaction(postID int, postType string, ctx context.Context) ([]*models.RespReaction, error)
	UpdateReaction(id int, Reaction models.React, ctx context.Context) (uint, error)
}

type ReactionService struct {
	repository repository.ReactionRepositoryInterface
}

func NewReactionService(repository repository.ReactionRepositoryInterface) ReactionServiceInterface {
	return &ReactionService{repository: repository}
}

func (u ReactionService) GetReaction(postID int, postType string, react_type string, ctx context.Context) ([]*models.RespReaction, error) {

	resp, err := u.repository.FindByReactID(postID, postType, react_type, ctx)
	if err != nil {
		return nil, err
	}

	reacts := []*models.RespReaction{}

	for _, react := range resp {
		if react.ReactType!=react_type {
			continue
		}
		tmpReact := models.RespReaction{}
		utils.CopyStructToStruct(react, &tmpReact)
		reacts = append(reacts, &tmpReact)
	}

	return reacts, nil
}

func (u ReactionService) CreateReact(reaction models.React, ctx context.Context) (uint, error) {
	vErr := reaction.Validate()
	if vErr != nil {
		return 0, vErr
	}
	logger.LogError(reaction)
	id, err := u.repository.CreateReact(reaction, ctx)
	return id, err

}

func (u ReactionService) AllReaction(postID int, postType string, ctx context.Context) ([]*models.RespReaction, error) {

	resp, err := u.repository.AllReaction(postID, postType, ctx)
	reacts := []*models.RespReaction{}

	for _, react := range resp {
		tmpReact := models.RespReaction{}
		utils.CopyStructToStruct(react, &tmpReact)
		reacts = append(reacts, &tmpReact)
	}
	return reacts, err
	//}

}

func (u ReactionService) UpdateReaction(id int, reaction models.React, ctx context.Context) (uint, error) {
	vErr := reaction.Validate()
	if vErr != nil {
		return 0, nil
	}
	_, err := u.repository.UpdateReaction(id, reaction, ctx)
	//if err != nil {
	return uint(id), err
	//}

}
