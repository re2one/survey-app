package repository

import "backend/model"

type AnsweredRepository interface {
	Get(*model.User, *model.Question) (map[uint]*model.Answered, error)
	Post(*model.User, *model.Question) (*model.Answered, error)
}
