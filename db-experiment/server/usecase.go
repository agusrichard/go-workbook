package server

import "db-experiment/usecases"

type usecases struct {
	todoUsecase usecase.TodoUsecase
	todoUsecaseV2 usecase.TodoUsecaseV2
}

func setupUsecases(r *repositories) *usecases {
	todoUsecase := usecase.InitializeTodoUsecase(r.todoRepository)
	todoUsecaseV2 := usecase.InitializeTodoUsecaseV2(r.todoRepositoryV2)

	return &usecases{
		todoUsecase: todoUsecase,
		todoUsecaseV2: todoUsecaseV2,
	}
}
