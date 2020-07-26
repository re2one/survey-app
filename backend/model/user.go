package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	// Id       uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	//Role     string `json:"role"`
	Password string `json:"password"`
}

func (User) TableName() string { return "users" }
