package repository

import (
	"github.com/jinzhu/gorm"
	"survey-app-backend/model"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Find(u *model.User) (*model.User, error)
	Add(u *model.User) (*model.User, error)
	Update(u *model.User) (*model.User, error)
	Delete(u *model.User) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Find(u *model.User) (*model.User, error) {
	err := ur.db.Where("name = ?", u.Name).First(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Add(u *model.User) (*model.User, error) {
	userExists := ur.db.NewRecord(u)
	var err error
	if userExists != true {
		ur.db.Create(&u)

	}

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Update(u *model.User) (*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Delete(u *model.User) (*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}