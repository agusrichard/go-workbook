package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(arr[0:3])
	var some []int = arr[3:5]
	fmt.Println(some)
}