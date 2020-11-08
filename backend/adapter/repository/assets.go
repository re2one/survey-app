package repository

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
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

func (a *assetsRepository) Upload(surveyId string, questionId string) error {
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

func (a *assetsRepository) SaveFile(surveyId string, questionId string, data image.Image, filename string) error {
	path := fmt.Sprintf("assets/survey_%v/question_%v/%v", surveyId, questionId, filename)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, data, nil)
	return nil
}

func (a *assetsRepository) GetFilenames(surveyId string, questionId string) ([]string, error) {
	path := fmt.Sprintf("assets/survey_%v/question_%v", surveyId, questionId)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	result := make([]string, len(files))
	for i, f := range files {
		result[i] = f.Name()
	}
	return result, nil
}
