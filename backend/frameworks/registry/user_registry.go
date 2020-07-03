package registry

import (
	"survey-app-backend/adapter/controller"
	ip "survey-app-backend/adapter/presenter"
	ir "survey-app-backend/adapter/repository"
	"survey-app-backend/usecase/interactor"
	up "survey-app-backend/usecase/presenter"
	ur "survey-app-backend/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}