package model

import (
	"time"
)

type TodoBase struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Todo struct {
	TodoBase
	Description NullString `json:"description" db:"description"`
	ModifiedAt  NullTime   `json:"modifiedAt" db:"modified_at"`
}

type TodoConcrete struct {
	TodoBase
	Description string    `json:"description" db:"description"`
	ModifiedAt  time.Time `json:"modifiedAt" db:"modified_at"`
}

func (t *Todo) IsValid() bool {
	return t.Username != "" || t.Title != ""
}

func (t *Todo) ToConcrete() TodoConcrete {
	return ToConcrete(t, TodoConcrete{}, TodoBase{}).(TodoConcrete)
}
