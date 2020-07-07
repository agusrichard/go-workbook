package main

import "fmt"

func main() {
	arr1 := []int{21, 9, 97}
	fmt.Println(arr1)
	
	arr2 := []struct{
		name string
		age int
	}{
		{"Sekar", 22},
		{"Saskia", 20},
	}
	
	fmt.Println(arr2)
}	