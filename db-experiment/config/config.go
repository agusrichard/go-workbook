package config

import (
	"db-experiment/models"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func GetConfig() *model.Config {
	_, err := os.Stat(".env")

	if !os.IsNotExist(err) {
		err := godotenv.Load(".env")

		if err != nil {
			log.Println("Error while reading the env file", err)
			panic(err)
		}
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	config := &model.Config{
		Database: model.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			DbName:   os.Getenv("DB"),
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
		},
		TimeZone:  "Asia/Jakarta",
		SecretKey: os.Getenv("SECRET_KEY"),
	}

	return config
}
