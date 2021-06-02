package model

import (
	"reflect"
)

func ToConcrete(abstractPointer interface{}, emptyConcretePointer interface{}, baseStruct interface{}) interface{} {
	tVal := reflect.ValueOf(abstractPointer).Elem()
	cVal := reflect.New(reflect.TypeOf(emptyConcretePointer))

	numOfFields := cVal.Elem().NumField()
	for i := 0; i < numOfFields; i++ {
		if tVal.Field(i).Type() == reflect.TypeOf(baseStruct) {
			cVal.Elem().Field(i).Set(tVal.Field(i))
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(NullString{}) {
			cVal.Elem().Field(i).SetString(tVal.Field(i).Interface().(NullString).String)
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(NullBool{}) {
			cVal.Elem().Field(i).SetBool(tVal.Field(i).Interface().(NullBool).Bool)
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(NullFloat64{}) {
			cVal.Elem().Field(i).SetFloat(tVal.Field(i).Interface().(NullFloat64).Float64)
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(NullInt64{}) {
			cVal.Elem().Field(i).SetInt(tVal.Field(i).Interface().(NullInt64).Int64)
			continue
		}

		if tVal.Field(i).Type() == reflect.TypeOf(NullTime{}) {
			cVal.Elem().Field(i).Set(reflect.ValueOf(tVal.Field(i).Interface().(NullTime).Time))
			continue
		}
	}

	return cVal.Elem().Interface()
}