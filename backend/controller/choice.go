package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type choiceController struct {
	choiceRepository   repository.ChoiceRepository
	questionRepository repository.QuestionRepository
}

type ChoiceResp struct {
	Choices []*model.Choice `json:"choices"`
}

type SingleChoiceResp struct {
	Choice *model.Choice `json:"choice"`
}

type ChoiceController interface {
	GetAll(writer http.ResponseWriter, request *http.Request)
	Get(writer http.ResponseWriter, request *http.Request)
	Post(writer http.ResponseWriter, request *http.Request)
	Put(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}

func NewChoiceController(c repository.ChoiceRepository, q repository.QuestionRepository) ChoiceController {
	return &choiceController{choiceRepository: c, questionRepository: q}
}

func (uc *choiceController) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	choices, err := uc.choiceRepository.GetAll(v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve choices.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(ChoiceResp{Choices: choices})
	return
}

func (uc *choiceController) Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	choice, err := uc.choiceRepository.Get(v["id"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve choice.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(SingleChoiceResp{Choice: choice})
	return
}

func (uc *choiceController) Post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	question, err := uc.questionRepository.Get(v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve survey.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var choice model.Choice
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&choice)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode choice post body")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	choice.Question = *question
	choice2, err := uc.choiceRepository.Post(v["surveyId"], &choice)
	if err != nil {
		log.Error().Err(err).Msg("unable to write post choice to db")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	r := SingleChoiceResp{Choice: choice2}
	json.NewEncoder(writer).Encode(r)
	return
}

func (uc *choiceController) Put(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var choice model.Choice
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&choice)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode choice post body")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	choice2, err := uc.choiceRepository.Put(&choice)
	if err != nil {
		log.Error().Err(err).Msg("unable to update choice to db")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	r := SingleChoiceResp{Choice: choice2}
	json.NewEncoder(writer).Encode(r)
	return
}

func (uc *choiceController) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	choice, err := uc.choiceRepository.Get(v["id"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve choice.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	choice2, err := uc.choiceRepository.Delete(choice)
	if err != nil {
		log.Error().Err(err).Msg("unable delete choice from db")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	r := SingleChoiceResp{Choice: choice2}
	json.NewEncoder(writer).Encode(r)
	return
}
