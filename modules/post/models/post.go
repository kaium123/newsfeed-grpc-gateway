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

type Post struct {
	gorm.Model
	AuthorID       int    `json:"author_id"`
	Content        string `json:"content"`
	DateOfCreation int    `json:"date_of_creation"`
	Attachments    []int  `json:"attachment_ids" gorm:"type:integer[]"`
	TagFriends     []int  `json:"tag_friends"`
}

type RespPost struct {
	ID             int          `json:"id"`
	AuthorID       int          `json:"author_id"`
	Content        string       `json:"content"`
	DateOfCreation int          `json:"date_of_creation"`
	Attachments    []Attachment `json:"attachment_ids"`
	Comments       []Comment    `json:"comments"`
	Reacts         []React      `json:"react"`
	DeletedAt      time.Time    `json:"deleted_at"`
	CreateAt       time.Time    `json:"create_post"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type Comment struct {
	ID              uint      `json:"id"`
	PostID          uint      `json:"post_id"`
	Content         string    `json:"content"`
	Attachments     []int     `json:"attachment_ids"`
	ParentCommentID uint      `json:"parent_comment_id"`
	CreateAt        time.Time `json:"create_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type React struct {
	ID            uint      `json:"id"`
	ReactType     string    `json:"react_type"`
	ReactedUserID uint      `json:"reacted_user_id"`
	CreateAt      time.Time `json:"create_post"`
	UpdatedAt     time.Time `json:"updated_at"`
}
