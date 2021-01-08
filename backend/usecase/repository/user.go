package repository

import "backend/model"

type UserRepository interface {
	Get(u *model.User) (*model.User, error)
	GetAll() ([]*model.User, error)
	GetIdFromEmail(string) (uint, error)
	Post(u *model.User) (*model.User, error)
}
