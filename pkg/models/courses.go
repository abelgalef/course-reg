package models

import (
	"time"
)

type Course struct {
	ID        int       `gorm:"primaryKey"`
	CourseID  string    `json:"course_id" gorm:"unique;not null;index;type:varchar(25)"`
	Credits   int       `json:"credits" gorm:"default:1"`
	Name      string    `json:"name" gorm:"unique; not null; index;type:varchar(255)"`
	Active    bool      `json:"-" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	DeptID    int       `json:"dept_id" binding:"required"`
	Files     []Files   `json:"resources" gorm:"foreignKey:ID"`
}

type CourseUser struct {
	CourseID  int       `gorm:"primaryKey"`
	UserID    int       `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;primaryKey"`
}
