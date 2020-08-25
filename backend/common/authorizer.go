package common

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"backend/model"

	"github.com/dgrijalva/jwt-go"
)

type authorizer struct {
	appKey string
}

// authorizer interface
type Authorizer interface {
	IsAuthorized(string, http.HandlerFunc) http.HandlerFunc
	UnwrapClaims(r *http.Request) (*model.CustomClaims, error)
}

// blub
func NewAuthorizer(appKey string) Authorizer {
	return &authorizer{appKey: appKey}
}

func (a *authorizer) IsAuthorized(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		claims, err := a.UnwrapClaims(request)
		if err != nil {
			log.Error().Err(err).Msg("Failed to decipher token.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}
		if claims.ExpiresAt < time.Now().Unix() {
			log.Error().Err(err).Msg("Token expired.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}
		if claims.Role != role && !(claims.Role == "admin") {
			log.Error().Err(err).Msg("Wrong role.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		log.Info().Msg("Authorized")
		next(writer, request)
	}
}

func (a *authorizer) UnwrapClaims(request *http.Request) (*model.CustomClaims, error) {
	bearerToken := request.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) != 2 {
		log.Error().Msg("Incorrect token.")
		return nil, errors.New("incorrect token")
	}

	claims := model.CustomClaims{}
	_, err := jwt.ParseWithClaims(strArr[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.appKey), nil
	})
	if err != nil {
		log.Error().Msg("failed to unwrap token.")
		return nil, errors.New("failed to unwrap token")
	}
	return &claims, nil
}
