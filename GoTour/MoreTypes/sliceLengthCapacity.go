package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	printSlice(arr)
	
	arr = arr[:0]
	printSlice(arr)
	arr = arr[:5]
	printSlice(arr)
	arr = arr[3:]
	printSlice(arr)
}

func printSlice(s []int) {
	fmt.Printf("len=%v, cap=%v, %v\n", len(s), cap(s), s)
}