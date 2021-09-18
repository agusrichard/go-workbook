package main

import "testing"

func BenchmarkCalculate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(1, 2)
	}
}
