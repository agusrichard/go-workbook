package main

import (
	"fmt"
	"sync"
)

// var (
// 	mutex   sync.Mutex
// 	balance int
// )

// func init() {
// 	balance = 1000
// }

// func deposit(value int, wg *sync.WaitGroup) {
// 	mutex.Lock()
// 	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
// 	balance += value
// 	mutex.Unlock()
// 	wg.Done()
// }

// func withdraw(value int, wg *sync.WaitGroup) {
// 	mutex.Lock()
// 	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
// 	balance -= value
// 	mutex.Unlock()
// 	wg.Done()
// }

// func main() {
// 	fmt.Println("Go Mutex Example")

// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go withdraw(700, &wg)
// 	fmt.Println("here1")
// 	go deposit(500, &wg)
// 	fmt.Println("here2")
// 	wg.Wait()

// 	fmt.Printf("New Balance %d\n", balance)
// }

func appendToSlice(s *[]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	fmt.Println("Appending the slice")
	*s = append(*s, 10)
	mu.Unlock()
	wg.Done()
}

func popTheSlice(s *[]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	fmt.Println("Popping the slice")
	*s = (*s)[:len(*s)-1]
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	s := []int{1, 2, 3, 4, 5}
	wg.Add(2)
	go appendToSlice(&s, &wg, &mu)
	go popTheSlice(&s, &wg, &mu)
	fmt.Println("I don't know what I am doing in here")
	fmt.Println("Yeah... Me too!")

	wg.Wait()
	fmt.Println("s", s)
}
