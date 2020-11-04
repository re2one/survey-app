package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"

	"backend/usecase/repository"
)

type assetsRepository struct {
	db *gorm.DB
}

func NewAssetsRepository(db *gorm.DB) repository.AssetsRepository {
	// db.AutoMigrate(&model.Answered{})
	return &assetsRepository{db}
}

func (a *assetsRepository) Post(surveyId string, questionId string) error {
	log.Info().Str("SurveyId", surveyId).Str("QuestionId", questionId).Msg("Asset-Folder Created")
	return nil
}
