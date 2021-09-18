package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// package level variable to be used by the benchmarking code
// to prevent the compiler from optimizing the code away
var result int64

func benchmarkCalculateRestAPI(x, y int64, b *testing.B) int64 {
	// Http method is GET and path is /?x=1&y=1 for example
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		b.Fatal(err)
	}

	// Add query parameter x and y which is the input for Calculate and CalculateSlow function
	q := req.URL.Query()
	q.Add("x", fmt.Sprint(x))
	q.Add("y", fmt.Sprint(y))
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)
	h.ServeHTTP(rr, req)

	result, err := strconv.ParseInt(rr.Body.String(), 10, 64)
	if err != nil {
		b.Fatal(err)
	}

	return result
}

func BenchmarkCalculateRestAPI1000000(b *testing.B) {
	var r int64
	for n := 0; n < b.N; n++ {
		// store the result of the calculation in a variable
		// to prevent the compiler from optimizing the code away
		r = benchmarkCalculateRestAPI(1000000, 1000000, b)
	}

	// store the result into a package level variable
	result = r
	fmt.Println("result", result)
}

func BenchmarkCalculateRestAPI100(b *testing.B) {
	var r int64
	for n := 0; n < b.N; n++ {
		// store the result of the calculation in a variable
		// to prevent the compiler from optimizing the code away
		r = benchmarkCalculateRestAPI(100, 100, b)
	}

	// store the result into a package level variable
	result = r
}

func BenchmarkCalculateRestAPI1(b *testing.B) {
	var r int64
	for n := 0; n < b.N; n++ {
		// store the result of the calculation in a variable
		// to prevent the compiler from optimizing the code away
		r = benchmarkCalculateRestAPI(1, 1, b)
	}

	// store the result into a package level variable
	result = r
}
