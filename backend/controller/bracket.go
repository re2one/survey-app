package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type bracketController struct {
	bracketRepository repository.BracketRepository
	surveyRepository  repository.SurveyRepository
}

type BracketResp struct {
	Brackets []*model.Bracket `json:"brackets"`
}

type SingleBracketResp struct {
	Bracket *model.Bracket `json:"bracket"`
}

type BracketController interface {
	Get(writer http.ResponseWriter, request *http.Request)
	Post(writer http.ResponseWriter, request *http.Request)
}

func NewBracketController(b repository.BracketRepository, s repository.SurveyRepository) BracketController {
	return &bracketController{bracketRepository: b, surveyRepository: s}
}

func (bc *bracketController) Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	brackets, err := bc.bracketRepository.Get(v["surveyId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve choice.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(BracketResp{Brackets: brackets})
	return
}

func (bc *bracketController) Post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	survey, err := bc.surveyRepository.Get(v["surveyId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve survey.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var bracket model.Bracket
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&bracket)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode choice post body")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	bracket.Survey = *survey
	newBracket, err := bc.bracketRepository.Post(v["surveyId"], &bracket)
	if err != nil {
		log.Error().Err(err).Msg("unable to write post choice to db")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	r := SingleBracketResp{Bracket: newBracket}
	json.NewEncoder(writer).Encode(r)
	return
}
