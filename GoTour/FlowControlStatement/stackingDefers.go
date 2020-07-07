package main

import "fmt"

func main() {
	fmt.Println("First")
	
	for i := 0; i < 10; i++ {
		defer fmt.Println("Ten times last")	
	}
	
	fmt.Println("Second")
}