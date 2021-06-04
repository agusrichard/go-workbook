package main

import "fmt"

func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(Fib(3))
}
