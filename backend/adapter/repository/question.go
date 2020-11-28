package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"backend/model"
	"backend/usecase/repository"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) repository.QuestionRepository {
	db.AutoMigrate(&model.Question{})
	return &questionRepository{db}
}

func (sr *questionRepository) Get(id string) (*model.Question, error) {
	//check if record exists
	var s model.Question
	err := sr.db.Where("ID = ?", id).First(&s).Error
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (sr *questionRepository) GetFirst(survey string) (*model.Question, error) {
	//check if record exists
	var s model.Question
	err := sr.db.Where("survey_id = ? and first = 'true'", survey).First(&s).Error
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (sr *questionRepository) GetAll(surveyId string) ([]*model.Question, error) {
	//check if record exists
	questions := make([]*model.Question, 0)
	if err := sr.db.Where("survey_id = ?", surveyId).Find(&questions).Error; err != nil {
		// log.Fatalln(err)
		return nil, err
	}

	return questions, nil
}

func (sr *questionRepository) Post(surveyId string, s *model.Question) (*model.Question, error) {

	err := sr.db.Where("title = ? and survey_id = ?", s.Title, surveyId).First(&s).Error

	if err == nil {
		err = errors.New("question already exists")
		return nil, err
	}

	sr.db.Create(&s)
	return s, nil
}

func (sr *questionRepository) Put(s *model.Question) (*model.Question, error) {

	var question model.Question
	err := sr.db.Where("ID = ?", s.ID).First(&question).Error

	if err != nil {
		err = errors.New("question does not exists")
		return nil, err
	}
	question.Title = s.Title
	question.Text = s.Text
	question.Type = s.Type
	question.First = s.First
	question.Bracket = s.Bracket
	sr.db.Save(question)
	return &question, nil
}

func (sr *questionRepository) Delete(s *model.Question) (*model.Question, error) {

	err := sr.db.Where("ID = ?", s.ID).First(&s).Error

	if err != nil {
		err = errors.New("No question to delete")
		return nil, err
	}
	sr.db.Delete(s)
	return s, nil
}
