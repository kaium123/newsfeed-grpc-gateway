package repository

import (
	"context"
	"net/http"
	"newsfeed/common/logger"
	"newsfeed/ent"
	"newsfeed/ent/comment"
	"newsfeed/ent/post"
	"newsfeed/ent/react"
	"newsfeed/errors"
	"newsfeed/modules/reaction/models"
)

type ReactionRepositoryInterface interface {
	FindByReactID(postID int, postType string, react_type string, ctx context.Context) ([]*ent.React, error)
	CreateReact(Reaction models.React, ctx context.Context) (uint, error)
	AllReaction(postID int, postType string, ctx context.Context) ([]*ent.React, error)
	UpdateReaction(id int, react models.React, ctx context.Context) (uint, error)
}

type ReactionRepository struct {
	Db     *ent.Client
	logger logger.LoggerInterface
}

func NewReactionRepository(db *ent.Client, logger logger.LoggerInterface) ReactionRepositoryInterface {
	return &ReactionRepository{Db: db, logger: logger}
}

func (r *ReactionRepository) FindByReactID(postID int, postType string, react_type string, ctx context.Context) ([]*ent.React, error) {

	if postType == "post" {
		baseQuery := r.Db.Post.Query().Where(post.ID(postID)).WithReacts()
		resp, err := baseQuery.First(ctx)
		if err != nil {
			return nil, &errors.ApplicationError{TranslationKey: "failedFindingReaction", HttpCode: http.StatusInternalServerError}
		}

		return resp.Edges.Reacts, nil
	} else {
		baseQuery := r.Db.Comment.Query().Where(comment.ID(postID)).WithReacts()
		resp, err := baseQuery.First(ctx)
		if err != nil {
			return nil, &errors.ApplicationError{TranslationKey: "failedFindingReaction", HttpCode: http.StatusInternalServerError}
		}

		return resp.Edges.Reacts, nil
	}

}

func (r *ReactionRepository) CreateReact(react models.React, ctx context.Context) (uint, error) {
	baseQuery := r.Db.React.Create().SetReactType(react.ReactType).AddReactedUserIDs(int(react.ReactedUserID))
	if react.PostType == "post" {
		baseQuery = baseQuery.AddPostIDs(int(react.PostID)).SetPostType(react.PostType)
	} else {
		baseQuery = baseQuery.AddCommentIDs(int(react.PostID)).SetPostType(react.PostType)
	}
	resp, err := baseQuery.Save(ctx)
	logger.LogError(err)
	if err != nil {
		return 0, &errors.ApplicationError{TranslationKey: "failedCreatePost", HttpCode: http.StatusInternalServerError}
	}

	logger.LogError(resp)

	return uint(resp.ID), nil
}

func (r *ReactionRepository) AllReaction(postID int, postType string, ctx context.Context) ([]*ent.React, error) {

	if postType == "post" {
		baseQuery := r.Db.Post.Query().Where(post.ID(postID)).WithReacts()
		resp, err := baseQuery.First(ctx)
		if err != nil {
			return nil, &errors.ApplicationError{TranslationKey: "failedFindingReaction", HttpCode: http.StatusInternalServerError}
		}

		return resp.Edges.Reacts, nil
	} else {
		baseQuery := r.Db.Comment.Query().Where(comment.ID(postID)).WithReacts()
		resp, err := baseQuery.First(ctx)
		if err != nil {
			return nil, &errors.ApplicationError{TranslationKey: "failedFindingReaction", HttpCode: http.StatusInternalServerError}
		}

		return resp.Edges.Reacts, nil
	}
}

func (r *ReactionRepository) UpdateReaction(id int, reacts models.React, ctx context.Context) (uint, error) {
	reacts.ID = uint(id)
	baseQuery := r.Db.React.Update().Where(react.And(react.ID(id), react.PostType(reacts.PostType))).
		SetReactType(reacts.ReactType).SetPostType(reacts.PostType)

	_, err := baseQuery.Save(ctx)
	logger.LogError(err)
	if err != nil {
		return 0, &errors.ApplicationError{TranslationKey: "failedCreatePost", HttpCode: http.StatusInternalServerError}
	}

	return uint(reacts.ID), nil
}
