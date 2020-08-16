package repository

import "backend/model"

type RoleRepository interface {
	Get(u *model.User) (*model.Role, error)
	Post(u *model.User, r string) (*model.Role, error)
}
