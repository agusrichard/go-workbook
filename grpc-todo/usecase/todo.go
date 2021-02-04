package usecase

import (
	"encoding/json"
	"grpc-todo/models"
	"grpc-todo/repository"
	"log"
	"strconv"
)

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

type TodoUsecase interface {
	CreateTodo(todo models.Todo) (bool, error)
	GetTodos(userID string) (string, error)
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

func (todoUsecase *todoUsecase) GetTodos(userID string) (string, error) {
	id, err := strconv.Atoi(userID)
	todos, err := todoUsecase.todoRepository.GetTodos(id)
	if err != nil {
		log.Println("Error get todos", err)
		return "", err
	}

	result, err := json.Marshal(todos)

	return string(result), nil
}
