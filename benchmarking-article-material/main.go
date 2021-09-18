package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Calculate(x, y int) int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return x + y
}
func main() {
	fmt.Println(Calculate(1, 2))
}
