package main

import "fmt"

func main() {
	x, y := 14, 21
	fmt.Println(x, y)
	
	i := &x
	j := &y
	fmt.Println(i, j)
	fmt.Println(*i, *j)
	
	*i = *i + 7
	*j = *j + 7
	fmt.Println(x, y, *i, *j)
	
}