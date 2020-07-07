package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	names := map[string]Person{
		"sekar": {"Sekardayu Hana Pradiani", 22},
		"saskia": {"Saskia Nurul Azhima", 20},
	}
	
	fmt.Println(names)
}