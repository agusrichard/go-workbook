package main

import (
	"fmt"
	"testing"
)

var result *interface{}
var i int

func setupTestCase(t testing.TB) func(t testing.TB) {
	t.Log("setup test case")
	return func(t testing.TB) {
		t.Log("teardown test case")
	}
}

func setupSubTest(t testing.TB) func(t testing.TB) {
	t.Log("setup sub test")
	return func(t testing.TB) {
		t.Log("teardown sub test")
	}
}

func TestCalculate(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{1, 3},
		{2, 4},
		{3, 5},
		{4, 6},
		{5, 7},
	}

	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range cases {
		fmt.Println("here")
		t.Run(fmt.Sprintf("running test when n=%d", tc.n), func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)

			actual := Calculate(tc.n)
			if actual != tc.expected {
				t.Fatalf("expect %v got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkCalculate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(1)
	}
}
