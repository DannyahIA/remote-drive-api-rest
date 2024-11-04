package model

import "time"

type Recent struct {
	RecentId  string    `json:"recent_id"`
	ItemId    string    `json:"item_id" gorm:"foreignKey:ItemId"`
	UserId    string    `json:"user_id" gorm:"foreignKey:UserId"`
	AcessedAt time.Time `json:"acessed_at"`
}