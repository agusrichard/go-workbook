package main

import "fmt"

type I interface{
	P()
}

type S struct {
	name string
}

func (s *S) P() {
	if s == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(s.name)
}

func describe(i I) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	var i I
	var s *S
	describe(i)
	i = s
	describe(i)
	i.P()
	
	var t = &S{"Sekardayu"}
	describe(t)
	t.P()
}

