package common

import (
	"backend/model"
	"backend/model/response"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type auth struct {
	AppKey string
}

// auth interface
type Auth interface {
	CreateToken(*model.User, *model.Role) (*response.Token, error)
}

// blub
func NewAuth(appKey string) Auth {
	return auth{appKey}
}

func (a auth) CreateToken(u *model.User, r *model.Role) (*response.Token, error) {
	var err error
	//Creating Access Token
	// os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = u.Email
	atClaims["password"] = u.Password
	atClaims["role"] = r.Role
	exp := time.Now().Add(time.Minute * 15).Unix()
	atClaims["expiresAt"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv(a.AppKey)))
	if err != nil {
		return nil, err
	}
	return &response.Token{Token: token, ExpiresAt: exp}, nil
}
