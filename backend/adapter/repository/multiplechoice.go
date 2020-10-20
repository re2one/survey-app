package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"backend/model"
	"backend/usecase/repository"
)

type multipleChoiceRepository struct {
	db *gorm.DB
}

func NewMultipleChoiceRepository(db *gorm.DB) repository.MultipleChoiceRepository {
	db.AutoMigrate(&model.ChoiceAnswer{})
	return &multipleChoiceRepository{db}
}

func (sr *multipleChoiceRepository) Post(multi *model.ChoiceAnswer) (*model.ChoiceAnswer, error) {

	var m2 model.ChoiceAnswer
	err := sr.db.Where("questionId = ? and email = ?", multi.QuestionId, multi.Email).First(&m2).Error

	if err == nil {
		err = errors.New("multiple choice answer already exists")
		return nil, err
	}

	sr.db.Create(&multi)
	return multi, nil
}

func (sr *multipleChoiceRepository) Get(questionId uint, email string) (*model.ChoiceAnswer, error) {
	var m2 model.ChoiceAnswer
	err := sr.db.Where("question_id = ? and email = ?", questionId, email).First(&m2).Error

	if err != nil {
		err = errors.New("multiple choice does not exist")
		return nil, err
	}

	return &m2, nil
}
