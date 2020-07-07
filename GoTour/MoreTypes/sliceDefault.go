package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[:])
}