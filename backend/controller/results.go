package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type resultsController struct {
	questionRepository       repository.QuestionRepository
	multiplechoiceRepository repository.MultipleChoiceRepository
	puzzleSolutionRepository repository.PuzzleRepository
	puzzleAnswerRepository   repository.PuzzleAnswerRepository
	userRepository           repository.UserRepository
}

type ResultsController interface {
	Get(writer http.ResponseWriter, request *http.Request)
	GetSingle(writer http.ResponseWriter, request *http.Request)
}

type Results struct {
	Questions []*Question `json:"questions"`
}

type Question struct {
	Text    string    `json:"text"`
	Type    string    `json:"type"`
	Id      uint      `json:"id"`
	Answers []*Answer `json:"answers"`
}

type Answer struct {
	User  string `json:"user"`
	Score string `json:"score"`
}

type ResultResponse struct {
	Result string `json:"result"`
}

func NewResultsController(
	questionRepository repository.QuestionRepository,
	multiplechoiceRepository repository.MultipleChoiceRepository,
	puzzleSolutionRepository repository.PuzzleRepository,
	puzzleAnswerRepository repository.PuzzleAnswerRepository,
	userRepository repository.UserRepository,
) ResultsController {
	return &resultsController{
		questionRepository:       questionRepository,
		multiplechoiceRepository: multiplechoiceRepository,
		puzzleSolutionRepository: puzzleSolutionRepository,
		puzzleAnswerRepository:   puzzleAnswerRepository,
		userRepository:           userRepository,
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
		switch q.Type {
		case "multiplechoice":
			answers, err := rc.multiplechoiceRepository.GetAll(q.ID)
			if err != nil {
				log.Error().Err(err).Msg("Unable to retrieve answers.")
				writer.WriteHeader(http.StatusInternalServerError)
			}
			resultAnswers := make([]*Answer, 0, len(answers))

			for _, a := range answers {
				resultAnswers = append(resultAnswers, &Answer{User: a.Email, Score: a.Text})
			}
			resultQuestions = append(resultQuestions, &Question{Type: q.Type, Text: q.Text, Id: q.ID, Answers: resultAnswers})
		case "puzzle":
			solutions, err := rc.puzzleSolutionRepository.GetAll(strconv.Itoa(int(q.ID)))
			if err != nil {
				log.Error().Err(err).Msg("Unable to retrieve puzzle solution.")
				writer.WriteHeader(http.StatusInternalServerError)
			}

			solutionsMap, err := rc.mapPositions(solutions)
			if err != nil {
				log.Error().Err(err).Msg("Unable to map their positions onto solutions.")
				writer.WriteHeader(http.StatusInternalServerError)
			}

			users, err := rc.userRepository.GetAll()
			if err != nil {
				log.Error().Err(err).Msg("Unable to retrieve the users.")
				writer.WriteHeader(http.StatusInternalServerError)
			}

			resultAnswers := make([]*Answer, 0)
			for _, user := range users {
				userPieces, err := rc.puzzleAnswerRepository.GetUserSolution(user.Email, strconv.Itoa(int(q.ID)))
				if err != nil {
					log.Error().Err(err).Msg("Unable to retrieve the submitted answers for this puzzle question.")
					writer.WriteHeader(http.StatusInternalServerError)
				}

				userPiecesMap, err := rc.mapUserSolution(userPieces)
				if err != nil {
					log.Error().Err(err).Msg("Unable to map their positions onto solutions.")
					writer.WriteHeader(http.StatusInternalServerError)
				}

				score := 0
				for i := 0; i < 24; i++ {
					solution, fieldFilled := solutionsMap[strconv.Itoa(i)]
					answer, fieldAnswered := userPiecesMap[strconv.Itoa(i)]
					if !fieldFilled && !fieldAnswered {
						score++
						continue
					}

					if fieldFilled && fieldAnswered {
						if solution.Image == answer.Image {
							score++
							continue
						}
					}
				}

				resultAnswers = append(resultAnswers, &Answer{User: user.Email, Score: strconv.Itoa(score)})
			}

			resultQuestions = append(resultQuestions, &Question{Type: q.Type, Text: q.Text, Id: q.ID, Answers: resultAnswers})
		default:
			continue
		}
	}

	transformedResult := rc.transformOutput(resultQuestions)

	json.NewEncoder(writer).Encode(&ResultResponse{Result: transformedResult})
}

func (rc *resultsController) GetSingle(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	question, err := rc.questionRepository.Get(v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable retrieve question.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if question.Type != "puzzle" {
		log.Error().Err(err).Msg("Cant retrieve score for non puzzle question.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	solutions, err := rc.puzzleSolutionRepository.GetAll(v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve puzzle solution.")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	solutionsMap, err := rc.mapPositions(solutions)
	if err != nil {
		log.Error().Err(err).Msg("Unable to map their positions onto solutions.")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	userPieces, err := rc.puzzleAnswerRepository.GetUserSolution(v["email"], v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve the submitted answers for this puzzle question.")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	userPiecesMap, err := rc.mapUserSolution(userPieces)
	if err != nil {
		log.Error().Err(err).Msg("Unable to map their positions onto solutions.")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	score := 0
	for i := 0; i < 24; i++ {
		solution, fieldFilled := solutionsMap[strconv.Itoa(i)]
		answer, fieldAnswered := userPiecesMap[strconv.Itoa(i)]
		if !fieldFilled && !fieldAnswered {
			score++
			continue
		}

		if fieldFilled && fieldAnswered {
			if solution.Image == answer.Image {
				score++
				continue
			}
		}
	}

	json.NewEncoder(writer).Encode(score)
}

func (rc *resultsController) mapPositions(solutions []*model.Puzzlepiece) (map[string]*model.Puzzlepiece, error) {
	result := make(map[string]*model.Puzzlepiece, len(solutions))
	for _, s := range solutions {
		result[s.Position] = s
	}
	return result, nil
}

func (rc *resultsController) mapUserSolution(solutions []*model.PuzzleAnswer) (map[string]*model.PuzzleAnswer, error) {
	result := make(map[string]*model.PuzzleAnswer, len(solutions))
	for _, s := range solutions {
		result[s.Position] = s
	}
	return result, nil
}

func (rc *resultsController) transformOutput(questions []*Question) string {
	// maps an users email address
	// to a map of question ids
	// to the users answers
	answersByUser := make(map[string]map[uint]string)
	columnLabels := make([]string, 0, len(questions)+1)
	columnLabels = append(columnLabels, "user")

	// iterating over questions
	// adding question id to column labels
	// iterating over answers per question,
	// populating a users-to-(questionid-to-answer map) map
	for _, q := range questions {
		columnLabels = append(columnLabels, fmt.Sprint(q.Id))
		for _, a := range q.Answers {
			if answersByUser[a.User] == nil {
				answersByUser[a.User] = make(map[uint]string)
			}
			answersByUser[a.User][q.Id] = a.Score
		}
	}

	// adding each questionid-to-answer map to a slice to get rid of the user key
	answerSlice := make([]map[uint]string, 0, len(answersByUser))
	for _, v := range answersByUser {
		answerSlice = append(answerSlice, v)
	}

	// creating the first row of column labels
	joinedColumnLabels := strings.Join(columnLabels, ",")

	// iterating over questionid-to-answer map slice
	// i indicates the new user id
	for i, v := range answerSlice {
		// init new "userid"-to-answer map
		newRow := make([]string, 0, len(questions)+1)

		// appending "userid" as first question
		newRow = append(newRow, fmt.Sprint(i))

		// iterate over questions to fetch answer
		for _, q := range questions {
			nextItem := ""
			if j, ok := v[q.Id]; ok {
				nextItem = j
			}
			newRow = append(newRow, nextItem)
		}
		joinedColumnLabels += "\n"
		joinedColumnLabels += strings.Join(newRow, ",")
	}

	return joinedColumnLabels
}
