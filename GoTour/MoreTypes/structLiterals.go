package main

import "fmt" 

type Person struct {
	name string
	age int
}

func main() {
	person1 := Person{"Agus", 22}
	person2 := Person{name: "Sekar"}
	person3 := &Person{"Saskia", 20}
	person$ := Person{}
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person4)
}