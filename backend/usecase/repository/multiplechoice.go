package repository

import "backend/model"

type MultipleChoiceRepository interface {
	Post(*model.ChoiceAnswer) (*model.ChoiceAnswer, error)
	Get(uint, string) (*model.ChoiceAnswer, error)
	GetAll(uint) ([]*model.ChoiceAnswer, error)
}
