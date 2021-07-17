package main

import "fmt"

func addition(x, y int) int {
	return x + y
}

func curryAddition(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	fmt.Println(addition(1, 2))

	add1 := curryAddition(1)
	add10 := add1(10)
	fmt.Println(add10)
}
