package model

import "time"

type Trash struct {
	TrashId   string    `json:"trash_id"`
	ItemId    string    `json:"item_id" gorm:"foreignKey:ItemId"`
	UserId    string    `json:"user_id" gorm:"foreignKey:UserId"`
	Path      string    `json:"path"`
	OldPath   string    `json:"old_path"`
	CreatedAt time.Time `json:"created_at"`
}