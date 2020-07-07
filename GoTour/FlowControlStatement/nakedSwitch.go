package main

import "fmt"

func main() {
	num := 100
	switch {
		case num < 10:
		fmt.Println("Less than 10")
		case num < 20:
		fmt.Println("Greater or equal to ten and less than 20")
		default:
		fmt.Println("What are you talking about!")
	}
}