package repository

import "backend/model"

type BracketRepository interface {
	Get(string) ([]*model.Bracket, error)
	Post(string, *model.Bracket) (*model.Bracket, error)
}
