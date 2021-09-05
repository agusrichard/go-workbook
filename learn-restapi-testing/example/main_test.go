package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// You can use testing.T, if you want to test the code without benchmarking
func setupSuite(tb testing.TB) func(tb testing.TB) {
	log.Println("setup suite")

	// Return a function to teardown the test
	return func(tb testing.TB) {
		log.Println("teardown suite")
	}
}

// Almost the same as the above, but this one is for single test instead of collection of tests
func setupTest(tb testing.TB, x float64) (func(tb testing.TB), *httptest.ResponseRecorder) {
	log.Println("setup test")

	// Http method is GET and path is /?x=1 for example
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		tb.Fatal(err)
	}

	// Add query parameter x which is the input for doubleMe function
	q := req.URL.Query()
	q.Add("x", fmt.Sprint(x))
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)
	h.ServeHTTP(rr, req)

	// Return tb and response recorder
	return func(tb testing.TB) {
		log.Println("teardown test")
	}, rr
}

func TestDoubleMe(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	table := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one", 1, 2},
		{"minus one", -1, -2},
		{"zero", 0, 0},
		{"minus one hundred", -100, -200},
		{"one hundred", 100, 200},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest, rr := setupTest(t, tc.input)
			defer teardownTest(t)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			actual := Response{}
			json.Unmarshal([]byte(rr.Body.String()), &actual)
			if actual.Result != tc.expected {
				t.Errorf("expected %f, got %f", tc.expected, actual.Result)
			}
		})
	}
}
