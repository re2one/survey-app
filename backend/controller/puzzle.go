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
	puzzleRepository repository.PuzzleRepository
}

type PuzzleController interface {
	Put(writer http.ResponseWriter, request *http.Request)
}

func NewPuzzleController(repo repository.PuzzleRepository) PuzzleController {
	return &puzzleController{puzzleRepository: repo}
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
	}

	v := mux.Vars(request)
	p, err := pc.puzzleRepository.Put(v["surveyId"], v["questionId"], pieces)
	if err != nil {
		log.Error().Err(err).Msg("unable to update pieces")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(writer).Encode(p)
	return
}
