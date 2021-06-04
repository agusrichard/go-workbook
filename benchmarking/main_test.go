package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	expected := 4
	result := Calculate(2)
	fmt.Println("result", result)
	if expected != result {
		t.Error(fmt.Sprintf("expect %v but got %v", expected, result))
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(1)
	}
}

func TestOther(t *testing.T) {
	fmt.Println("Testing something else")
	fmt.Println("This shouldn't run with -run=calc")
}
