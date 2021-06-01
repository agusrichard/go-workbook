package usecase

import repository "db-experiment/repositories"

type todoUsecase struct {
	r repository.TodoRepository
}

type TodoUsecase interface {

}

func InitializeTodoUsecase(r repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		r: r,
	}
}
