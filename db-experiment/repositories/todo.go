package repository

import "github.com/jmoiron/sqlx"

type todoRepository struct {
	db *sqlx.DB
}

type TodoRepository interface {

}

func InitializeTodoRepository(db *sqlx.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}