package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
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
	Description sql.NullString `json:"description" db:"description"`
	ModifiedAt  sql.NullTime   `json:"modifiedAt" db:"modified_at"`
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

func ToConcrete(abstractPointer interface{}, emptyConcretePointer interface{}, baseStruct interface{}) interface{} {
	tVal := reflect.ValueOf(abstractPointer).Elem()
	cVal := reflect.New(reflect.TypeOf(emptyConcretePointer))

	numOfFields := cVal.Elem().NumField()
	for i := 0; i < numOfFields; i++ {
		if tVal.Field(i).Type() == reflect.TypeOf(baseStruct) {
			cVal.Elem().Field(i).Set(tVal.Field(i))
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(sql.NullString{}) {
			cVal.Elem().Field(i).SetString(tVal.Field(i).Interface().(sql.NullString).String)
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(sql.NullBool{}) {
			cVal.Elem().Field(i).SetBool(tVal.Field(i).Interface().(sql.NullBool).Bool)
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(sql.NullFloat64{}) {
			cVal.Elem().Field(i).SetFloat(tVal.Field(i).Interface().(sql.NullFloat64).Float64)
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(sql.NullInt64{}) {
			cVal.Elem().Field(i).SetInt(tVal.Field(i).Interface().(sql.NullInt64).Int64)
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(sql.NullTime{}) {
			cVal.Elem().Field(i).Set(reflect.ValueOf(tVal.Field(i).Interface().(sql.NullTime).Time))
			continue
		}
	}

	return cVal.Elem().Interface()
}

func main() {
	todo := Todo{
		TodoBase: TodoBase{
			ID:       1,
			Username: "sekardayu",
			Title:    "title",
			CreatedAt: time.Now(),
		},
		Description: sql.NullString{
			String: "description",
			Valid:  true,
		},
		ModifiedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
	}
	fmt.Println(todo)
	cTodo := todo.ToConcrete()
	fmt.Println(cTodo)

	result, err := json.Marshal(cTodo)
	if err != nil {
		panic(err)
	}

	println(string(result))
}
