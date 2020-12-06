package repository

import (
	"github.com/jinzhu/gorm"

	"backend/model"
	"backend/usecase/repository"
)

type puzzleAnswerRepository struct {
	db *gorm.DB
}

func NewPuzzleAnswerRepository(db *gorm.DB) repository.PuzzleAnswerRepository {
	db.AutoMigrate(&model.PuzzleAnswer{})
	return &puzzleAnswerRepository{db}
}

func (par *puzzleAnswerRepository) Post(pieces []*model.PuzzleAnswer) ([]*model.PuzzleAnswer, error) {
	for _, piece := range pieces {
		par.db.Create(&piece)
	}

	return nil, nil
}
