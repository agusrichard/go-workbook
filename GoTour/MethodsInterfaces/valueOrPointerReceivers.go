package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) Scale(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

func main() {
	p := Point{3, 4}
	fmt.Printf("Value before scaling: %f\n", p.Abs())
	p.Scale(5)
	fmt.Printf("Value after scaling: %f\n", p.Abs())
}