package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
		case string:
		fmt.Printf("I am a string, and I am %v\n", v)
		case int:
		fmt.Printf("Two plus two is %v\n", v)
		default:
		fmt.Printf("What am I? %T\n", v)
	}
}


func main() {
	i := 10
	do(i)
	
	j := "Sekardayu"
	do(j)
}