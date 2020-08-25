package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type surveyController struct {
	surveyRepository repository.SurveyRepository
}

type SurveyResp struct {
	Surveys []*model.Survey `json:"surveys"`
}

type SingleSurveyResp struct {
	Survey *model.Survey `json:"survey"`
}

// UserController defines functions available for requesat handling
type SurveyController interface {
	GetAll(writer http.ResponseWriter, request *http.Request)
	Get(writer http.ResponseWriter, request *http.Request)
	Post(writer http.ResponseWriter, request *http.Request)
	Put(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}

// NewUserController provides functions for request handling
func NewSurveyController(repo repository.SurveyRepository) SurveyController {
	return &surveyController{surveyRepository: repo}
}

func (uc *surveyController) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	surveys, err := uc.surveyRepository.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve surveys.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(SurveyResp{Surveys: surveys})
	return
}

func (uc *surveyController) Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	survey, err := uc.surveyRepository.Get(v["title"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve survey.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(SingleSurveyResp{Survey: survey})
	return
}

func (uc *surveyController) Post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var survey model.Survey
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&survey)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode survey post body")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	survey2, err := uc.surveyRepository.Post(&survey)
	if err != nil {
		log.Error().Err(err).Msg("unable to write post survey to db")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	r := SingleSurveyResp{Survey: survey2}
	json.NewEncoder(writer).Encode(r)
	return
}

func (uc *surveyController) Put(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// xyz := &Resp{Message: "hey everybody, shit is authed!!"}
	json.NewEncoder(writer).Encode(nil)
	return
}

func (uc *surveyController) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// xyz := &Resp{Message: "hey everybody, shit is authed!!"}
	json.NewEncoder(writer).Encode(nil)
	return
}
