package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b Abser
	x := MyFloat(math.Sqrt2)
	y := Point{3, 4}
	
	fmt.Println(x, y)
	a = &x
	b = &y
	
	fmt.Println(a.Abs())
	fmt.Println(b.Abs())
}

type Abser interface {
	Abs() float64
}

type Point struct {
	X, Y float64
}

type MyFloat float64

func (f *MyFloat) Abs() float64 {
	if *f < 0 {
		return -float64(*f)
	}
	return float64(*f)
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}