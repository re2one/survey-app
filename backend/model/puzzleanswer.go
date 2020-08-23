package model

import (
	"github.com/jinzhu/gorm"
)

type PuzzleAnswer struct {
	gorm.Model
	Question   Question
	QuestionId uint   `gorm:"foreignkey:QuestionRefer" json:"questionid"`
	UserId     uint   `json:"userid"`
	Position   uint   `json:"position"`
	Image      string `json:"image"`
}

func (PuzzleAnswer) TableName() string { return "puzzleanswers" }
