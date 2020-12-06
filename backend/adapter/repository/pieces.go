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
	for _, p := range pieces {

		/*var p2 model.PuzzleAnswer
		err := par.db.Where("questionId = ? and email = ?", p.Question.ID, p.Email).First(&p2).Error

		if err == nil {
			log.Error().Msg("puzzle answer already exists")
			continue
		}*/

		par.db.Create(&p)
	}

	return pieces, nil
}
