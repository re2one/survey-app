package repository

import (
	"backend/model"
	"backend/usecase/repository"

	"fmt"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
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

func NewUserRepository(db *gorm.DB) repository.UserRepository {
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

func (ur *userRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) Post(u *model.User) (*model.User, error) {
	// userExists := ur.db.NewRecord(u)
	// var user model.User
	err := ur.db.Where("email = ?", u.Email).First(&u).Error

	if err == nil {
		err = UserAlreadyExistsError{u}
		return nil, err
	}

	ur.db.Create(&u)
	return u, nil
}

func (ur *userRepository) GetIdFromEmail(email string) (uint, error) {
	var user model.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
