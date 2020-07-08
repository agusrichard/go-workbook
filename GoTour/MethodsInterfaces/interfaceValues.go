package main

import "fmt"

type I interface {
	P()
}

type S struct {
	name string
}

func (s *S) P() {
	fmt.Println(s.name)
}

type F float64

func (f F) P() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("Value: %v, Type: %T", i, i)
}

func main() {
	var i I
	i = &S{"Sekar"}
	describe(i)
	i.P()
	
	var j I
	j = F(-14)
	describe(j)
	j.P()
}

