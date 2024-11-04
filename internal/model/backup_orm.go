package model

import "time"

type Backup struct {
	BackupId  string    `json:"backup_id"`
	UserId    string    `json:"user_id" gorm:"foreignKey:UserId"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}
