package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/usecase/repository"
)

type resultsController struct {
	questionRepository       repository.QuestionRepository
	multiplechoiceRepository repository.MultipleChoiceRepository
}

type ResultsController interface {
	Get(writer http.ResponseWriter, request *http.Request)
}

type Results struct {
	Questions []*Question `json:"questions"`
}

type Question struct {
	Text    string    `json:"text"`
	Type    string    `json:"type"`
	Answers []*Answer `json:"answers"`
}

type Answer struct {
	User  string `json:"user"`
	Score string `json:"score"`
}

func NewResultsController(
	questionRepository repository.QuestionRepository,
	multiplechoiceRepository repository.MultipleChoiceRepository,
) ResultsController {
	return &resultsController{
		questionRepository:       questionRepository,
		multiplechoiceRepository: multiplechoiceRepository,
	}
}

func (rc *resultsController) Get(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	questions, err := rc.questionRepository.GetAll(v["surveyId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve questions.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	resultQuestions := make([]*Question, 0, len(questions))

	for _, q := range questions {
		if q.Type == "multiplechoice" {
			answers, err := rc.multiplechoiceRepository.GetAll(q.ID)
			if err != nil {
				log.Error().Err(err).Msg("Unable to retrieve answers.")
				writer.WriteHeader(http.StatusInternalServerError)
			}
			resultAnswers := make([]*Answer, 0, len(answers))

			for _, a := range answers {
				resultAnswers = append(resultAnswers, &Answer{User: a.Email, Score: a.Text})
			}
			resultQuestions = append(resultQuestions, &Question{Type: q.Type, Text: q.Text, Answers: resultAnswers})
		}
	}

	json.NewEncoder(writer).Encode(&Results{Questions: resultQuestions})
}
