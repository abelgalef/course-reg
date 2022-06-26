package models

import "time"

type Department struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255)" binding:"required, min=3"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Resources []User    `json:"-" gorm:"foreignKey:ID"`
	Courses   []Course  `json:"courses" gorm:"foreignKey:ID"`
}
