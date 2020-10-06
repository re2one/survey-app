package model

import (
	"github.com/jinzhu/gorm"
)

type ChoiceAnswer struct {
	gorm.Model
	Question   Question
	QuestionId uint   `gorm:"foreignkey:QuestionRefer" json:"questionid"`
	Email      string `json:"email"`
	Text       string `json:"text"`
}

func (ChoiceAnswer) TableName() string { return "choiceanswers" }
