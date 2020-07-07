package main

import "fmt"

func something(num int) (x, y int) {
	x = num - 5
	y = num + 5
	return 
}

func main() {
	fmt.Println(something(10))	
}

