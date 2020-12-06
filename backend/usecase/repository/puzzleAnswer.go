package repository

import "backend/model"

type PuzzleAnswerRepository interface {
	Post([]*model.PuzzleAnswer) ([]*model.PuzzleAnswer, error)
}
