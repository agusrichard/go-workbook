package main

import "fmt"

func main() {
	var i interface{} = "Sekardayu"
	
	a := i.(string)
	fmt.Println(a)
	
	a, ok := i.(string)
	fmt.Println(a, ok)
	
	b, ok := i.(int)
	fmt.Println(b, ok)
	
	b = i.(int)			// Panic
	fmt.Println(b)
}