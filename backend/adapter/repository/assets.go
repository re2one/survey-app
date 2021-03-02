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

func (a *assetsRepository) PostAssetFolder(surveyId string, dirType string) error {
	surveyPath, err := a.createDirectory("assets", "survey", surveyId)
	if err != nil {
		return err
	}
	introPath, err := a.createDirectoryForIntroduction(surveyPath, dirType)
	if err != nil {
		return err
	}
	log.Info().Str("question path", introPath).Msg("Asset-Folder Created")
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

func (a *assetsRepository) createDirectoryForIntroduction(path string, dirType string) (string, error) {

	dirPath := fmt.Sprintf("%v/%v", path, dirType)
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

func (a *assetsRepository) SavePDF(filetype string, surveyId string, data []byte) error {
	var path string
	switch filetype {
	case "termsandconditions":
		path = "assets/termsandconditions.pdf"
	case "impressum":
		path = "assets/impressum.pdf"
	case "datenschutz":
		path = "assets/datenschutz.pdf"
	default:
		path = fmt.Sprintf("assets/survey_%v/%v/%v.pdf", surveyId, filetype, filetype)
	}
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(data)
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

func (a *assetsRepository) Get(surveyId string, questionId string, filename string) (*os.File, error) {

	path := fmt.Sprintf("assets/survey_%v/question_%v/%v", surveyId, questionId, filename)
	img, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer img.Close()
	return img, nil
}
