package models

import (
	"errors"
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	Path string `json:"path"`
	Name string `json:"name"`
}

type React struct {
	gorm.Model
	PostID        uint   `json:"post_id"`
	ReactType     string `json:"react_type"`
	ReactedUserID uint   `json:"reacted_user_id"`
	PostType string `json:"post_type"`
}

type RespReaction struct {
	ID            uint      `json:"id"`
	PostID        uint      `json:"post_id"`
	ReactType     string    `json:"react_type"`
	ReactedUserID uint      `json:"reacted_user_id"`
	CreateAt      time.Time `json:"create_post"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (u *React) Validate() error {
	return v.ValidateStruct(u,
		v.Field(&u.ReactType,
			v.By(func(value interface{}) error {
				if u.ReactType != "like" && u.ReactType != "love" && u.ReactType != "care" {
					return errors.New("invalid react type")
				}
				return nil
			}),
		),
	)
}
