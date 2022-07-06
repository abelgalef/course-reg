package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(255)" binding:"required,min=3"`
	Active    bool           `json:"-" gorm:"defalut:true"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Teachers  []User         `json:"-" gorm:"foreignKey:DeptID"`
	Courses   []Course       `json:"courses" gorm:"foreignKey:DeptID"`
}
