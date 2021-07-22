package model

import (
	"time"
)

type Todo struct {
	ID           int          `json:"id" db:"id"`
	Username     string       `json:"username" db:"username"`
	Title        string       `json:"title" db:"title"`
	Description  MyNullString `json:"description" db:"description"`
	Deadline     NullTime     `json:"deadline" db:"deadline"`
	IsImportant  NullBool     `json:"is_important" db:"is_important"`
	BudgetAmount NullFloat64  `json:"budget_amount" db:"budget_amount"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	ModifiedAt   NullTime     `json:"modofied_at" db:"modified_at"`
}

func (t *Todo) IsValid() bool {
	return t.Username != "" && t.Title != ""
}
