package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"backend/model"
	"backend/usecase/repository"
)

type surveyRepository struct {
	db *gorm.DB
}

func NewSurveyRepository(db *gorm.DB) repository.SurveyRepository {
	db.AutoMigrate(&model.Survey{})
	return &surveyRepository{db}
}

func (sr *surveyRepository) Get(title string) (*model.Survey, error) {
	//check if record exists
	var s model.Survey
	err := sr.db.Where("ID = ?", title).First(&s).Error
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (sr *surveyRepository) GetAll() ([]*model.Survey, error) {
	//check if record exists
	surveys := make([]*model.Survey, 0)
	if err := sr.db.Find(&surveys).Error; err != nil {
		log.Fatalln(err)
	}

	return surveys, nil
}

func (sr *surveyRepository) Post(s *model.Survey) (*model.Survey, error) {

	err := sr.db.Where("title = ?", s.Title).First(&s).Error

	if err == nil {
		err = errors.New("Survey already exists")
		return nil, err
	}

	sr.db.Create(&s)
	return s, nil
}

func (sr *surveyRepository) Put(s *model.Survey) (*model.Survey, error) {

	var survey model.Survey
	err := sr.db.Where("ID = ?", s.Title).First(&survey).Error

	if err != nil {
		err = errors.New("Survey does not exists")
		return nil, err
	}
	survey.Summary = s.Summary
	survey.Disclaimer = s.Disclaimer
	survey.Introduction = s.Introduction
	survey.Title = s.Title
	sr.db.Save(survey)
	return &survey, nil
}

func (sr *surveyRepository) Delete(s *model.Survey) (*model.Survey, error) {

	err := sr.db.Where("ID = ?", s.Title).First(&s).Error

	if err != nil {
		err = errors.New("No Survey to delete")
		return nil, err
	}
	sr.db.Delete(s)
	return s, nil
}
