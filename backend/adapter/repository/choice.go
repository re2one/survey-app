package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"backend/model"
	"backend/usecase/repository"
)

type choiceRepository struct {
	db *gorm.DB
}

func NewChoiceRepository(db *gorm.DB) repository.ChoiceRepository {
	db.AutoMigrate(&model.Choice{})
	return &choiceRepository{db}
}

func (sr *choiceRepository) Get(title string) (*model.Choice, error) {
	//check if record exists
	var s model.Choice
	err := sr.db.Where("ID = ?", title).First(&s).Error
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (sr *choiceRepository) GetByText(questionId uint, text string) (*model.Choice, error) {
	//check if record exists
	var s model.Choice
	err := sr.db.Where("question_id = ? and text = ?", questionId, text).First(&s).Error
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (sr *choiceRepository) GetAll(questionId string) ([]*model.Choice, error) {
	//check if record exists
	choices := make([]*model.Choice, 0)
	if err := sr.db.Where("question_id = ?", questionId).Find(&choices).Error; err != nil {
		log.Fatalln(err)
	}

	return choices, nil
}

func (sr *choiceRepository) Post(questionId string, s *model.Choice) (*model.Choice, error) {

	err := sr.db.Where("text = ? and question_id = ?", s.Text, questionId).First(&s).Error

	if err == nil {
		err = errors.New("choice already exists")
		return nil, err
	}

	sr.db.Create(&s)
	return s, nil
}

func (sr *choiceRepository) Put(s *model.Choice) (*model.Choice, error) {

	var choice model.Choice
	err := sr.db.Where("ID = ?", s.ID).First(&choice).Error

	if err != nil {
		err = errors.New("choice does not exists")
		return nil, err
	}
	choice.Text = s.Text
	choice.NextQuestion = s.NextQuestion
	sr.db.Save(choice)
	return &choice, nil
}

func (sr *choiceRepository) Delete(s *model.Choice) (*model.Choice, error) {

	err := sr.db.Where("ID = ?", s.ID).First(&s).Error

	if err != nil {
		err = errors.New("No choice to delete")
		return nil, err
	}
	sr.db.Delete(s)
	return s, nil
}
