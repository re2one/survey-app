package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/model/response"
	"backend/usecase/repository"
)

type fullQuestionsController struct {
	questionRepository       repository.QuestionRepository
	answeredRepository       repository.AnsweredRepository
	surveyRepository         repository.SurveyRepository
	userRepository           repository.UserRepository
	multiplechoiceRepository repository.MultipleChoiceRepository
	choiceRepository         repository.ChoiceRepository
}

type FullQuestionsController interface {
	GetAll(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

func NewFullQuestionsController(
	questionRepository repository.QuestionRepository,
	answeredRepository repository.AnsweredRepository,
	surveyRepository repository.SurveyRepository,
	userRepository repository.UserRepository,
	multiplechoiceRepository repository.MultipleChoiceRepository,
	choiceRepository repository.ChoiceRepository,
) FullQuestionsController {
	return &fullQuestionsController{
		questionRepository:       questionRepository,
		surveyRepository:         surveyRepository,
		answeredRepository:       answeredRepository,
		userRepository:           userRepository,
		multiplechoiceRepository: multiplechoiceRepository,
		choiceRepository:         choiceRepository,
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

	fullQuestions := make([]*response.FullQuestion, 0)

	currentQuestion, err := uc.questionRepository.GetFirst(surveyId)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve first question.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	finished := false
	for answered := true; answered; {

		state, err := uc.answeredRepository.Get(retrievedUser, currentQuestion)
		if err != nil {
			log.Error().Err(err).Msg("Unable to retrieve if current question has been answered.")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, ok := state[currentQuestion.ID]; !ok {
			answered = false
		}

		fullQuestion := response.FullQuestion{
			QuestionId: currentQuestion.ID,
			Title:      currentQuestion.Title,
			Type:       currentQuestion.Type,
			Answered:   answered,
		}

		fullQuestions = append(fullQuestions, &fullQuestion)

		if _, ok := state[currentQuestion.ID]; !ok {
			continue
		}

		userAnswer, err := uc.multiplechoiceRepository.Get(currentQuestion.ID, email)
		if err != nil {
			log.Error().Err(err).Msg("Unable to fetch the users answer.")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		usedChoice, err := uc.choiceRepository.GetByText(currentQuestion.ID, userAnswer.Text)
		if err != nil {
			log.Error().Err(err).Msg("Unable to fetch the corresponding choice for an answer.")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		currentQuestion, err = uc.questionRepository.Get(fmt.Sprint(usedChoice.NextQuestion))
		if err != nil {
			log.Error().Err(err).Msg("Unable to retrieve the following question question.")
			finished = true
			answered = false
		}
	}

	json.NewEncoder(writer).Encode(response.FullQuestionsResponse{
		Questions: fullQuestions,
		Finished:  finished,
	})
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