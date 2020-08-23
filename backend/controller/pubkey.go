package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type pubkeyController struct{}

type PubkeyResponse struct {
	Pubkey string `json:"pubkey"`
}

func NewPubkeyController() *pubkeyController {
	return &pubkeyController{}
}

func (p *pubkeyController) Get(writer http.ResponseWriter, request *http.Request) {
	publicKey, err := ioutil.ReadFile("common/keys/jwtRS256.key.pub")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	response := PubkeyResponse{Pubkey: string(publicKey)}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
