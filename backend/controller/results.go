package controller

import (
	"encoding/json"
	"errors"
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
	GetUserAverage(writer http.ResponseWriter, request *http.Request)
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
	Order string `json:"order"`
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

	users, err := rc.userRepository.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve the users.")
		writer.WriteHeader(http.StatusInternalServerError)
	}

	usersByEmail := make(map[string]*model.User, len(users))
	for _, user := range users {
		usersByEmail[user.Email] = user
	}

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

				if _, ok := usersByEmail[a.Email]; !ok {
					log.Error().Err(err).Msg("Unable to find email in retrieved users.")
					continue
				}
				id := usersByEmail[a.Email].ID

				answered, err := rc.answeredRepository.GetSingle(id, q)
				if err != nil {
					resultAnswers = append(resultAnswers, &Answer{User: a.Email, Score: ""})
					continue
				}
				if len(answered) < 1 {
					resultAnswers = append(resultAnswers, &Answer{User: a.Email, Score: ""})
					continue
				}
				if answered[0].Answered == false {
					resultAnswers = append(resultAnswers, &Answer{User: a.Email, Score: ""})
					continue
				}

				resultAnswers = append(resultAnswers, &Answer{User: a.Email, Score: a.Text, Order: strconv.Itoa(int(answered[0].Order))})
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

			resultAnswers := make([]*Answer, 0)
			for _, user := range users {

				answered, err := rc.answeredRepository.GetSingle(user.ID, q)
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

				resultAnswers = append(resultAnswers, &Answer{User: user.Email, Score: strconv.Itoa(score), Order: strconv.Itoa(int(answered[0].Order))})
			}

			resultQuestions = append(resultQuestions, &Question{Type: q.Type, Text: q.Text, Id: q.ID, Answers: resultAnswers})
		default:
			continue
		}
	}

	transformedResult := rc.transformOutput(resultQuestions, users)

	json.NewEncoder(writer).Encode(&ResultResponse{Result: transformedResult})
}

func (rc *resultsController) GetSingle(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	score, err := rc.getSingleScore(v["email"], v["questionId"])
	if err != nil {
		log.Error().Err(err).Msg("Unable compute single question score.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(score)
}

func (rc *resultsController) getSingleScore(email string, questionId string) (int, error) {
	question, err := rc.questionRepository.Get(questionId)
	if err != nil {
		return -1, errors.New("Unable retrieve question.")
	}

	if question.Type != "puzzle" {
		return -1, errors.New("Cant retrieve score for non puzzle question.")
	}

	solutions, err := rc.puzzleSolutionRepository.GetAll(questionId)
	if err != nil {
		return -1, errors.New("Unable to retrieve puzzle solution.")
	}

	solutionsMap, err := rc.mapPositions(solutions)
	if err != nil {
		return -1, errors.New("Unable to map their positions onto solutions.")
	}

	userPieces, err := rc.puzzleAnswerRepository.GetUserSolution(email, questionId)
	if err != nil {
		return -1, errors.New("Unable to retrieve the submitted answers for this puzzle question.")
	}

	userPiecesMap, err := rc.mapUserSolution(userPieces)
	if err != nil {
		return -1, errors.New("Unable to map their positions onto solutions.")
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
	return score, nil
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

func (rc *resultsController) transformOutput(questions []*Question, users []*model.User) string {
	// maps an users email address
	// to a map of question ids
	// to the users answers
	answersByUser := make(map[uint]map[uint]string)
	columnLabels := make([]string, 0, len(questions)+1)
	columnLabels = append(columnLabels, "user")

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
				log.Error().Str("email", a.User).Msg("cant find id for users")
				continue
			}
			if answersByUser[id] == nil {
				answersByUser[id] = make(map[uint]string)
			}
			var score string = ""
			if a.Order != "" && a.Score != "" {
				score = fmt.Sprintf("%v$%v", a.Order, a.Score)
			}
			answersByUser[id][q.Id] = score
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

func (rc *resultsController) GetUserAverage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	v := mux.Vars(request)

	surveyId := v["surveyId"]
	email := v["email"]

	lookupUser := model.User{Email: email}
	retrievedUser, err := rc.userRepository.Get(&lookupUser)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve user.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	answered, err := rc.answeredRepository.Get(retrievedUser, surveyId)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve answered questions for user.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	score := 0.0
	puzzleCount := 0.0
	for _, a := range answered {
		singleScore, err := rc.getSingleScore(email, strconv.Itoa(int(a.QuestionId)))
		if err != nil {
			log.Error().Err(err).Str("questionId", strconv.Itoa(int(a.QuestionId))).Str("userId", strconv.Itoa(int(retrievedUser.ID))).Msg("Unable to evaluate score for user + question pair.")
			continue
		}
		score += float64(singleScore)
		puzzleCount++
	}

	json.NewEncoder(writer).Encode(fmt.Sprintf("%.2f", score/puzzleCount))
}
