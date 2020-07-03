package interactor

import (
	"survey-app-backend/model"
	"survey-app-backend/usecase/presenter"
	"survey-app-backend/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	GetAll(u []*model.User) ([]*model.User, error)
	Get(u *model.User) (*model.User, error)
	Update(u *model.User) (*model.User, error)
	Add(u *model.User) (*model.User, error)
	Delete(u *model.User) (*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) GetAll(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userInteractor) Get(u *model.User) (*model.User, error) {
	u, err := us.UserRepository.Find(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}

func (us *userInteractor) Add(u *model.User) (*model.User, error) {
	u, err := us.UserRepository.Add(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}

func (us *userInteractor) Update(u *model.User) (*model.User, error) {
	u, err := us.UserRepository.Update(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}

func (us *userInteractor) Delete(u *model.User) (*model.User, error) {
	u, err := us.UserRepository.Delete(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}