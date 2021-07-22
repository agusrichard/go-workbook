package repository

import (
	model "db-experiment/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type todoRepository struct {
	db *sqlx.DB
}

type TodoRepository interface {
	CreateTodo(todo *model.Todo) error
	GetAllTodos() (*[]model.Todo, error)
	GetTodoByID(id int) (*model.Todo, error)
	FilterTodos(filterQuery string, skip, take int) (*[]model.Todo, error)
	UpdateTodo(todo *model.Todo) error
	DeleteTodo(id int) error
}

func InitializeTodoRepository(db *sqlx.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) CreateTodo(todo *model.Todo) error {
	if todo == nil {
		return errors.New("todo repository: create todo: todo is nil;")
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "todo repository: create todo: failed to initiate transaction;")
	}

	err = insertTodo(tx, todo)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "todo repository: create todo: failed to insert todo in repository;")
	}

	tx.Commit()

	return nil
}

func insertTodo(tx *sqlx.Tx, todo *model.Todo) error {
	_, err := tx.NamedExec(`
		INSERT INTO todos(username, title, description, deadline, is_important, budget_amount)
		VALUES (:username, :title, :description, :deadline, :is_important, :budget_amount);
;	`, todo)

	return err
}

func (r *todoRepository) GetAllTodos() (*[]model.Todo, error) {
	var todos []model.Todo

	query := `SELECT id, username, title, description, deadline, is_important, budget_amount, created_at, modified_at FROM todos`
	err := r.db.Select(&todos, query)
	if err != nil {
		return nil, errors.Wrap(err, "todo repository: get all todo: failed")
	}

	return &todos, nil
}

func (r *todoRepository) GetTodoByID(id int) (*model.Todo, error) {
	var todo model.Todo

	err := r.db.Get(&todo, `SELECT id, username, title, description, deadline, is_important, budget_amount, created_at, modified_at FROM todos WHERE id=$1;`, id)
	if err != nil {
		return nil, errors.Wrap(err, "todo repository: get todo by id: failed")
	}

	return &todo, nil
}

func (r *todoRepository) FilterTodos(filterQuery string, skip, take int) (*[]model.Todo, error) {
	var todos []model.Todo

	query := fmt.Sprintf(`
		SELECT id, username, title, description, created_at, modified_at
		FROM todos
		%s
		OFFSET %d
		LIMIT %d;
	`, filterQuery, skip, take)

	fmt.Println("query", query)

	err := r.db.Select(&todos, query)
	if err != nil {
		return nil, errors.Wrap(err, "todo repository: get all todo: failed")
	}

	return &todos, nil
}

func (r *todoRepository) UpdateTodo(todo *model.Todo) error {
	if todo == nil {
		return errors.New("todo repository: update todo: todo is nil;")
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "todo repository: update todo: failed to initiate transaction;")
	}

	err = updateTodo(tx, todo)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "todo repository: update todo: failed;")
	}

	tx.Commit()

	return nil
}

func updateTodo(tx *sqlx.Tx, todo *model.Todo) error {
	_, err := tx.NamedExec(`
		UPDATE todos
		SET username=:username,
		    title=:title,
		    description=:description
			deadline=:deadline,
			is_important=:is_important,
			budget_amount=:budget_amount
		WHERE id=:id;
	`, todo)

	return err
}

func (r *todoRepository) DeleteTodo(id int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "todo repository: delete todo: failed to initiate transaction;")
	}

	err = deleteTodo(tx, id)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "todo repository: delete todo: failed;")
	}

	tx.Commit()

	return nil
}

func deleteTodo(tx *sqlx.Tx, id int) error {
	_, err := tx.Exec(`
		DELETE FROM todos WHERE id=$1;
	`, id)

	return err
}
