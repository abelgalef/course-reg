package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"type:varchar(255)" binding:"required,min=3"`
	DeptCode    string         `json:"dept_code" binding:"required" gorm:"unique"`
	Active      bool           `json:"-" gorm:"defalut:true"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Courses     []Course       `json:"courses" gorm:"foreignKey:DeptID"`
	Constraints []Constraint   `json:"-" gorm:"foreignKey:DeptID"`
}

type Constraint struct {
	ID                 int       `json:"id" gorm:"primaryKey"`
	AllowedNumOfCourse int       `json:"allowed_num_of_courses" binding:"required"`
	AllowedNumOfCredit int       `json:"allowed_num_of_credits" binding:"required"`
	SemsterExpiry      time.Time `json:"semster_expiry" binding:"required"`
	DeptID             int       `json:"dept" binding:"required"`
	CreatedBy          int       `json:"created_by"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
}
