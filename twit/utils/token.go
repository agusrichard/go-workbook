package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"twit/configs"
	"twit/models"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte(configs.GetConfig().SecretKey)

func GenerateToken(user models.User) (string, *models.RequestError) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		err := &models.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("INTERNAL SERVER ERROR"),
		}
		Logging(err)
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (models.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		log.Println("Error to parse token", err)
		return models.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idStr := fmt.Sprintf("%v", claims["id"])
		id, _ := strconv.ParseInt(idStr, 10, 64)
		email := claims["email"].(string)
		return models.User{Email: email, ID: uint(id)}, nil
	}

	return models.User{}, err
}
