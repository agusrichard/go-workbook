package utils

import (
	"errors"
	"net/http"
	"twit/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *models.RequestError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", &models.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("INTERNAL SERVER ERROR"),
		}
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
