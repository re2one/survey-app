package model

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Survey             Survey
	SurveyId           uint   `gorm:"foreignkey:SurveyRefer" json:"surveyid"`
	Title              string `json:"title"`
	Text               string `json:"text"`
	Type               string `json:"type"`
	First              string `json:"first"`
	Bracket            string `json:"bracket"`
	Next               string `json:"next"`
	SecondToNext       string `json:"secondToNext"`
	TypeOfNextQuestion string `json:"typeOfNextQuestion"`
}

func (Question) TableName() string { return "questions" }
