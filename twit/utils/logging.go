package utils

import (
	"log"
	"twit/models"
)

func Logging(err *models.RequestError) {
	if err != nil {
		log.Println(err.Error())
	}
}
