package server

import (
	repository "db-experiment/repositories"
	"github.com/jmoiron/sqlx"
)

type repositories struct {
	todoRepository repository.TodoRepository
	todoRepositoryV2 repository.TodoRepositoryV2
}

func setupRepositories(db *sqlx.DB) *repositories {
	todoRepository := repository.InitializeTodoRepository(db)
	todoRepositoryV2 := repository.InitializeTodoRepositoryV2(db)

	return &repositories{
		todoRepository: todoRepository,
		todoRepositoryV2: todoRepositoryV2,
	}
}