package main

import (
	"fmt"
	"log"
	"testing"
)

func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Println("setup test")

	return func(tb testing.TB) {
		log.Println("teardown test")
	}
}

func setupSubTest(tb testing.TB) func(tb testing.TB) {
	log.Println("setup subtest")

	return func(tb testing.TB) {
		log.Println("teardown subtest")
	}
}

func TestMyInt_Useless(t *testing.T) {
	var cases []struct{name string; input, expected int}

	teardownTest := setupTest(t)
	defer teardownTest(t)

	for i := 0; i < 10; i++ {
		cases = append(cases, struct{name string; input, expected int}{
			name: fmt.Sprintf("case_%d", i),
			input: i,
			expected: i,
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)

			m := MyInt{number: tc.input}

			actual := m.Useless()
			if actual != tc.expected {
				t.Fatalf("expect %v got %v", tc.expected, actual)
			}
		})
	}
}

func TestUsingUseless_One(t *testing.T) {
	var cases []struct{name string; input, expected int}

	teardownTest := setupTest(t)
	defer teardownTest(t)

	for i := 0; i < 10; i++ {
		cases = append(cases, struct{name string; input, expected int}{
			name: fmt.Sprintf("case_%d", i),
			input: i,
			expected: i,
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)

			m := MyInt{number: tc.input}

			actual := UsingUseless(&m)
			if actual != tc.expected {
				t.Fatalf("expect %v got %v", tc.expected, actual)
			}
		})
	}
}

type MyIntMocked struct {
	number int
	Useless func(n int) int
}


func TestUsingUseless_Mocked(t *testing.T) {
	m := MyIntMocked{number: 1}
	m.Useless = func(n int) int {
		return m.number
	}

	expected, input := 1, 1
	actual := m.Useless(input)
	if actual != expected {
		t.Fatalf("expect %v got %v", expected, actual)
	}
}