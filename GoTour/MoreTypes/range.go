package main

import "fmt"

var doubles = []int{2, 4, 6, 8, 10}

func main() {
	for i, v := range doubles {
		fmt.Printf("2*%d = %d\n", i+1, v)	
	}
}