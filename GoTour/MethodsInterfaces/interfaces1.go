package main

import "fmt"

func main() {
	var a LivingBeing = Human{"Sekar", 22}
	var b LivingBeing = Animal{"Dog", 4}

	fmt.Println(a)
	fmt.Println(b)
}

type LivingBeing interface {
	Eat() string
}

type Human struct {
	name string
	age  int
}

type Animal struct {
	species string
	leg     int
}

func (h Human) Eat() string {
	return "I will eat everything!"
}

func (a Animal) Eat() string {
	return "Munch... Munch... Munch..."
}
