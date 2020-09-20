package model

import (
	"github.com/jinzhu/gorm"
)

type Answered struct {
	gorm.Model
	User       User
	UserId     uint `gorm:"foreignkey:UserRefer" json:"userid"`
	Question   Question
	QuestionId uint `gorm:"foreignkey:QuestionRefer" json:"questionid"`
	Answered   bool `json:"answered"`
}

func (Answered) TableName() string { return "answered" }
