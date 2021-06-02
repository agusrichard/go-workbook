package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type Todo struct {
	ID          int            `json:"id" db:"id"`
	Username    string         `json:"username" db:"username"`
	Title       string         `json:"title" db:"title"`
	CreatedAt   time.Time      `json:"createdAt" db:"created_at"`
	Description sql.NullString `json:"description" db:"description"`
	ModifiedAt  sql.NullTime   `json:"modifiedAt" db:"modified_at"`
}

func (t *Todo) ToNotNullable() interface{} {
	return InitilizeNotNullStruct(t)
}

func InitilizeNotNullStruct(nullableStruct interface{}) interface{} {
	fmt.Println("InitilizeNotNullStruct")
	var sfs []reflect.StructField

	nst := reflect.TypeOf(nullableStruct).Elem()
	nsv := reflect.ValueOf(nullableStruct).Elem()

	fmt.Println(sfs)

	numFields := nsv.NumField()
	for i := 0; i < numFields; i++ {
		var t reflect.Type

		switch reflect.TypeOf(nst.Field(i)) {
		case reflect.TypeOf(sql.NullString{}):
			t = reflect.TypeOf("")
		case reflect.TypeOf(sql.NullBool{}):
			t = reflect.TypeOf(true)
		case reflect.TypeOf(sql.NullFloat64{}):
			t = reflect.TypeOf("0.1")
		case reflect.TypeOf(sql.NullInt64{}):
			t = reflect.TypeOf("0")
		case reflect.TypeOf(sql.NullTime{}):
			t = reflect.TypeOf(time.Now())
		default:
			t = nst.Field(i).Type
		}
		sfs = append(sfs, reflect.StructField{
			Name: nst.Field(i).Name,
			Type: t,
			Tag:  nst.Field(i).Tag,
		})
	}

	fmt.Println(sfs)

	return nil
}

func main() {
	todo := Todo{
		ID:        1,
		Username:  "sekardayu",
		Title:     "title",
		CreatedAt: time.Now(),
		Description: sql.NullString{
			String: "description",
			Valid:  true,
		},
		ModifiedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	fmt.Println(todo)
	cTodo := todo.ToNotNullable()
	fmt.Println(cTodo)

	//result, err := json.Marshal(cTodo)
	//if err != nil {
	//	panic(err)
	//}
	//
	//println(string(result))
}
