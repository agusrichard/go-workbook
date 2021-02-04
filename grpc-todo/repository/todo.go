package repository

import (
	"database/sql"
	"grpc-todo/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type todoRepository struct {
	db *sqlx.DB
}

type TodoRepository interface {
	CreateTodo(todo models.Todo) (bool, error)
	GetTodos(userID int64) ([]models.Todo, error)
	UpdateTodo(todo models.Todo) (bool, error)
	DeleteTodo(id int64) (bool, error)
}

func InitTodoRepository(db *sqlx.DB) TodoRepository {
	return &todoRepository{
		db,
	}
}

func (todoRepository *todoRepository) CreateTodo(todo models.Todo) (bool, error) {
	var err error
	var result bool

	tx, errTx := todoRepository.db.Begin()
	if errTx != nil {
		log.Println("Error create todo: ", errTx)
	} else {
		err = insertTodo(tx, todo)
		if err != nil {
			log.Println("Error create todo: ", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		log.Println("Error create todo: ", err)
	}

	return result, err
}

func insertTodo(tx *sql.Tx, todo models.Todo) error {
	_, err := tx.Exec(`
	INSERT INTO todos (
		title,
		description,
		user_id
	)
	VALUES(
		$1,
		$2,
		$3
	);
	`,
		todo.Title,
		todo.Description,
		todo.UserID,
	)

	return err
}

func (todoRepository *todoRepository) GetTodos(userID int64) ([]models.Todo, error) {
	var todos []models.Todo
	rows, err := todoRepository.db.Query(`
		SELECT id, title, description FROM todos WHERE user_id=$1;
	`, userID)

	if err != nil {
		log.Println("Error get todos", err)
		return []models.Todo{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&(todo.ID), &(todo.Title), &(todo.Description))
		if err != nil {
			log.Println("Error get todos", err)
			return []models.Todo{}, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (todoRepository *todoRepository) UpdateTodo(todo models.Todo) (bool, error) {
	var err error
	var result bool

	tx, errTx := todoRepository.db.Begin()
	if errTx != nil {
		log.Println("Error update todo: ", errTx)
	} else {
		err = updateTodo(tx, todo)
		if err != nil {
			log.Println("Error update todo: ", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		log.Println("Error update todo: ", err)
	}

	return result, err
}

func updateTodo(tx *sql.Tx, todo models.Todo) error {
	_, err := tx.Exec(`
	UPDATE todos
	SET
		title=$1,
		description=$2
	WHERE id=$3
	
	`,
		todo.Title,
		todo.Description,
		todo.ID,
	)

	return err
}

func (todoRepository *todoRepository) DeleteTodo(id int64) (bool, error) {
	var err error
	var result bool

	tx, errTx := todoRepository.db.Begin()
	if errTx != nil {
		log.Println("Error to delete todo", err)
	} else {
		err = deleteTodo(tx, id)
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		log.Println("Error to delete todo", err)
	}

	return result, err
}

func deleteTodo(tx *sql.Tx, id int64) error {
	_, err := tx.Exec(`
		DELETE FROM todos
		WHERE id=$1
	`, id,
	)

	return err
}
