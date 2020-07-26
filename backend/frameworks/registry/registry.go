package registry

import (
	"backend/adapter/controller"
	"backend/common"

	"github.com/jinzhu/gorm"
)

type registry struct {
	db   *gorm.DB
	auth *common.Auth
}

// Registry xyz
type Registry interface {
	NewUserHandler() controller.UserController
}

// registry xyz
func NewRegistry(db *gorm.DB, auth *common.Auth) Registry {
	return registry{db, auth}
}

func (r registry) NewUserHandler() controller.UserController {
	return r.NewUserController()
}
