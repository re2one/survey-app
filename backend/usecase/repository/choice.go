package repository

import "backend/model"

type ChoiceRepository interface {
	GetAll(string) ([]*model.Choice, error)
	Get(string) (*model.Choice, error)
	GetByText(uint, string) (*model.Choice, error)
	Post(string, *model.Choice) (*model.Choice, error)
	Put(*model.Choice) (*model.Choice, error)
	Delete(*model.Choice) (*model.Choice, error)
}
