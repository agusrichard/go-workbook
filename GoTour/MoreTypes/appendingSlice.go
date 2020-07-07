package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}	
	printSlice(arr)
	
	arr = append(arr, 4)
	printSlice(arr)
	
	arr = append(arr, 5, 6, 7)
	printSlice(arr)
}

func printSlice(s []int) {
	fmt.Println(len(s), cap(s), s)	
}