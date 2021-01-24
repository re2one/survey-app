package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type puzzleController struct {
	puzzleRepository   repository.PuzzleRepository
	answeredRepository repository.AnsweredRepository
	userRepository     repository.UserRepository
	questionRepository repository.QuestionRepository
}

type PuzzleController interface {
	Put(writer http.ResponseWriter, request *http.Request)
	GetAll(writer http.ResponseWriter, request *http.Request)
	GetAllForQuestionaire(writer http.ResponseWriter, request *http.Request)
}

type PuzzleResponse struct {
	Pieces []*model.Puzzlepiece `json:"pieces"`
}

func NewPuzzleController(
	repo repository.PuzzleRepository,
	answeredRepository repository.AnsweredRepository,
	userRepository repository.UserRepository,
	questionRepository repository.QuestionRepository,
) PuzzleController {
	return &puzzleController{
		puzzleRepository:   repo,
		answeredRepository: answeredRepository,
		userRepository:     userRepository,
		questionRepository: questionRepository,
	}
}

func (pc *puzzleController) Put(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var pieces []*model.Puzzlepiece
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&pieces)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode puzzle put body")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	v := mux.Vars(request)
	p, err := pc.puzzleRepository.Put(v["surveyId"], v["questionId"], pieces)
	if err != nil {
		log.Error().Err(err).Msg("unable to update pieces")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(p)
	return
}

func (pc *puzzleController) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)
	pieces, err := pc.puzzleRepository.GetAll(v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve questions.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(PuzzleResponse{Pieces: pieces})
	return
}

func (pc *puzzleController) GetAllForQuestionaire(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	lookupUser := model.User{Email: v["email"]}
	retrievedUser, err := pc.userRepository.Get(&lookupUser)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve user.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	q, err := pc.questionRepository.Get(v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve question.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ans, err := pc.answeredRepository.GetSingle(retrievedUser.ID, q)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve question state.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(ans) > 0 {
		log.Error().Err(err).Msg("Question has already been watched.")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	pieces, err := pc.puzzleRepository.GetAll(v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve questions.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(PuzzleResponse{Pieces: pieces})
	return
}
