package registry

import (
	"backend/adapter/controller"
	ip "backend/adapter/presenter"
	ir "backend/adapter/repository"
	"backend/usecase/interactor"
	up "backend/usecase/presenter"
	ur "backend/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewRoleRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter(r.auth)
}

func (r *registry) NewRoleRepository() ur.RoleRepository {
	return ir.NewRoleRepository(r.db)
}
