package presenter

import (
	"backend/model"
	"backend/model/response"
)

type UserPresenter interface {
	SignupResponse(u *model.User) *model.User
	LoginResponse(u *model.User, r *model.Role) (*response.UserResponse, error)
}
