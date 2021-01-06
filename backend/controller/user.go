package controller

import (
	"encoding/json"
	"net/http"

	"backend/model"
	"backend/model/response"
	"backend/usecase/interactor"

	"github.com/rs/zerolog/log"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

// UserController defines functions available for requesat handling
type UserController interface {
	Login(writer http.ResponseWriter, request *http.Request)
	Signup(writer http.ResponseWriter, request *http.Request)
	RefreshToken(writer http.ResponseWriter, request *http.Request)
	Get(writer http.ResponseWriter, request *http.Request)
}

// NewUserController provides functions for request handling
func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) Login(writer http.ResponseWriter, request *http.Request) {

	// setting up a json decoder that returns an error if it encounters any fields not present in Message-struct
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	var user model.User
	decoderErr := decoder.Decode(&user)
	if decoderErr != nil {
		handleDecoderError(decoderErr, writer)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	if user.Password == "" || user.Email == "" {
		log.Error().Str("username", user.Email).Str("Password", user.Password).Msg("Loginattempt with empty username or Password.")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	result, err := uc.userInteractor.Get(&user)
	if err != nil {
		return
	}

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(result)
	return
}

func (uc *userController) Signup(writer http.ResponseWriter, request *http.Request) {

	var defaultRole string = "user"

	// setting up a json decoder that returns an error if it encounters any fields not present in Message-struct
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	var user model.User
	decoderErr := decoder.Decode(&user)
	if decoderErr != nil {
		handleDecoderError(decoderErr, writer)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	if user.Password == "" || user.Email == "" {
		log.Error().Msg("Loginattempt with empty username or Password.")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	result, err := uc.userInteractor.Post(&user, defaultRole)
	if err != nil {
		return
	}

	json.NewEncoder(writer).Encode(result)
	return
}

func (uc *userController) RefreshToken(writer http.ResponseWriter, request *http.Request) {
	result, err := uc.userInteractor.Refresh(request)
	if err != nil {
		log.Error().Err(err).Msg("Unable to update token.")
		writer.WriteHeader(http.StatusForbidden)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(result)
	return
}

func handleDecoderError(err error, writer http.ResponseWriter) {
	log.Err(err).Caller().Msg("Error while decoding the passed User.")
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (uc *userController) Get(writer http.ResponseWriter, request *http.Request) {
	var users []*response.SmolUserResponse
	users, err := uc.userInteractor.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve smol users.")
		writer.WriteHeader(http.StatusForbidden)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}
