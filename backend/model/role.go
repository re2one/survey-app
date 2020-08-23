package model

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	User   User
	UserId uint   `gorm:"foreignkey:UserRefer" json:"userid"`
	Role   string `json:"role"`
}

func (Role) TableName() string { return "roles" }
