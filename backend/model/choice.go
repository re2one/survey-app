package model

import (
	"github.com/jinzhu/gorm"
)

type Choice struct {
	gorm.Model
	Question     Question
	QuestionId   uint   `gorm:"foreignkey:QuestionRefer" json:"questionid"`
	Text         string `json:"text"`
	NextQuestion string `json:"nextQuestion"`
}

func (Choice) TableName() string { return "choices" }
