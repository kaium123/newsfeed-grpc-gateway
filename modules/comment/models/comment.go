package models

import (
	"time"

	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	Path string `json:"path"`
	Name string `json:"name"`
}

type Comment struct {
	gorm.Model
	PostID             uint   `json:"post_id"`
	Content            string `json:"content"`
	CommentAttachments []int  `json:"attachments" gorm:"type:integer[]"`
	ParentCommentID    uint   `json:"parent_comment_id"`
	Rank               uint   `json:"rank"`
}

type RespComment struct {
	ID              uint         `json:"id"`
	PostID          uint         `json:"post_id"`
	Content         string       `json:"content"`
	Attachments     []Attachment `json:"attachment_ids"`
	ParentCommentID uint         `json:"parent_comment_id"`
	React           []React      `json:"react"`
	CreateAt        time.Time    `json:"create_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

type React struct {
	ID            uint      `json:"id"`
	ReactType     string    `json:"react_type"`
	ReactedUserID uint      `json:"reacted_user_id"`
	CreateAt      time.Time `json:"create_post"`
	UpdatedAt     time.Time `json:"updated_at"`
}
