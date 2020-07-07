package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	sekar := Person{"Sekardayu Hana Pradiani", 22}
	fmt.Println(sekar.name)
	girlfriend := &sekar
	fmt.Println(girlfriend)
	girlfriend.age =17
	fmt.Println(girlfriend)
	fmt.Println(sekar)
}
