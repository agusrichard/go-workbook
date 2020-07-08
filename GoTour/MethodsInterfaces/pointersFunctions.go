package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Abs(p Point) float64{
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func Scale(p *Point, factor float64) {
	p.X = p.X * factor
	p.Y = p.Y * factor
}

func main() {
	p := Point{3, 4}
	fmt.Println(Abs(p))
	Scale(&p, 5)
	fmt.Println(p)
}