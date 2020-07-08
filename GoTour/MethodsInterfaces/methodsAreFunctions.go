package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y) 	
}

func main() {
	p := Point{5, 12}
	fmt.Println(Distance(p))
}