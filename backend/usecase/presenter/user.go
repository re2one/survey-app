package presenter

import (
	"backend/model"
	"backend/model/response"
)

type UserPresenter interface {
	SignupResponse(u *model.User, r *model.Role) (*response.UserResponse, error)
	LoginResponse(u *model.User, r *model.Role) (*response.UserResponse, error)
}
