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

func (rr *answeredRepository) Get(u *model.User, q *model.Question) (map[uint]*model.Answered, error) {

	answered := make([]model.Answered, 0)
	err := rr.db.Model(u).Find(&answered).Error
	if err != nil {
		return nil, err
	}
	log.Info().Int("length", len(answered)).Msg("number of retrieved states")
	result := make(map[uint]*model.Answered, len(answered))
	for _, v := range answered {
		result[v.QuestionId] = &v
	}

	return result, nil
}

func (rr *answeredRepository) Post(u *model.User, q *model.Question) (*model.Answered, error) {
	// userExists := ur.db.NewRecord(u)
	answered := model.Answered{User: *u, QuestionId: q.ID, Answered: true}
	err := rr.db.Create(&answered).Error
	if err != nil {
		return nil, err
	}
	return &answered, nil
}
