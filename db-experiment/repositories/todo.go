package repository

import (
	model "db-experiment/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type todoRepository struct {
	db *sqlx.DB
}

type TodoRepository interface {
	CreateTodo(todo *model.Todo) error
}

func InitializeTodoRepository(db *sqlx.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}


func (r *todoRepository) CreateTodo(todo *model.Todo) error {
	if todo == nil {
		return errors.New("create todo: todo is nil")
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "create todo: failed to instantiate transaction")
	}

	err = insertTodo(tx, todo)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "create todo: failed to insert todo in repository")
	}

	tx.Commit()

	return nil
}

func insertTodo(tx *sqlx.Tx, todo *model.Todo) error {
	_, err := tx.NamedExec(`
		INSERT INTO todos(username, title, description)
		VALUES (:username, :title, :description)
;	`, todo)

	return err
}