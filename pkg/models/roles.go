package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Tag       string         `json:"tag" gorm:"type:varchar(50);not null;index" binding:"required,min=3"`
	Desc      string         `json:"description" gorm:"type:varchar(255);not null"`
	Active    bool           `json:"-" gorm:"default:true"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Users     []User         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rights    []Right        `json:"premissions" gorm:"many2many:role_permission;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type RoleRight struct {
	RoleID    uint      `gorm:"primaryKey"`
	RightID   uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

var NeededRole = Role{
	Tag:  "ADMIN",
	Desc: "This role grants the user to create, delete and edit any number of employees, products, catagories, permissions and roles.",
	Users: []User{
		{FirstName: "root", LastName: "root", Email: "root@localhost", Password: "XohImNooBHFR0OVvjcYpJ3NgPQ1qq73WKhHvch0VQtg=", RoleID: 1},
	},
	Rights: []Right{
		{Tag: "CREATE_EMPLOYEE", Desc: "This permission grants the role to create any number of employees."},
		{Tag: "REVOKE_PERMISSION", Desc: "This permission grants the role to revoke any number of given permissions."},
		{Tag: "GRANT_PERMISSION", Desc: "This permission grants the role to grant any number of permissions to any number of employees."},
		{Tag: "DELETE_PERMISSION", Desc: "This permission grants the role to delete any number of permissions."},
	},
}
