package repository

import "backend/model"

type PuzzleAnswerRepository interface {
	Post([]*model.PuzzleAnswer) ([]*model.PuzzleAnswer, error)
	GetUserSolution(string, string) ([]*model.PuzzleAnswer, error)
}
