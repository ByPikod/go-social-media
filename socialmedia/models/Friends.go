package models

type Friends struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Inviting  uint `json:"author_id" gorm:"index:idx_friend,unique"`
	Accepting uint `json:"friend_id" gorm:"index:idx_friend,unique"`
}
