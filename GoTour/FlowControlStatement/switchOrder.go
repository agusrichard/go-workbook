package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	switch today {
	case today + 0:
		fmt.Println("Today wohooo!")
	default:
		fmt.Println("I dont't know when")

	}
}
