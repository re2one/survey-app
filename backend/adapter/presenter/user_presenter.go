package presenter

import (
	"backend/common"
	"backend/model"
	"backend/model/response"
)

type userPresenter struct {
	auth *common.Auth
}

// fooo
type UserPresenter interface {
	LoginResponse(us *model.User, r *model.Role) (*response.UserResponse, error)
	SignupResponse(us *model.User, r *model.Role) (*response.UserResponse, error)
}

// exported function
func NewUserPresenter(a *common.Auth) UserPresenter {
	return &userPresenter{a}
}

func (up *userPresenter) LoginResponse(user *model.User, r *model.Role) (*response.UserResponse, error) {
	token, err := (*up.auth).CreateToken(user, r)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{Role: r.Role, Username: user.Name, Token: token}, nil
}

func (up *userPresenter) SignupResponse(user *model.User, r *model.Role) (*response.UserResponse, error) {
	token, err := (*up.auth).CreateToken(user, r)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{Role: r.Role, Username: user.Name, Token: token}, nil
}
