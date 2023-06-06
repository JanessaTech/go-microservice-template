package helper

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/dto"
	"golang.org/x/crypto/bcrypt"
)

var privateKey = []byte("THIS_IS_NOT_SO_SECRET+YOU_SHOULD_DEFINITELY_CHANGE_IT")

func GenerateJWT(accDto *dto.AccountDTO) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  accDto.ID,
		"iat": time.Now().Unix(),
	})
	return token.SignedString(privateKey)
}

func EncodePassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

func ValidateJWT(ctx *gin.Context) error {
	token, err := getToken(ctx)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func GetCurrentAccountId(ctx *gin.Context) (int, error) {
	token, err := getToken(ctx)
	if err != nil {
		return -1, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	userId, _ := claims["id"].(int)
	return userId, nil
}
