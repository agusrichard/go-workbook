package utils

import (
	"fmt"
	"grpc-auth/models"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	SecretKey = []byte("ThisIsASecretKey")
)

func GenerateToken(tokenResult models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = tokenResult.ID
	claims["username"] = tokenResult.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Println("Error in generating key")
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
		username := claims["username"].(string)
		return models.User{Username: username, ID: id}, nil
	}

	return models.User{}, err
}
