package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type PuzzleAnswerController interface {
	Post(writer http.ResponseWriter, request *http.Request)
}

type puzzleAnswerController struct {
	puzzleAnswerRepository repository.PuzzleAnswerRepository
}

func NewPuzzleAnswerController(pr repository.PuzzleAnswerRepository) PuzzleAnswerController {
	return &puzzleAnswerController{puzzleAnswerRepository: pr}
}

func (pc *puzzleAnswerController) Post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var pieces []*model.PuzzleAnswer
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&pieces)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode puzzle answer post body")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	pieces2, err := pc.puzzleAnswerRepository.Post(pieces)
	if err != nil {
		log.Error().Err(err).Msg("unable to write post choice to db")
		writer.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(writer).Encode(pieces2)
	return
}
