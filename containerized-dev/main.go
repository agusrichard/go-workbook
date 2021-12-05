package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func getDBConnection() *sqlx.DB {
	// Note that the port that was exposed to local is 3000, but in here specified 5432
	// because in here we're connecting container to container
	dbConnectionStr := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		"localhost",
		5432,
		"containerized_dev",
		"postgres",
		"postgres",
	)

	fmt.Println("dbConnectionst", dbConnectionStr)

	db, err := sqlx.Open("postgres", dbConnectionStr)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//TODO: experiment with correct values
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(5)

	fmt.Println("Connected to DB")
	return db
}

func newHandler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		message := r.URL.Query().Get("x")

		query := `
			INSERT INTO goals (message)
			VALUES ($1);
		`

		_, err := db.Exec(query, message)
		if err != nil {
			resp := struct {
				Success bool   `json:"success"`
				Message string `json:"message"`
			}{
				Success: false,
				Message: "Failed",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}

		resp := struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Success: true,
			Message: "Success",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func main() {
	db := getDBConnection()
	http.HandleFunc("/", newHandler(db))
	http.ListenAndServe(":8000", nil)
}
