package model

import (
	"github.com/jinzhu/gorm"
)

type Answered struct {
	gorm.Model
	User       User
	UserId     uint `gorm:"foreignkey:UserRefer" json:"userid"`
	QuestionId uint `json:"questionid"`
	Answered   bool `json:"answered"`
}

func (Answered) TableName() string { return "answered" }
