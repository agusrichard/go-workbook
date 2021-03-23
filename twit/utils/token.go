package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"twit/configs"
	"twit/models"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var SecretKey = []byte(configs.GetConfig().SecretKey)

func GenerateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
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
		email := claims["email"].(string)
		return models.User{Email: email, Model: gorm.Model{ID: uint(id)}}, nil
	}

	return models.User{}, err
}
