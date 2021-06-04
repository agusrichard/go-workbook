package repository

import (
	"database/sql"
	model "db-experiment/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

type todoRepositoryV2 struct {
	db *sqlx.DB
}

type TodoRepositoryV2 interface {
	CreateTodo(todo *model.TodoModel) error
	GetAllTodos() (*[]model.TodoModel, error)
	GetTodoByID(todoID int) (*model.TodoModel, error)
	FilterTodos(filterQuery string, skip, take int) (*[]model.TodoModel, error)
	UpdateTodo(todo *model.TodoModel) error
	DeleteTodo(id int) error
}

func InitializeTodoRepositoryV2(db *sqlx.DB) TodoRepositoryV2 {
	return &todoRepositoryV2{
		db: db,
	}
}

func (r *todoRepositoryV2) CreateTodo(todo *model.TodoModel) error {
	if todo == nil {
		return errors.New("todo repository: create todo: todo is nil;")
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "todo repository: create todo: failed to initiate transaction;")
	}

	err = insertTodoV2(tx, todo)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "todo repository: create todo: failed to insert todo in repository;")
	}

	tx.Commit()

	return nil
}

func insertTodoV2(tx *sqlx.Tx, todo *model.TodoModel) error {
	_, err := tx.Exec(`
		INSERT INTO todos(username, title, description)
		VALUES ($1, $2, $3);
;	`, todo.Username, todo.Title, todo.Description.String)

	return err
}

func (r *todoRepositoryV2) GetAllTodos() (*[]model.TodoModel, error) {
	var todos []model.TodoModel

	query := `SELECT id, username, title, description, created_at, modified_at FROM todos`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "todo repository: get all todos: failed to query the data")
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, title string
		var description sql.NullString
		var createdAt time.Time
		var modifiedAt sql.NullTime

		err = rows.Scan(&id, &username, &title, &description, &createdAt, &modifiedAt)
		if err != nil {
			return nil, errors.Wrap(err, "todo repository: get all todos: failed scan the rows")
		}

		todos = append(todos, model.TodoModel{
			ID: id,
			Username: username,
			Title: title,
			Description: description,
			CreatedAt: createdAt,
			ModifiedAt: modifiedAt,
		})
	}

	return &todos, nil
}

func (r *todoRepositoryV2) GetTodoByID(todoID int) (*model.TodoModel, error) {
	row := r.db.QueryRow(`SELECT id, username, title, description, created_at, modified_at FROM todos WHERE id=$1;`, todoID)

	var id int
	var username, title string
	var description sql.NullString
	var createdAt time.Time
	var modifiedAt sql.NullTime

	err := row.Scan(&id, &username, &title, &description, &createdAt, &modifiedAt)
	if err != nil {
		return nil, errors.Wrap(err, "todo repository: get todo by id: failed to scan row")
	}

	result := model.TodoModel{
		ID: id,
		Username: username,
		Title: title,
		Description: description,
		CreatedAt: createdAt,
		ModifiedAt: modifiedAt,
	}

	return &result, nil
}

func (r *todoRepositoryV2) FilterTodos(filterQuery string, skip, take int) (*[]model.TodoModel, error) {
	var todos []model.TodoModel

	query := fmt.Sprintf(`
		SELECT id, username, title, description, created_at, modified_at
		FROM todos
		%s
		OFFSET %d
		LIMIT %d;
	`, filterQuery, skip, take)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "todo repository: get all todos: failed to query the data")
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, title string
		var description sql.NullString
		var createdAt time.Time
		var modifiedAt sql.NullTime

		err = rows.Scan(&id, &username, &title, &description, &createdAt, &modifiedAt)
		if err != nil {
			return nil, errors.Wrap(err, "todo repository: get all todos: failed scan the rows")
		}

		todos = append(todos, model.TodoModel{
			ID: id,
			Username: username,
			Title: title,
			Description: description,
			CreatedAt: createdAt,
			ModifiedAt: modifiedAt,
		})
	}

	return &todos, nil
}

func (r *todoRepositoryV2) UpdateTodo(todo *model.TodoModel) error {
	if todo == nil {
		return errors.New("todo repository: update todo: todo is nil;")
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "todo repository: update todo: failed to initiate transaction;")
	}

	err = updateTodoV2(tx, todo)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "todo repository: update todo: failed;")
	}

	tx.Commit()

	return nil
}

func updateTodoV2(tx *sqlx.Tx, todo *model.TodoModel) error {
	fmt.Println("todo model", todo)

	_, err := tx.Exec(`
		UPDATE todos
		SET username=$1,
		    title=$2,
		    description=$3
		WHERE id=$4;
	`,
		todo.Username,
		todo.Title,
		todo.Description.String,
		todo.ID,
	)

	fmt.Println("err", err)

	return err
}

func (r *todoRepositoryV2) DeleteTodo(id int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "todo repository: delete todo: failed to initiate transaction;")
	}

	err = deleteTodoV2(tx, id)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "todo repository: delete todo: failed;")
	}

	tx.Commit()

	return nil
}

func deleteTodoV2(tx *sqlx.Tx, id int) error {
	_, err := tx.Exec(`
		DELETE FROM todos WHERE id=$1;
	`, id)

	return err
}