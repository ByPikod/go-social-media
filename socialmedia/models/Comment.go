package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	AuthorID     uint   `json:"author_id"`
	Author       User   `json:"author" gorm:"foreignKey:AuthorID"`
	PostID       uint   `json:"post_id"`
	PostType     string `json:"post_type" gorm:"index"` // reply or post
	Content      string `json:"content"`
	AttachmentID uint   `json:"attachment_id"`
}
