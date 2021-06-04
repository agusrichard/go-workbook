package main

import (
	"fmt"
	"testing"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test")
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}

func TestCalculate(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{ 1, 3},
		{ 2, 4},
		{ 3, 5},
		{ 4, 6},
		{ 5, 7},
	}

	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range cases {
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
