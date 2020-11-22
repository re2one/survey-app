package model

import (
	"github.com/jinzhu/gorm"
)

type Puzzlepiece struct {
	gorm.Model
	Question   Question
	QuestionId uint   `gorm:"foreignkey:QuestionRefer" json:"questionid"`
	Position   string `json:"position"`
	Image      string `json:"image"`
	Tapped     bool   `json:"tapped"`
	Empty      bool   `json:"empty"`
}

func (Puzzlepiece) TableName() string { return "puzzlepieces" }
