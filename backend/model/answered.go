package model

import (
	"github.com/jinzhu/gorm"
)

type Answered struct {
	gorm.Model
	User       User
	UserId     uint `gorm:"foreignkey:UserRefer" json:"userid"`
	QuestionId uint `json:"questionid"`
	SurveyId   uint `json:"surveyid"`
	Answered   bool `json:"answered"`
	Viewed     bool `json:"viewed"`
	Order      uint `json:"order"`
}

func (Answered) TableName() string { return "answered" }
