package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/model/response"
	"backend/usecase/repository"
)

type fullQuestionsController struct {
	questionRepository repository.QuestionRepository
	answeredRepository repository.AnsweredRepository
	userRepository     repository.UserRepository
}

type FullQuestionsController interface {
	GetAll(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

func NewFullQuestionsController(
	questionRepository repository.QuestionRepository,
	answeredRepository repository.AnsweredRepository,
	userRepository repository.UserRepository,
) FullQuestionsController {
	return &fullQuestionsController{
		questionRepository: questionRepository,
		answeredRepository: answeredRepository,
		userRepository:     userRepository,
	}
}

func (uc *fullQuestionsController) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	surveyId := v["surveyid"]
	email := v["email"]
	lookupUser := model.User{Email: email}

	retrievedUser, err := uc.userRepository.Get(&lookupUser)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve user.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	questions, err := uc.questionRepository.GetAll(surveyId)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve questions.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	fullQuestions := make([]*response.FullQuestion, 0, len(questions))
	for _, q := range questions {

		answered := false
		state, err := uc.answeredRepository.Get(retrievedUser, q)
		if err == nil {
			answered = state.Answered
		}
		log.Error().Err(err).Msg("No info for question x user pair. question is unanswered.")
		fullQuestion := response.FullQuestion{
			QuestionId: q.ID,
			Title:      q.Title,
			Type:       q.Type,
			Answered:   answered,
		}
		fullQuestions = append(fullQuestions, &fullQuestion)
	}

	json.NewEncoder(writer).Encode(response.FullQuestionsResponse{Questions: fullQuestions})
	return
}

func (uc *fullQuestionsController) Post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	email := v["email"]
	lookupUser := model.User{Email: email}

	retrievedUser, err := uc.userRepository.Get(&lookupUser)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve user.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var question model.Question
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&question)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode question post body")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	answered, err := uc.answeredRepository.Post(retrievedUser, &question)
	if err != nil {
		log.Error().Err(err).Msg("unable post question state")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(answered)
	return
}
