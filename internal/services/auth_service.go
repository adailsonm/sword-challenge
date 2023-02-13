package services

import (
	"errors"
	"os"
	"time"

	"github.com/adailsonm/desafio-sword/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
}

func NewAuthService() AuthService {
	return AuthService{}
}

func (s AuthService) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("token expired")
		}
	}
	return false, errors.New("couldn't handle token")
}

func (s AuthService) CreateToken(user *models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{}
	claims["userId"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(secret))
}
