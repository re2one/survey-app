package repository

import (
	"backend/model"

	"github.com/jinzhu/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

type RoleRepository interface {
	Get(u *model.User) (*model.Role, error)
	Post(u *model.User, r string) (*model.Role, error)
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	db.AutoMigrate(&model.Role{})
	return &roleRepository{db}
}

func (rr *roleRepository) Get(u *model.User) (*model.Role, error) {

	var role model.Role
	err := rr.db.Model(&u).Related(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (rr *roleRepository) Post(u *model.User, r string) (*model.Role, error) {
	// userExists := ur.db.NewRecord(u)
	role := model.Role{User: *u, Role: r}
	err := rr.db.Create(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
