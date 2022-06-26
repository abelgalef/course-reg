package models

import "time"

type Files struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`
	Type      string    `gorm:"type:varchar(255)"`
	Size      string    `gorm:"type:varchar(255)"`
	Loc       string    `json:"url" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
