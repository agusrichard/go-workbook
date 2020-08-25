package main

import "fmt"

func main() {
	var num1 int = 21
	var num2 *int = &num1

	fmt.Println(num1)
	fmt.Println(num2)
}
