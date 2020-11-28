package model

import "github.com/jinzhu/gorm"

type Bracket struct {
	gorm.Model
	Survey   Survey
	SurveyId uint   `gorm:"foreignkey:SurveyRefer" json:"surveyid"`
	Name     string `json:"name"`
}

func (Bracket) TableName() string { return "brackets" }
