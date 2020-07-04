package repository

import (
	"fmt"
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

type UserAlreadyExistsError struct{
	user *model.User
}

func (uaee UserAlreadyExistsError) Error () string {
	return fmt.Sprintf("User %+v already exists.\n", uaee.user)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&model.User{})
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
	// userExists := ur.db.NewRecord(u)
	var foundUsers []model.User
	ur.db.Where("email = ?", u.Email).Find(&foundUsers)

	if len(foundUsers) < 1 {
		ur.db.Create(&u)
	} else {
		err := UserAlreadyExistsError{u}
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