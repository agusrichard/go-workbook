package main

import "fmt"

func main() {
	switch num := 20; num {
		case 10:
			fmt.Println("Ten")
		case 20:
			fmt.Println("Twenty")
		default:
			fmt.Println("Default")
	}
}