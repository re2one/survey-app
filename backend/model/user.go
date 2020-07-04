package model

import (
	
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	//ID			uint	`gorm:"primary_key" json:"id"`
	Name		string	`json:"name"`
	Email		string	`gorm:"unique"`
	Role		string	`json:"role"`
	Password	string 	`json:"password"`
}

func (User) TableName() string { return "users" }