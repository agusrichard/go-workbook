package main

import "fmt"

func something(name string, age int) (string, int) {
	return "Hello my name is " + name, age + 5
}

func main() {
	fmt.Println(something("Agus Richard Lubis", 22))
}