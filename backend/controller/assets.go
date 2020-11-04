package controller

import (
	"encoding/json"
	"net/http"

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
