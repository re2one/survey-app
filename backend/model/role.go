package model

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	// Id       uint   `gorm:"primary_key" json:"id"`
	User   User
	UserId uint   `gorm:"foreignkey:UserRefer" json:"userid"`
	Role   string `json:"role"`
}

func (Role) TableName() string { return "roles" }
