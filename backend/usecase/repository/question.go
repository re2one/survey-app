package repository

import "backend/model"

type QuestionRepository interface {
	GetAll(string) ([]*model.Question, error)
	Get(string) (*model.Question, error)
	Post(string, *model.Question) (*model.Question, error)
	Put(*model.Question) (*model.Question, error)
	Delete(*model.Question) (*model.Question, error)
}
