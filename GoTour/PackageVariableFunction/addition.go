package main

import "fmt"

func add0(x int, y int) int {
	return x + y	
}

func add1(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add0(21, 14))	
}