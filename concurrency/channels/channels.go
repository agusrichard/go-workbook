package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func CalculateValue(values chan int) {
// 	value := rand.Intn(10)
// 	fmt.Printf("Calculated Random Value: %d\n", value)
// 	values <- value
// }

// func main() {
// 	fmt.Println("Go channel tutorial")

// 	values := make(chan int)
// 	defer close(values)

// 	go CalculateValue(values)
// 	value := <-values
// 	fmt.Println("here")
// 	fmt.Println(value)
// }

// UNBUFFERED CHANNEL
// func CalculateValue(c chan int) {
// 	value := rand.Intn(10)
// 	fmt.Println("Calculated Random Value: {}", value)
// 	time.Sleep(1000 * time.Millisecond)
// 	c <- value
// 	fmt.Println("This executes regardless as the send is now non-blocking", value)
// }

// func main() {
// 	fmt.Println("Go Channel Tutorial")

// 	valueChannel := make(chan int, 2)
// 	defer close(valueChannel)

// 	go CalculateValue(valueChannel)
// 	go CalculateValue(valueChannel)

// 	values := <-valueChannel
// 	fmt.Println(values)

// 	time.Sleep(1000 * time.Millisecond)
// }

func main() {
	fmt.Println("Working in channels")

	name := make(chan string)
	defer close(name)

	go func(name chan string) {
		fmt.Println("Goroutine one initialized!")
		dur := time.Millisecond * time.Duration(rand.Intn(5000))
		time.Sleep(dur)
		name <- "Sekardayu Hana Pradiani"
		fmt.Println("Goroutine one done!", dur)
	}(name)

	age := make(chan int)
	defer close(age)
	go func(age chan int) {
		fmt.Println("Goroutine two initialized!")
		dur := time.Millisecond * time.Duration(rand.Intn(500))
		time.Sleep(dur)
		age <- 21
		fmt.Println("Goroutine two done!", dur)
	}(age)

	dob := make(chan time.Time)
	defer close(dob)
	go func(dob chan time.Time) {
		fmt.Println("Goroutine three initialized!")
		dur := time.Millisecond * time.Duration(rand.Intn(2500))
		time.Sleep(dur)
		dob <- time.Now()
		fmt.Println("Goroutine three done!", dur)
	}(dob)

	fmt.Println("name", <-name)
	fmt.Println("age", <-age)
	fmt.Println("dob", <-dob)
}
