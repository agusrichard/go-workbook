package main

import "fmt"

func main() {
	avg := average(1, 2, 3, 4, 5)
	fmt.Println(avg)
}

func average(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	fmt.Println(numbers)
	var avg = float64(total) / float64(len(numbers))
	return avg
}
