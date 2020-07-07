package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 12
	var y int = 13
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(x, y, f)
}