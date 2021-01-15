package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name               string `json:"name"`
	Email              string `json:"email" gorm:"unique"`
	Salt               string `json:"salt"`
	Password           string `json:"password"`
	Thesis             string `json:"thesis"`
	TermsAndConditions string `json:"termsAndConditions"`
}

func (User) TableName() string { return "users" }
