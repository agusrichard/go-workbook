package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	pOne := Point{3, 4}
	fmt.Println(pOne.Distance())
	pTwo := Point{5, 12}
	fmt.Println(pTwo.Distance())
}