package main

import "fmt"

const eulerNumber = 2.71828

func main() {
	eulerNumber = 3.14		// Can't reassign
	fmt.Println(eulerNumber)
}