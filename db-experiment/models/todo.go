package model

import (
	"time"
)

type Todo struct {
	ID          int        `json:"id" db:"id"`
	Username    string     `json:"username" db:"username"`
	Title       string     `json:"title" db:"title"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	Description NullString `json:"description" db:"description"`
	ModifiedAt  NullTime   `json:"modifiedAt" db:"modified_at"`
}

func (t *Todo) IsValid() bool {
	return t.Username != "" && t.Title != ""
}
