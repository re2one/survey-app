package model

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	Name       string `json:"name"`
	Authorized bool   `json:"authorized"`
	Email      string `json:"email"`
	// Password string`json:"password"`
	Role string `json:"role"`

	jwt.StandardClaims
}
