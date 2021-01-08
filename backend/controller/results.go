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
	answeredRepository       repository.AnsweredRepository
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
	answeredRepository repository.AnsweredRepository,
) ResultsController {
	return &resultsController{
		questionRepository:       questionRepository,
		multiplechoiceRepository: multiplechoiceRepository,
		puzzleSolutionRepository: puzzleSolutionRepository,
		puzzleAnswerRepository:   puzzleAnswerRepository,
		userRepository:           userRepository,
		answeredRepository:       answeredRepository,
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

				answered, err := rc.answeredRepository.GetSingle(user, q)
				if err != nil {
					resultAnswers = append(resultAnswers, &Answer{User: user.Email, Score: ""})
					continue
				}
				if len(answered) < 1 {
					resultAnswers = append(resultAnswers, &Answer{User: user.Email, Score: ""})
					continue
				}
				if answered[0].Answered == false {
					resultAnswers = append(resultAnswers, &Answer{User: user.Email, Score: ""})
					continue
				}
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
	answersByUser := make(map[uint]map[uint]string)
	columnLabels := make([]string, 0, len(questions)+1)
	columnLabels = append(columnLabels, "user")

	users, err := rc.userRepository.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("unable to retrieve users")
		return ""
	}

	uids := rc.mapEmailsToIds(users)
	qids := make([]uint, 0, len(questions))

	// iterating over questions
	// adding question id to column labels
	// iterating over answers per question,
	// populating a users-to-(questionid-to-answer map) map
	for _, q := range questions {
		columnLabels = append(columnLabels, fmt.Sprint(q.Id))
		qids = append(qids, q.Id)
		for _, a := range q.Answers {
			id, ok := uids[a.User]
			if !ok {
				log.Error().Err(err).Str("email", a.User).Msg("cant find id for users")
				continue
			}
			if answersByUser[id] == nil {
				answersByUser[id] = make(map[uint]string)
			}
			answersByUser[id][q.Id] = a.Score
		}
	}

	joinedColumnLabels := strings.Join(columnLabels, ",")

	for userId, answersMap := range answersByUser {

		newRow := make([]string, 0, len(questions)+1)
		newRow = append(newRow, fmt.Sprint(userId))
		for _, id := range qids {
			a, ok := answersMap[id]
			if ok {
				newRow = append(newRow, a)
				continue
			}
			newRow = append(newRow, "")
		}

		joinedColumnLabels += "\n"
		joinedColumnLabels += strings.Join(newRow, ",")
	}

	return joinedColumnLabels
}

func (rc *resultsController) mapEmailsToIds(users []*model.User) map[string]uint {
	result := make(map[string]uint)
	for _, u := range users {
		result[u.Email] = u.ID
	}
	return result
}
