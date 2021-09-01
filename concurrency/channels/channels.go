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
func CalculateValue(c chan int) {
	value := rand.Intn(10)
	fmt.Printf("Calculated random value: %d\n", value)
	time.Sleep(time.Second)
	c <- value
	fmt.Println("Only executes after another goroutine performs a receive on the channel", value)
}

func main() {
	fmt.Println("Go channel tutorial")

	valueChannel := make(chan int)
	defer close(valueChannel)

	go CalculateValue(valueChannel)
	go CalculateValue(valueChannel)

	values := <-valueChannel
	fmt.Println(values)
}

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
