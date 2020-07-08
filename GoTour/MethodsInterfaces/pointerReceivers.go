package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) Scale(factor float64) {
	p.X = p.X * factor
	p.Y = p.Y * factor
}

func main() {
	p := Point{3, 4}
	p.Scale(5)
	fmt.Println(p.Abs())
}