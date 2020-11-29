package model

import (
	"github.com/jinzhu/gorm"
)

type Choice struct {
	gorm.Model
	Question           Question
	QuestionId         uint   `gorm:"foreignkey:QuestionRefer" json:"questionid"`
	Text               string `json:"text"`
	NextQuestion       string `json:"nextQuestion"`
	SecondToNext       string `json:"secondToNext"`
	TypeOfNextQuestion string `json:"typeOfNextQuestion"`
}

func (Choice) TableName() string { return "choices" }
