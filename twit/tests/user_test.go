package tests

import (
	"testing"
	"twit/usecases"
)

func TestAddition(t *testing.T) {
	var tests = []struct{ x, y, expected int }{
		{1, 1, 2},
		{2, 2, 4},
		{-1, -1, -2},
		{0, 1, 1},
		{10000, 2000, 12000},
	}
	for _, test := range tests {
		if output := usecases.Addition(test.x, test.y); output != test.expected {
			t.Error("Test Failed: {} x, {} y, {} expected, recieved: {}", test.x, test.y, test.expected, output)
		}
	}
}
