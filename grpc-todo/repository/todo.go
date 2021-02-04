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
