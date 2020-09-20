package repository

import (
	"backend/model"
	"backend/usecase/repository"

	"github.com/jinzhu/gorm"
)

type answeredRepository struct {
	db *gorm.DB
}

func NewAnsweredRepository(db *gorm.DB) repository.AnsweredRepository {
	db.AutoMigrate(&model.Answered{})
	return &answeredRepository{db}
}

func (rr *answeredRepository) Get(u *model.User, q *model.Question) (*model.Answered, error) {

	var answered model.Answered
	err := rr.db.Model(&u).Model(q).Related(&answered).Error
	if err != nil {
		return nil, err
	}

	return &answered, nil
}

func (rr *answeredRepository) Post(u *model.User, q *model.Question) (*model.Answered, error) {
	// userExists := ur.db.NewRecord(u)
	answered := model.Answered{User: *u, Question: *q, Answered: true}
	err := rr.db.Create(&answered).Error
	if err != nil {
		return nil, err
	}
	return &answered, nil
}
