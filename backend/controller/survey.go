package controller

import (
	"encoding/json"
	"net/http"
)

type surveyController struct {
}

type Resp struct {
	Message string `json:"message"`
}

// UserController defines functions available for requesat handling
type SurveyController interface {
	Test(writer http.ResponseWriter, request *http.Request)
}

// NewUserController provides functions for request handling
func NewSurveyController() SurveyController {
	return &surveyController{}
}

func (uc *surveyController) Test(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	xyz := &Resp{Message: "hey everybody, shit is authed!!"}
	json.NewEncoder(writer).Encode(xyz)
	return
}
