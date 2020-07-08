package main

import (
	"fmt"
	"time"
)

type MyError struct{
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v with problem %v", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "What is wrong bruh?"}
}

func main() {
	v := run()
	fmt.Println(v)
}