package model

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Survey   Survey
	SurveyId uint   `gorm:"foreignkey:SurveyRefer" json:"surveyid"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Type     string `json:"type"`
	First    string `json:"first"`
}

func (Question) TableName() string { return "questions" }
