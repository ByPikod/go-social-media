package models

type Like struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	AuthorID uint   `json:"author_id" gorm:"index:idx_like,unique"`
	PostID   uint   `json:"post_id" gorm:"index:idx_like,unique"`
	PostType string `json:"post_type" gorm:"index:idx_like,unique"` // comment or post
}
