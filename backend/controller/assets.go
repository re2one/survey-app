package controller

import (
	"encoding/json"
	"image/jpeg"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/usecase/repository"
)

type assetsController struct {
	assetsRepository repository.AssetsRepository
}

type SingleAssetResponse struct {
	SurveyId   string `json:"surveyId"`
	QuestionId string `json:"questionId"`
}

type AssetsController interface {
	Post(writer http.ResponseWriter, request *http.Request)
	Upload(writer http.ResponseWriter, request *http.Request)
}

func NewAssetsController(a repository.AssetsRepository) AssetsController {
	return &assetsController{assetsRepository: a}
}

func (a *assetsController) Post(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	err := a.assetsRepository.Post(v["surveyId"], v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("unable to create asset folder")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	r := SingleAssetResponse{SurveyId: v["surveyId"], QuestionId: v["questionId"]}
	json.NewEncoder(writer).Encode(r)
	return
}

func (a *assetsController) Upload(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	request.ParseMultipartForm(32 << 20) // limit your max input length!

	// in your case file would be fileupload
	file, header, err := request.FormFile("fileKey")
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to unwrap file")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")
	log.Info().Str("Filename", name[0]).Msg("Name of transferred file")

	image, err := jpeg.Decode(file)

	// save file to disk
	err = a.assetsRepository.SaveFile(v["surveyId"], v["questionId"], image, header.Filename)
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to write file to disc")
		writer.WriteHeader(http.StatusInternalServerError)

	}

	return
}
