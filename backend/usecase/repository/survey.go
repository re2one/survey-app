package repository

import "backend/model"

type SurveyRepository interface {
	GetAll() ([]*model.Survey, error)
	Get(u *model.Survey) (*model.Survey, error)
	Post(u *model.Survey) (*model.Survey, error)
	Put(u *model.Survey) (*model.Survey, error)
	Delete(u *model.Survey) (*model.Survey, error)
}
