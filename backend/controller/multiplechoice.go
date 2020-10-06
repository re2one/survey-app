package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type multipleChoiceController struct {
	multipleChoiceRepository repository.MultipleChoiceRepository
	questionRepository       repository.QuestionRepository
}

type MultipleChoiceController interface {
	Post(writer http.ResponseWriter, request *http.Request)
}

func NewMultipleChoiceController(c repository.MultipleChoiceRepository, q repository.QuestionRepository) MultipleChoiceController {
	return &multipleChoiceController{multipleChoiceRepository: c, questionRepository: q}
}

func (m *multipleChoiceController) Post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var choice model.ChoiceAnswer
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&choice)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode multiplechoice answer post body")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	choice2, err := m.multipleChoiceRepository.Post(&choice)
	if err != nil {
		log.Error().Err(err).Msg("unable to write post choice to db")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(writer).Encode(choice2)
	return
}
