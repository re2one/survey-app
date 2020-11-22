package repository

import "backend/model"

type PuzzleRepository interface {
	Put(string, string, []*model.Puzzlepiece) ([]*model.Puzzlepiece, error)
	GetAll(string) ([]*model.Puzzlepiece, error)
}
