package models

import (
	"time"
)

type Right struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Tag       string    `json:"tag" gorm:"type:varchar(50);not null;index" binding:"required,min=3"`
	Desc      string    `json:"description" gorm:"type:varchar(255);not null"`
	Active    bool      `json:"-" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
