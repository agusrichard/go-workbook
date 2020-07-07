package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	names := map[string]Person{
		"sekar": Person{
			"Sekardayu Hana Pradiani",
			22,
		},
		"saskia": Person{
			"Saskia Nurul Azhima",
			20,
		},
	}
	
	fmt.Println(names)
	fmt.Println(names["sekar"])
	fmt.Println(names["saskia"])
}