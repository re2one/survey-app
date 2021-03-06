package repository

import "backend/model"

type AnsweredRepository interface {
	Get(*model.User, string) (map[uint]model.Answered, error)
	GetSingle(uint, *model.Question) ([]*model.Answered, error)
	Post(*model.User, *model.Question, string) (*model.Answered, error)
	Viewed(*model.User, *model.Question) (*model.Answered, error)
}
