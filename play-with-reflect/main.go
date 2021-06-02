package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

type UserBase struct {
	Username string `json:"username" db:"username"`
}

type User struct {
	UserBase
	Fullname sql.NullString `json:"fullname" db:"fullname"`
}

type user struct {
	UserBase
	Fullname string `json:"username" db:"username"`
}

type TodoBase struct {
	ID    int    `json:"id" db:"id"`
	User  User   `json:"user" db:"user"`
	Title string `json:"title" db:"title"`
}

type Todo struct {
	TodoBase
	Description sql.NullString `json:"description" db:"description"`
}

type todo struct {
	TodoBase
	Description string `json:"description" db:"description"`
}

func (t *Todo) ToConcrete() todo {
	return ToConcrete(t, &todo{}).(todo)
}

func ToConcrete(abstractValue interface{}, concreteType interface{}) interface{} {
	todoValue := reflect.ValueOf(abstractValue).Elem()
	concreteValue := reflect.ValueOf(concreteType)

	numOfFields := concreteValue.Elem().NumField()
	for i := 0; i < numOfFields; i++ {
		if todoValue.Field(i).Type() == reflect.TypeOf(TodoBase{}) {
			concreteValue.Elem().Field(i).Set(todoValue.Field(i))
			continue
		}

		if todoValue.Field(i).Type() == reflect.TypeOf(sql.NullString{}) {
			concreteValue.Elem().Field(i).SetString(todoValue.Field(i).Interface().(sql.NullString).String)
			continue
		}

		//if todoValue.Field(i).Kind() == reflect.Struct {
		//
		//}
	}

	return concreteValue.Elem().Interface()
}

func main() {
	todo := Todo{
		TodoBase: TodoBase{
			ID:       1,
			User: User{
				UserBase: UserBase{
					Username: "sekardayu",
				},
				Fullname: sql.NullString{
					String: "Sekardayu Hana Pradiani",
					Valid: true,
				},
			},
			Title:    "title",
		},
		Description: sql.NullString{
			String: "description",
			Valid:  true,
		},
	}
	fmt.Println(todo)
	newTodo := todo.ToConcrete()
	fmt.Println(newTodo)

	result, err := json.Marshal(newTodo)
	if err != nil {
		panic(err)
	}

	println(string(result))
}
