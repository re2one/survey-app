package repository

import "backend/model"

type MultipleChoiceRepository interface {
	Post(*model.ChoiceAnswer) (*model.ChoiceAnswer, error)
}
