package usecase

import (
	"encoding/json"
	"grpc-todo/models"
	"grpc-todo/repository"
	"log"
)

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

type TodoUsecase interface {
	CreateTodo(todo models.Todo) (bool, error)
	GetTodos(userID int64) (string, error)
	UpdateTodo(todo models.Todo) (bool, error)
	DeleteTodo(id int64) (bool, error)
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

func (todoUsecase *todoUsecase) GetTodos(userID int64) (string, error) {
	todos, err := todoUsecase.todoRepository.GetTodos(userID)
	if err != nil {
		log.Println("Error get todos", err)
		return "", err
	}

	result, err := json.Marshal(todos)

	return string(result), nil
}

func (todoUsecase *todoUsecase) UpdateTodo(todo models.Todo) (bool, error) {
	_, err := todoUsecase.todoRepository.UpdateTodo(todo)
	if err != nil {
		log.Println("Error update todo", err)
		return false, err
	}

	return true, nil
}

func (todoUsecase *todoUsecase) DeleteTodo(id int64) (bool, error) {
	_, err := todoUsecase.todoRepository.DeleteTodo(id)
	if err != nil {
		log.Println("Error delete todo", err)
		return false, err
	}

	return true, nil
}
