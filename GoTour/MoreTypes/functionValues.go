package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)	
}

func main() {
	hypothenuse := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	
	fmt.Println(hypothenuse(3, 4))
	fmt.Println(compute(hypothenuse))
	fmt.Println(compute(math.Pow))
}