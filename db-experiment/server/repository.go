package server

import (
	repository "db-experiment/repositories"
	"github.com/jmoiron/sqlx"
)

type repositories struct {
	todoRepository repository.TodoRepository
}

func setupRepositories(db *sqlx.DB) *repositories {
	todoRepository := repository.InitializeTodoRepository(db)

	return &repositories{
		todoRepository: todoRepository,
	}
}