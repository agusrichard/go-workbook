package main

import (
	"fmt"
	"math/rand"
	"time"
)

func expensiveComputation() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func sleepyHead() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		dur := time.Duration(rand.Intn(250))
		time.Sleep(dur * time.Millisecond)
	}
}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
		time.Sleep(1000 * time.Millisecond)
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
		time.Sleep(500 * time.Millisecond)
	}
}

func printer(c <-chan string) {
	for {
		fmt.Println(<-c)
	}
}

func hello(quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			return
		default:
			fmt.Println("Hello")
		}
	}
}

func main() {
	fmt.Println("Running")
	defer func() {
		fmt.Println("Done")
	}()

	go expensiveComputation()

	for i := 0; i < 10; i++ {
		go sleepyHead()
	}

	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	quit := make(chan struct{})
	go hello(quit)
	time.Sleep(1 * time.Second)
	quit <- struct{}{}

	var input string
	fmt.Scanln(&input)
}
