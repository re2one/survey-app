package model

import (
	"github.com/jinzhu/gorm"
)

type PuzzleAnswer struct {
	gorm.Model
	Question   Question
	QuestionId uint   `gorm:"foreignkey:QuestionRefer" json:"questionid"`
	Email      string `json:"email"`
	Position   string `json:"position"`
	Image      string `json:"image"`
	Tapped     bool   `json:"tapped"`
}

func (PuzzleAnswer) TableName() string { return "puzzleanswers" }
