package usecase

import (
	model "db-experiment/models"
	repository "db-experiment/repositories"
	"github.com/pkg/errors"
)

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

type TodoUsecase interface {
	CreateTodo(todo *model.Todo) error
}

func InitializeTodoUsecase(r repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		todoRepository: r,
	}
}


func (u *todoUsecase) CreateTodo(todo *model.Todo) error {
	if todo == nil {
		return errors.New("create todo: todo is nil")
	}

	if !todo.IsValid() {
		return errors.New("create todo: username and title can not be empty")
	}

	err := u.todoRepository.CreateTodo(todo)
	if err != nil {
		return errors.Wrap(err, "create todo: failed to create todo in usecase")
	}

	return nil
}