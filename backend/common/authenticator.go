package common

import (
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/model/response"

	"time"

	"github.com/dgrijalva/jwt-go"
)

type authenticator struct {
	appKey string
}

// authenticator interface
type Authenticator interface {
	CreateToken(*model.User, *model.Role) (*response.UserResponse, error)
}

// blub
func NewAuthenticator(appKey string) Authenticator {
	return &authenticator{appKey: appKey}
}

func (a authenticator) CreateToken(u *model.User, r *model.Role) (*response.UserResponse, error) {
	var err error
	expiration := time.Now().Add(time.Minute * 15).Unix()
	claims := &model.CustomClaims{
		Name:       u.Name,
		Authorized: true,
		Email:      u.Email,
		Role:       r.Role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expiration,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.appKey))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return nil, err
	}
	return &response.UserResponse{Token: tokenString, ExpiresAt: expiration, Username: u.Name, Role: r.Role}, nil
}
