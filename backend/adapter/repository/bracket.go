package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"backend/model"
	"backend/usecase/repository"
)

type bracketRepository struct {
	db *gorm.DB
}

func NewBracketRepository(db *gorm.DB) repository.BracketRepository {
	db.AutoMigrate(&model.Bracket{})
	return &bracketRepository{db}
}

func (bc *bracketRepository) Get(surveyId string) ([]*model.Bracket, error) {
	//check if record exists
	brackets := make([]*model.Bracket, 0)
	if err := bc.db.Where("survey_id = ?", surveyId).Find(&brackets).Error; err != nil {
		return nil, err
	}

	return brackets, nil
}

func (br *bracketRepository) Post(surveyId string, b *model.Bracket) (*model.Bracket, error) {

	err := br.db.Where("name = ? and survey_id = ?", b.Name, surveyId).First(&b).Error

	if err == nil {
		err = errors.New("choice already exists")
		return nil, err
	}

	br.db.Create(&b)
	return b, nil
}
