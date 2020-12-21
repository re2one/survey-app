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

func (par *puzzleAnswerRepository) GetUserSolution(email, questionId string) ([]*model.PuzzleAnswer, error) {
	var retrievedPieces []*model.PuzzleAnswer
	if err := par.db.Where("question_id = ? and email = ?", questionId, email).Find(&retrievedPieces).Error; err != nil {
		return nil, err
	}
	return retrievedPieces, nil
}
