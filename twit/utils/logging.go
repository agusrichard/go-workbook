package utils

import (
	"fmt"
	"log"
)

func Logging(err error, location, errorMessage string) {
	if err != nil {
		log.Println(fmt.Sprintf("Location: %s; Error Message: %s\n", location, errorMessage))
	}
}
