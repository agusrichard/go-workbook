package configs

import (
	"log"
	"os"
	"twit/models"

	"github.com/joho/godotenv"
)

func GetConfig() *models.Config {
	_, err := os.Stat(".env")

	if !os.IsNotExist(err) {
		err := godotenv.Load(".env")

		if err != nil {
			log.Println("Error while reading the env file", err)
			panic(err)
		}
	}

	config := &models.Config{
		Database: models.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			DbName:   os.Getenv("DB"),
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
		},
		TimeZone:  "Asia/Jakarta",
		SecretKey: os.Getenv("SECRET_KEY"),
	}

	return config
}
