package repository

import "survey-app-backend/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Find(u *model.User) (*model.User, error)
	Update(u *model.User) (*model.User, error)
	Add(u *model.User) (*model.User, error)
	Delete(u *model.User) (*model.User, error)
}