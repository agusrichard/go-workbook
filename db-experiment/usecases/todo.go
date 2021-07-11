package usecase

import (
	model "db-experiment/models"
	repository "db-experiment/repositories"
	"fmt"
	"github.com/pkg/errors"
)

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

type TodoUsecase interface {
	CreateTodo(todo *model.Todo) error
	GetAllTodos() (*[]model.Todo, error)
	GetTodoByID(id int) (*model.Todo, error)
	FilterTodos(filterQuery string, skip, take int) (*[]model.Todo, error)
	UpdateTodo(todo *model.Todo) error
	DeleteTodo(id int) error
}

func InitializeTodoUsecase(r repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		todoRepository: r,
	}
}


func (u *todoUsecase) CreateTodo(todo *model.Todo) error {
	if todo == nil {
		return errors.New("todo usecase: create todo: todo is nil;")
	}

	if !todo.IsValid() {
		fmt.Println("here")
		return errors.New("todo usecase: create todo: username and title can not be empty;")
	}

	err := u.todoRepository.CreateTodo(todo)
	if err != nil {
		return errors.Wrap(err, "todo usecase: create todo: failed to create todo in usecase;")
	}

	return nil
}

func (u *todoUsecase) GetAllTodos() (*[]model.Todo, error) {
	todos, err := u.todoRepository.GetAllTodos()
	if err != nil {
		return nil, errors.Wrap(err, "todo usecase: get all todos: error get all todos;")
	}

	return todos, nil
}

func (u *todoUsecase) GetTodoByID(id int) (*model.Todo, error) {
	todo, err := u.todoRepository.GetTodoByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "todo usecase: get todo by id: error get data;")
	}

	return todo, nil
}

func (u *todoUsecase) FilterTodos(filterQuery string, skip, take int) (*[]model.Todo, error) {
	todos, err := u.todoRepository.FilterTodos(filterQuery, skip, take)
	if err != nil {
		return nil, errors.Wrap(err, "todo usecase: filter todos: error get data")
	}

	return todos, nil
}

func (u *todoUsecase) UpdateTodo(todo *model.Todo) error {
	if todo == nil {
		return errors.New("todo usecase: create todo: todo is nil;")
	}

	if todo.ID == 0 {
		return errors.New("todo usecase: provide id to specify which todo needs to be updated")
	}

	err := u.todoRepository.UpdateTodo(todo)
	if err != nil {
		return errors.Wrap(err, "todo usecase: update todo: failed")
	}

	return nil
}

func (u *todoUsecase) DeleteTodo(id int) error {
	err := u.todoRepository.DeleteTodo(id)
	if err != nil {
		return errors.Wrap(err, "todo usecase: delete todo: failed")
	}

	return nil
}