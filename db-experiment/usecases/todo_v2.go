package usecase

import (
	"database/sql"
	model "db-experiment/models"
	repository "db-experiment/repositories"
	"fmt"
	"github.com/pkg/errors"
)

type todoUsecaseV2 struct {
	todoRepository repository.TodoRepositoryV2
}

type TodoUsecaseV2 interface {
	CreateTodo(todo *model.TodoShape) error
	GetAllTodos() (*[]model.TodoShape, error)
	GetTodoByID(id int) (*model.TodoShape, error)
	FilterTodos(filterQuery string, skip, take int) (*[]model.TodoShape, error)
	UpdateTodo(todo *model.TodoShape) error
	DeleteTodo(id int) error
}

func InitializeTodoUsecaseV2(r repository.TodoRepositoryV2) TodoUsecaseV2 {
	return &todoUsecaseV2{
		todoRepository: r,
	}
}


func (u *todoUsecaseV2) CreateTodo(todo *model.TodoShape) error {
	if todo == nil {
		return errors.New("todo usecase: create todo: todo is nil;")
	}

	if !todo.IsValid() {
		return errors.New("todo usecase: create todo: username and title can not be empty;")
	}

	todoModel := model.TodoModel{
		Username: todo.Username,
		Title: todo.Title,
		Description: sql.NullString{
			String: todo.Description,
		},
	}

	err := u.todoRepository.CreateTodo(&todoModel)
	if err != nil {
		return errors.Wrap(err, "todo usecase: create todo: failed to create todo in usecase;")
	}

	return nil
}

func (u *todoUsecaseV2) GetAllTodos() (*[]model.TodoShape, error) {
	var result []model.TodoShape
	todos, err := u.todoRepository.GetAllTodos()
	if err != nil {
		return nil, errors.Wrap(err, "todo usecase: get all todos: error get all todos;")
	}

	for _, value := range *todos {
		result = append(result, model.TodoShape{
			ID: value.ID,
			Username: value.Username,
			Title: value.Title,
			Description: value.Description.String,
			CreatedAt: value.CreatedAt,
			ModifiedAt: value.ModifiedAt.Time,
		})
	}

	return &result, nil
}

func (u *todoUsecaseV2) GetTodoByID(id int) (*model.TodoShape, error) {
	todo, err := u.todoRepository.GetTodoByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "todo usecase: get todo by id: error get data;")
	}
	result := model.TodoShape{
		ID: todo.ID,
		Username: todo.Username,
		Title: todo.Title,
		Description: todo.Description.String,
		CreatedAt: todo.CreatedAt,
		ModifiedAt: todo.ModifiedAt.Time,
	}

	return &result, nil
}

func (u *todoUsecaseV2) FilterTodos(filterQuery string, skip, take int) (*[]model.TodoShape, error) {
	var result []model.TodoShape

	todos, err := u.todoRepository.FilterTodos(filterQuery, skip, take)
	if err != nil {
		return nil, errors.Wrap(err, "todo usecase: filter todos: error get data")
	}

	for _, value := range *todos {
		result = append(result, model.TodoShape{
			ID: value.ID,
			Username: value.Username,
			Title: value.Title,
			Description: value.Description.String,
			CreatedAt: value.CreatedAt,
			ModifiedAt: value.ModifiedAt.Time,
		})
	}


	return &result, nil
}

func (u *todoUsecaseV2) UpdateTodo(todo *model.TodoShape) error {
	if todo == nil {
		return errors.New("todo usecase: create todo: todo is nil;")
	}

	if todo.ID == 0 {
		return errors.New("todo usecase: provide id to specify which todo needs to be updated")
	}

	fmt.Println("todo", todo)

	todoModel := model.TodoModel{
		ID: todo.ID,
		Username: todo.Username,
		Title: todo.Title,
		Description: sql.NullString{
			String: todo.Description,
		},
	}

	err := u.todoRepository.UpdateTodo(&todoModel)
	if err != nil {
		return errors.Wrap(err, "todo usecase: update todo: failed")
	}

	return nil
}

func (u *todoUsecaseV2) DeleteTodo(id int) error {
	err := u.todoRepository.DeleteTodo(id)
	if err != nil {
		return errors.Wrap(err, "todo usecase: delete todo: failed")
	}

	return nil
}