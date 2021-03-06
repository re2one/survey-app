package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

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
	Viewed(http.ResponseWriter, *http.Request)
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

	state, err := uc.answeredRepository.Get(retrievedUser, surveyId)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve the answered questions.")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	finished := false
	for answered := true; answered; {

		if _, ok := state[currentQuestion.ID]; !ok {
			answered = false
		}

		fullQuestion := response.FullQuestion{
			QuestionId: currentQuestion.ID,
			Title:      currentQuestion.Title,
			Type:       currentQuestion.Type,
			Answered:   answered,
			Example:    currentQuestion.Example,
		}

		fullQuestions = append(fullQuestions, &fullQuestion)

		if _, ok := state[currentQuestion.ID]; !ok {
			continue
		}

		var nextQuestion string
		switch currentQuestion.Type {
		case "multiplechoice":
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
			switch usedChoice.TypeOfNextQuestion {
			case "random":
				questionsInBracket, err := uc.questionRepository.GetBracket(surveyId, usedChoice.NextQuestion)
				if err != nil {
					log.Error().Err(err).Msg("Unable to fetch questions in random bracket.")
					writer.WriteHeader(http.StatusInternalServerError)
					return
				}

				possibleNextQuestions := make([]uint, 0)
				var alreadyPreviewed uint = 0

				for k, v := range questionsInBracket {
					if _, ok := state[k]; !ok {
						previewed, err := uc.answeredRepository.GetSingle(retrievedUser.ID, v)
						if err == nil && len(previewed) > 0 {
							if previewed[0].Viewed {
								alreadyPreviewed = v.ID
							}
						}
						possibleNextQuestions = append(possibleNextQuestions, k)
						continue
					}
					fq := response.FullQuestion{
						QuestionId: v.ID,
						Title:      v.Title,
						Type:       v.Type,
						Answered:   answered,
						Example:    v.Example,
					}
					fullQuestions = append(fullQuestions, &fq)
					delete(questionsInBracket, k)
				}

				if alreadyPreviewed != 0 {
					nextQuestion = fmt.Sprint(questionsInBracket[alreadyPreviewed].ID)
					break
				}

				if len(possibleNextQuestions) < 1 {
					nextQuestion = fmt.Sprint(usedChoice.SecondToNext)
					break
				}

				if len(possibleNextQuestions) > 0 {
					randomIndex := rand.Intn(len(possibleNextQuestions))
					nextQuestion = fmt.Sprint(questionsInBracket[possibleNextQuestions[uint(randomIndex)]].ID)
				}

			default:
				nextQuestion = fmt.Sprint(usedChoice.NextQuestion)
			}
		case "puzzle":
			q, err := uc.questionRepository.Get(strconv.Itoa(int(currentQuestion.ID)))
			if err != nil {
				log.Error().Err(err).Msg("Unable to fetch question.")
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			switch q.TypeOfNextQuestion {
			case "random":
				questionsInBracket, err := uc.questionRepository.GetBracket(surveyId, q.Next)
				if err != nil {
					log.Error().Err(err).Msg("Unable to fetch questions in random bracket.")
					writer.WriteHeader(http.StatusInternalServerError)
					return
				}

				possibleNextQuestions := make([]uint, 0)
				var alreadyPreviewed uint = 0

				for k, v := range questionsInBracket {
					if _, ok := state[k]; !ok {
						previewed, err := uc.answeredRepository.GetSingle(retrievedUser.ID, v)
						if err == nil && len(previewed) > 0 {
							if previewed[0].Viewed {
								alreadyPreviewed = v.ID
							}
						}
						possibleNextQuestions = append(possibleNextQuestions, k)
						continue
					}
					fq := response.FullQuestion{
						QuestionId: v.ID,
						Title:      v.Title,
						Type:       v.Type,
						Answered:   answered,
						Example:    v.Example,
					}
					fullQuestions = append(fullQuestions, &fq)
					delete(questionsInBracket, k)
				}

				if alreadyPreviewed != 0 {
					nextQuestion = fmt.Sprint(questionsInBracket[alreadyPreviewed].ID)
					break
				}

				if len(possibleNextQuestions) < 1 {
					nextQuestion = fmt.Sprint(q.SecondToNext)
					break
				}

				if len(possibleNextQuestions) > 0 {
					randomIndex := rand.Intn(len(possibleNextQuestions))
					nextQuestion = fmt.Sprint(questionsInBracket[possibleNextQuestions[uint(randomIndex)]].ID)
				}

			default:
				nextQuestion = fmt.Sprint(q.Next)
			}
		default:
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		if nextQuestion == "lastQuestion" {
			log.Info().Msg("This was the last question.")
			finished = true
			answered = false
			continue
		}

		currentQuestion, err = uc.questionRepository.Get(nextQuestion)
		if err != nil {
			log.Error().Err(err).Msg("Unable to retrieve the following question.")
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
	order := v["order"]
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

	answered, err := uc.answeredRepository.Post(retrievedUser, &question, order)
	if err != nil {
		log.Error().Err(err).Msg("unable post question state")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(answered)
	return
}

func (uc *fullQuestionsController) Viewed(writer http.ResponseWriter, request *http.Request) {
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

	answered, err := uc.answeredRepository.Viewed(retrievedUser, &question)
	if err != nil {
		log.Error().Err(err).Msg("unable post question state")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(answered)
	return
}
