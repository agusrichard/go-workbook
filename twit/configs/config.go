package configs

import (
	"fmt"
	"os"
	"twit/models"

	"github.com/joho/godotenv"
)

func GetConfig() *models.Config {
	_, err := os.Stat(".env")

	if !os.IsNotExist(err) {
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Println("Error while reading the env file", err)
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
		TimeZone: "Asia/Jakarta",
	}

	return config
}
