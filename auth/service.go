package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToke(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("PROJECT_s3kr3T_k3y")

func NewServices() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	sign, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return sign, err
	}

	return sign, nil
}

func (s *jwtService) ValidateToke(encoded string) (*jwt.Token, error) {
	token, err := jwt.Parse(encoded, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
