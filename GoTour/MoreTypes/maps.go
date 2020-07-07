package main

import "fmt"

type Person struct {
	name string
	age int
}

var m map[string]Person

func main() {
	fmt.Println(m)
	m = make(map[string]Person)
	m["sekar"] = Person{"Sekardayu Hana Pradiani", 22}
	fmt.Println(m)
}