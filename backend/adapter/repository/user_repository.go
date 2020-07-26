package repository

import (
	"backend/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Get(u *model.User) (*model.User, error)
	Post(u *model.User) (*model.User, error)
}

type UserAlreadyExistsError struct {
	user *model.User
}

func (u UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("User %+v already exists.\n", u.user)
}

type UserNotFoundError struct {
	user *model.User
}

func (u UserNotFoundError) Error() string {
	return fmt.Sprintf("User %+v not found.\n", u.user)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&model.User{})
	return &userRepository{db}
}

func (ur *userRepository) Get(u *model.User) (*model.User, error) {
	//check if record exists
	err := ur.db.Where("email = ?", u.Email).First(&u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Post(u *model.User) (*model.User, error) {
	// userExists := ur.db.NewRecord(u)
	// var user model.User
	err := ur.db.Where("email = ?", u.Email).First(&u).Error

	if err != nil {
		ur.db.Create(&u)
	} else {
		err = UserAlreadyExistsError{u}
		return nil, err
	}
	return u, nil
}
