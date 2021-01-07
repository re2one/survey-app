package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type questionController struct {
	questionRepository repository.QuestionRepository
	surveyRepository   repository.SurveyRepository
	answeredRepository repository.AnsweredRepository
	userRepository     repository.UserRepository
	piecesRepository   repository.PuzzleAnswerRepository
}

type QuestionResp struct {
	Questions []*model.Question `json:"questions"`
}

type SingleQuestionResp struct {
	Question *model.Question `json:"question"`
}

type AnswersResponse struct {
	Submissions map[string][]*model.PuzzleAnswer `json:"submissions"`
}

type QuestionController interface {
	GetAll(writer http.ResponseWriter, request *http.Request)
	Get(writer http.ResponseWriter, request *http.Request)
	GetAnswered(writer http.ResponseWriter, request *http.Request)
	Post(writer http.ResponseWriter, request *http.Request)
	Put(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}

func NewQuestionController(
	q repository.QuestionRepository,
	s repository.SurveyRepository,
	a repository.AnsweredRepository,
	u repository.UserRepository,
	p repository.PuzzleAnswerRepository,
) QuestionController {
	return &questionController{
		questionRepository: q,
		surveyRepository:   s,
		answeredRepository: a,
		userRepository:     u,
		piecesRepository:   p,
	}
}

func (uc *questionController) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	questions, err := uc.questionRepository.GetAll(v["surveyId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve questions.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(QuestionResp{Questions: questions})
	return
}

func (uc *questionController) Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	question, err := uc.questionRepository.Get(v["id"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve question.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(SingleQuestionResp{Question: question})
	return
}

func (uc *questionController) GetAnswered(writer http.ResponseWriter, request *http.Request) {
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

	// add survey id to query!!!
	state, err := uc.answeredRepository.Get(retrievedUser, v["surveyId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve the answered questions.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	submissions := make(map[string][]*model.PuzzleAnswer)

	for _, v := range state {
		questionId := strconv.Itoa(int(v.QuestionId))

		question, err := uc.questionRepository.Get(questionId)
		if err != nil {
			log.Error().Err(err).Msg("Unable to retrieve question.")
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if question.Type != "puzzle" {
			continue
		}

		pieces, err := uc.piecesRepository.GetUserSolution(email, questionId)
		if err != nil {
			log.Error().Err(err).Msg("Unable to retrieve the submitted answers for this puzzle question.")
			writer.WriteHeader(http.StatusInternalServerError)
		}

		submissions[questionId] = pieces
	}

	json.NewEncoder(writer).Encode(&AnswersResponse{Submissions: submissions})
	return
}

func (uc *questionController) Post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	survey, err := uc.surveyRepository.Get(v["surveyId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve survey.")
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
	question.Survey = *survey
	question2, err := uc.questionRepository.Post(v["surveyId"], &question)
	if err != nil {
		log.Error().Err(err).Msg("unable to write post question to db")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	r := SingleQuestionResp{Question: question2}
	json.NewEncoder(writer).Encode(r)
	return
}

func (uc *questionController) Put(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var question model.Question
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&question)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode question post body")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	question2, err := uc.questionRepository.Put(&question)
	if err != nil {
		log.Error().Err(err).Msg("unable to update question to db")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	r := SingleQuestionResp{Question: question2}
	json.NewEncoder(writer).Encode(r)
	return
}

func (uc *questionController) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	question, err := uc.questionRepository.Get(v["id"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve question.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	question2, err := uc.questionRepository.Delete(question)
	if err != nil {
		log.Error().Err(err).Msg("unable delete question from db")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	r := SingleQuestionResp{Question: question2}
	json.NewEncoder(writer).Encode(r)
	return
}
