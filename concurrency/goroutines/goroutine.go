package main

import (
	"fmt"
)

func expensiveComputation() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Running")
	defer func() {
		fmt.Println("Done")
	}()

	go expensiveComputation()

	var input string
	fmt.Scanln(&input)
}
