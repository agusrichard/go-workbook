package main

import "fmt"

func Calculate(n int) int {
	result := n+2
	return result
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(Calculate(3))
}
