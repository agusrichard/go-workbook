package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Calculate(x, y int64) int64 {
	// get random integer between 0 and 2000
	n := rand.Intn(2000)
	// sleep for a random duration between 0 and 2000 milliseconds
	time.Sleep(time.Duration(n) * time.Millisecond)

	return x + y
}

func CalculateSlow(x, y int64) int64 {
	// sleep for 3 seconds to mock heavy calculation
	time.Sleep(2 * time.Second)

	return x + y
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the query string from the request
	queryX, queryY := r.URL.Query().Get("x"), r.URL.Query().Get("y")

	// Parse the input from user to int or returns error if it fails
	x, err := strconv.ParseInt(queryX, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the input from user to int or returns error if it fails
	y, err := strconv.ParseInt(queryY, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := make(chan int64, 1)
	go func(r chan<- int64) {
		r <- Calculate(x, y)
	}(result)

	resultSlow := make(chan int64, 1)
	go func(r chan<- int64) {
		r <- CalculateSlow(x, y)
	}(resultSlow)

	fmt.Fprintf(w, "%d", <-result+<-resultSlow)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
