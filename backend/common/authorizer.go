package common

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"backend/model"

	"github.com/dgrijalva/jwt-go"
)

type authorizer struct {
	appKey string
}

// authorizer interface
type Authorizer interface {
	IsUser(http.HandlerFunc) http.HandlerFunc
	IsAdmin(http.HandlerFunc) http.HandlerFunc
	IsOwner(http.HandlerFunc) http.HandlerFunc
	UnwrapClaims(r *http.Request) (*model.CustomClaims, error)
}

// blub
func NewAuthorizer(appKey string) Authorizer {
	return &authorizer{appKey: appKey}
}

func (a *authorizer) IsUser(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		claims, err := a.UnwrapClaims(request)
		if err != nil {
			log.Error().Err(err).Msg("Failed to decipher token.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}
		expired := a.isExpired(claims)
		if !expired {
			log.Error().Err(err).Msg("Token expired.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		perms := a.hasPermissions("user", claims)
		if !perms {
			log.Error().Err(err).Msg("Wrong role.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		log.Info().Msg("Authorized")
		next.ServeHTTP(writer, request)
	}
}

func (a *authorizer) IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		claims, err := a.UnwrapClaims(request)
		if err != nil {
			log.Error().Err(err).Msg("Failed to decipher token.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}
		expired := a.isExpired(claims)
		if !expired {
			log.Error().Err(err).Msg("Token expired.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		perms := a.hasPermissions("admin", claims)
		if !perms {
			log.Error().Err(err).Msg("Wrong role.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		log.Info().Msg("Authorized")
		next.ServeHTTP(writer, request)
	}
}

func (a *authorizer) IsOwner(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		claims, err := a.UnwrapClaims(request)
		if err != nil {
			log.Error().Err(err).Msg("Failed to decipher token.")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		v := mux.Vars(request)

		if v["email"] != claims.Email {
			log.Error().Err(err).Msg("User does not own the requested resource")
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		log.Info().Msg("User owns the requested resource")
		next.ServeHTTP(writer, request)
	}
}

func (a *authorizer) isExpired(claims *model.CustomClaims) bool {

	if claims.ExpiresAt < time.Now().Unix() {
		return false
	}
	return true
}

func (a *authorizer) hasPermissions(role string, claims *model.CustomClaims) bool {

	if claims.Role != role && !(claims.Role == "admin") {
		return false
	}
	return true
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
