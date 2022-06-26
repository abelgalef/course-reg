package models

import "time"

type Course struct {
	ID        int       `gorm:"primaryKey"`
	CourseID  string    `json:"id" gorm:"unique;not null;index;type:varchar(25)"`
	Name      string    `json:"name" gorm:"unique; not null; index;type:varchar(255)"`
	Active    bool      `json:"-" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Files     []Files   `json:"Resources" gorm:"foreignKey:ID"`
	Users     []User    `json:"-" gorm:"many2many:course_user"`
}

type CourseUser struct {
	CourseID int `gorm:"primaryKey"`
	UserID   int `gorm:"primaryKey"`
}
