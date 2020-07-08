package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var a interface{}
	describe(a)
	
	b := 21
	describe(b)
	
	c := "Sekardayu"
	describe(c)
}