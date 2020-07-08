package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) Scale(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

func ScaleFunc(p *Point, f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

func main() {
	p1 := Point{3, 4}
	p1.Scale(5)
	fmt.Println(p1)
	ScaleFunc(&p1, 5)
	fmt.Println(p1)
	
	p2 := &Point{5, 12}
	p2.Scale(3)
	fmt.Println(p2)
	ScaleFunc(p2, 3)
	fmt.Println(p2)
}