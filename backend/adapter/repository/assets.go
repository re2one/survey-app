package repository

import (
	"fmt"
	"os"

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
	surveyPath, err := a.createDirectory("assets", "survey", surveyId)
	if err != nil {
		return err
	}
	questionPath, err := a.createDirectory(surveyPath, "question", questionId)
	if err != nil {
		return err
	}
	log.Info().Str("question path", questionPath).Msg("Asset-Folder Created")
	return nil
}

func (a *assetsRepository) createDirectory(path string, dirType string, id string) (string, error) {

	dirPath := fmt.Sprintf("%v/%v_%v", path, dirType, id)
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirPath, 0755)
		if errDir != nil {
			return "", errDir
		}
	}
	return dirPath, nil
}
