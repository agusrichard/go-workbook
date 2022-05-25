package main

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	Print([]int{1, 2, 3})
	Print([]string{"a", "b", "c"})
}
