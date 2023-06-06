package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/dto"
)

var privateKey = []byte("THIS_IS_NOT_SO_SECRET+YOU_SHOULD_DEFINITELY_CHANGE_IT")

func GenerateJWT(accDto *dto.AccountDTO) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  accDto.ID,
		"iat": time.Now().Unix(),
	})
	return token.SignedString(privateKey)
}
