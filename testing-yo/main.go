package main

import "fmt"

type MyInt struct {
	number int
}

func (m *MyInt) Useless() int {
	return m.number
}

func UsingUseless(m *MyInt) int {
	return m.Useless()
}

func main() {
	fmt.Println("Hello World")
}
