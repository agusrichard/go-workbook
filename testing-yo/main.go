package main

import "fmt"

type MyInt struct {
	number int
}

type MyIntI interface {
	Useless() int
}

func (m *MyInt) Useless() int {
	return m.number
}

func UsingUseless(m MyIntI) int {
	return m.Useless()
}

func main() {
	fmt.Println("Hello World")
}
