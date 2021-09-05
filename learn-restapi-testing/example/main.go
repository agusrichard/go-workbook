package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Message string  `json:"message"`
	Result  float64 `json:"result"`
}

// Main function to be tested
func doubleMe(x float64) float64 {
	return x * 2
}

// Main handler function that receives the request and returns the response
func handler(w http.ResponseWriter, r *http.Request) {

	// Get the query string from the request
	queryX := r.URL.Query().Get("x")

	// Parse the input from user to float and returns error if it fails
	x, err := strconv.ParseFloat(queryX, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send back the appropriate response
	resp := Response{
		Message: "Hello!",
		Result:  doubleMe(x),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
