package config

import (
	"db-experiment/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DB *sqlx.DB

// ConnectDB to get all needed db connections for application
func ConnectDB(config *model.Config) *sqlx.DB {
	DB = getDBConnection(config)

	return DB
}

func getDBConnection(config *model.Config) *sqlx.DB {
	var dbConnectionStr string

	dbConnectionStr = fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName,
		config.Database.Username,
		config.Database.Password,
	)

	db, err := sqlx.Open("postgres", dbConnectionStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//TODO: experiment with correct values
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(5)

	log.Println("Connected to DB")
	return db
}