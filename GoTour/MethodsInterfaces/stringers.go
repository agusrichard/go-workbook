package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v, and I am %v years old.", p.Name, p.Age)
}

func main() {
	sekar := Person{"Sekardayu Hana Pradiani", 22}
	saskia := Person{"Saskia Nurul Azhima", 20}
	
	fmt.Println(sekar)
	fmt.Println(saskia)
}