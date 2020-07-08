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

func AbsFunc(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p1 := Point{3, 4}
	fmt.Println(p1.Abs())
	fmt.Println(AbsFunc(p1))
	
	p2 := &Point{5, 12}
	fmt.Println(p2.Abs())
	fmt.Println(AbsFunc(*p2))
}

