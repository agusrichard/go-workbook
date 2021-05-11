package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		x, y   float64
		output float64
	}{
		{1, 2, 3},
		{-1, 2, 1},
		{1.51, 2.31, 3.82},
		{0.1, 1.1, 1.2},
		{0, 0, 0},
	}

	for _, c := range testCases {
		assert.InDelta(t, Add(c.x, c.y), c.output, 0.0000001)
	}
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		x, y   float64
		output float64
	}{
		{3, 2, 1},
		{-1, 2, -3},
		{1.51, 2.31, -0.8},
		{0.1, 1.1, -1},
		{0, 0, 0},
	}

	for _, c := range testCases {
		assert.InDelta(t, Subtract(c.x, c.y), c.output, 0.0000001)
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		x, y   float64
		output float64
	}{
		{3, 2, 6},
		{-1, 2, -2},
		{1.51, 2.31, 3.4881},
		{0.1, 1.1, 0.11},
		{0, 0, 0},
	}

	for _, c := range testCases {
		assert.InDelta(t, Multiply(c.x, c.y), c.output, 0.0000001)
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		x, y   float64
		output float64
	}{
		{3, 2, 1.5},
		{-1, 2, -0.5},
		{1.51, 2.31, 0.65367965368},
		{0.1, 1.1, 0.09090909091},
		{1, 1, 1},
	}

	for _, c := range testCases {
		assert.InDelta(t, Divide(c.x, c.y), c.output, 0.0000001)
	}
}
