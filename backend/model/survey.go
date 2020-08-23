package model

import (
	"github.com/jinzhu/gorm"
)

type Survey struct {
	gorm.Model
	Title        string `json:"title" gorm:"unique`
	Summary      string `json:"summary"`
	Disclaimer   string `json:"disclaimer"`
	Introduction string `json:"introduction"`
}

func (Survey) TableName() string { return "surveys" }
