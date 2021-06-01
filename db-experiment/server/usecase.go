package server

import "db-experiment/usecases"

type usecases struct {
	todoUsecase usecase.TodoUsecase
}

func setupUsecases(r *repositories) *usecases {
	todoUsecase := usecase.InitializeTodoUsecase(r.todoRepository)

	return &usecases{
		todoUsecase: todoUsecase,
	}
}
