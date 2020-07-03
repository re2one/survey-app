package presenter

import "survey-app-backend/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
	ResponseUser(u *model.User) *model.User
}