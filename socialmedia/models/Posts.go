package models

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	Content      string `json:"content"`
	AuthorID     uint   `json:"author_id"`
	Author       User   `json:"author" gorm:"foreignKey:AuthorID"`
	AttachmentID uint   `json:"attachment_id"`
}
