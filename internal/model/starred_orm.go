package model 

import "time"

type Starred struct {
	StarredId string    `json:"starred_id"`
	ItemId    string    `json:"item_id" gorm:"foreignKey:ItemId"`
	UserId    string    `json:"user_id" gorm:"foreignKey:UserId"`
	MarkedAt  time.Time `json:"marked_at"`
}