package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name" gorm:"type:varchar(30)" binding:"required,min=3"`
	LastName  string    `json:"last_name" gorm:"type:varchar(30)" binding:"required,min=3"`
	Email     string    `json:"email" binding:"required,email" gorm:"unique;not null;index;type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)" binding:"required,min=6"`
	Active    bool      `json:"-" gorm:"default:true"`
	RoleID    int       `json:"role_id" gorm:"default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type LoginResponse struct {
	User
	Token string `json:"token"`
}

type LoginData struct {
	Email string `binding:"required,email"`
	Pass  string
}

type RegResponse struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}
