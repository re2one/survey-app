package presenter

import "survey-app-backend/model"

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
	ResponseUser(us *model.User) *model.User
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	// for _, u := range us {
	// 	u.Name = "Mr." + u.Name
	// }
	return us
}

func (up *userPresenter) ResponseUser(u *model.User) *model.User {
	// for _, u := range us {
	// 	u.Name = "Mr." + u.Name
	// }
	return u
}