package model

import (
	"database/sql"
	"time"
)

type TodoShape struct {
	ID          int       `json:"id" db:"id"`
	Username    string    `json:"username" db:"username"`
	Title       string    `json:"title" db:"title"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	Description string    `json:"description" db:"description"`
	ModifiedAt  time.Time `json:"modifiedAt" db:"modified_at"`
}

type TodoModel struct {
	ID          int            `json:"id" db:"id"`
	Username    string         `json:"username" db:"username"`
	Title       string         `json:"title" db:"title"`
	CreatedAt   time.Time      `json:"createdAt" db:"created_at"`
	Description sql.NullString `json:"description" db:"description"`
	ModifiedAt  sql.NullTime   `json:"modifiedAt" db:"modified_at"`
}

func (t *TodoShape) IsValid() bool {
	return t.Username != "" || t.Title != ""
}