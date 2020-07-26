package interactor

// import (
// 	"backend/model"
// 	"backend/usecase/presenter"
// 	"backend/usecase/repository"
// )

// type roleInteractor struct {
// 	RoleRepository repository.roleRepository
// 	// UserPresenter  presenter.UserPresenter
// }

// type UserInteractor interface {
// 	Get(u *model.User) (*model.User, error)
// 	Post(u *model.User) (*model.User, error)
// }

// func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
// 	return &userInteractor{r, p}
// }

// func (us *userInteractor) Get(u *model.User) (*model.User, error) {
// 	u, err := us.UserRepository.Get(u)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return us.UserPresenter.LoginResponse(u), nil
// }

// func (us *userInteractor) Post(u *model.User) (*model.User, error) {
// 	u, err := us.UserRepository.Post(u)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return us.UserPresenter.SignupResponse(u), nil
// }
