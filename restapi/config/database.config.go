package config

import (
	"database/sql"
	"fmt"

	// Postgres
	_ "github.com/lib/pq"
)

// DB to use database
var DB *sql.DB

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// InitDb Initializing database connection
func InitDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[host], config[port],
		config[user], config[password], config[dbname])

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	conf[host] = host
	conf[port] = port
	conf[user] = user
	conf[password] = password
	conf[dbname] = dbname
	return conf
}
