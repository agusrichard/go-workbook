package main

import (
	"fmt"
	"math"
)

func square_root(x float64) (y string) {
	if x < 0 {
		return square_root(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(square_root(-4), square_root(4))	
}