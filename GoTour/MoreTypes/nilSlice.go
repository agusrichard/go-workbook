package main

import "fmt"

func main() {
	var arr []int
	fmt.Println(arr, len(arr), cap(arr))
	
	if arr == nil {
		fmt.Println("Nilllllllll!!!!!!!!!!")	
	} else {
		fmt.Println("Relax bruh!")
	}
					
}