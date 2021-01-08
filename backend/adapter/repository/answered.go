package repository

import (
	"github.com/rs/zerolog/log"

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

func (rr *answeredRepository) Get(u *model.User, surveyId string) (map[uint]model.Answered, error) {

	answered := make([]model.Answered, 0)

	err := rr.db.Where("user_id = ? and answered = ? and survey_id = ?", u.ID, true, surveyId).Find(&answered).Error
	if err != nil {
		return nil, err
	}
	log.Info().Int("length", len(answered)).Msg("number of retrieved states")
	result := make(map[uint]model.Answered, len(answered))
	for _, v := range answered {
		result[v.QuestionId] = v
	}

	return result, nil
}

func (rr *answeredRepository) GetSingle(u *model.User, q *model.Question) ([]*model.Answered, error) {

	answered := make([]*model.Answered, 0)
	err := rr.db.Where("user_id = ? and question_id = ?", u.ID, q.ID).Find(&answered).Error
	if err != nil {
		return nil, err
	}
	log.Info().Int("length", len(answered)).Msg("number of retrieved states")

	return answered, nil
}

func (rr *answeredRepository) Post(u *model.User, q *model.Question) (*model.Answered, error) {
	// userExists := ur.db.NewRecord(u)
	var answered model.Answered
	err := rr.db.Where("user_id = ? and question_id = ?", u.ID, q.ID).Find(&answered).Error
	if err != nil {
		answered := model.Answered{User: *u, QuestionId: q.ID, Answered: true, SurveyId: q.SurveyId}
		err := rr.db.Create(&answered).Error
		if err != nil {
			return nil, err
		}
		return &answered, nil
	}
	answered.Answered = true
	rr.db.Save(answered)
	return &answered, nil

}

func (rr *answeredRepository) Viewed(u *model.User, q *model.Question) (*model.Answered, error) {
	// userExists := ur.db.NewRecord(u)
	answered := model.Answered{User: *u, QuestionId: q.ID, Viewed: true, Answered: false, SurveyId: q.SurveyId}
	err := rr.db.Create(&answered).Error
	if err != nil {
		return nil, err
	}
	return &answered, nil
}
