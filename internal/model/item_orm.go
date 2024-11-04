package model

import "time"

type Item struct {
	ItemId    string    `json:"item_id"`
	OwnerId   string    `json:"owner_id" gorm:"foreignKey:OwnerId"`
	ItemName  string    `json:"item_name"`
	ItemPath  string    `json:"item_path"`
	Type      string    `json:"type"`
	Size      string    `json:"size"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}