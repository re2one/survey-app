package model

import (
	"github.com/jinzhu/gorm"
)

type Puzzlepiece struct {
	gorm.Model
	Question   Question
	QuestionId uint   `gorm:"foreignkey:QuestionRefer" json:"questionid"`
	Position   uint   `json:"position"`
	Image      string `json:"image"`
}

func (Puzzlepiece) TableName() string { return "puzzlepieces" }
