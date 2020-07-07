package main

import "fmt"

func adder() func(float64) float64 {
	var sum float64 = 0
	return func(num float64) float64 {
		sum += num
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(float64(i)), neg(-float64(i)))
	}
}