package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type surveyController struct {
	surveyRepository repository.SurveyRepository
}

type Resp struct {
	Message string `json:"message"`
}
type SurveyResp struct {
	Surveys []*model.Survey `json:"surveys"`
}

// UserController defines functions available for requesat handling
type SurveyController interface {
	Test(writer http.ResponseWriter, request *http.Request)
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

func (uc *surveyController) Test(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	xyz := &Resp{Message: "hey everybody, shit is authed!!"}
	json.NewEncoder(writer).Encode(xyz)
	return
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

	uc.surveyRepository.GetAll()
	json.NewEncoder(writer).Encode(nil)
	return
}

func (uc *surveyController) Post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	xyz := &Resp{Message: "hey everybody, shit is authed!!"}
	json.NewEncoder(writer).Encode(xyz)
	return
}

func (uc *surveyController) Put(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	xyz := &Resp{Message: "hey everybody, shit is authed!!"}
	json.NewEncoder(writer).Encode(xyz)
	return
}

func (uc *surveyController) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	xyz := &Resp{Message: "hey everybody, shit is authed!!"}
	json.NewEncoder(writer).Encode(xyz)
	return
}
