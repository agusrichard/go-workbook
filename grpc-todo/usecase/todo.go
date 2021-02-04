package usecase

import (
	"grpc-todo/models"
	"grpc-todo/repository"
	"log"
)

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

type TodoUsecase interface {
	CreateTodo(todo models.Todo) (bool, error)
}

func InitUserUsecase(todoRepository repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		todoRepository,
	}
}

func (todoUsecase *todoUsecase) CreateTodo(todo models.Todo) (bool, error) {
	_, err := todoUsecase.todoRepository.CreateTodo(todo)
	if err != nil {
		log.Println("Error create todo", err)
		return false, err
	}

	return true, nil
}
