package interactor

import (
	"backend/model"
	"backend/model/response"
	"backend/usecase/presenter"
	"backend/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(u *model.User) (*response.UserResponse, error)
	Post(u *model.User, role string) (*response.UserResponse, error)
}

func NewUserInteractor(
	ur repository.UserRepository,
	rr repository.RoleRepository,
	p presenter.UserPresenter,
) UserInteractor {
	return &userInteractor{ur, rr, p}
}

func (us *userInteractor) Get(u *model.User) (*response.UserResponse, error) {

	var r *model.Role
	var err error
	var res *response.UserResponse

	u, err = us.UserRepository.Get(u)
	if err != nil {
		return nil, err
	}

	r, err = us.RoleRepository.Get(u)
	if err != nil {
		return nil, err
	}

	// hhhhh
	res, err = us.UserPresenter.LoginResponse(u, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (us *userInteractor) Post(u *model.User, role string) (*response.UserResponse, error) {
	var err error
	var res *response.UserResponse
	var r *model.Role

	u, err = us.UserRepository.Post(u)
	if err != nil {
		return nil, err
	}

	r, err = us.RoleRepository.Post(u, role)
	if err != nil {
		return nil, err
	}

	// hhhhh
	res, err = us.UserPresenter.SignupResponse(u, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}
