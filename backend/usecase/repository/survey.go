package repository

import "backend/model"

type SurveyRepository interface {
	GetAll() ([]*model.Survey, error)
	Get(string) (*model.Survey, error)
	Post(*model.Survey) (*model.Survey, error)
	Put(*model.Survey) (*model.Survey, error)
	Delete(*model.Survey) (*model.Survey, error)
}
