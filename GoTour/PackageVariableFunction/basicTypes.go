package main

import (
	"fmt"
	"math/cmplx"
)

var (
	agus string = "Agus Richard Lubis"
	age int = 22
	compNum complex128 = cmplx.Sqrt(8 + 9i)
)

func main() {
	fmt.Printf("Type %T Value %v\n", agus, agus)
	fmt.Printf("Type %T Value %v\n", compNum, compNum)
}