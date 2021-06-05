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
	UselessMocked func() int
}

func (m *MyIntMocked) Useless() int {
	return m.UselessMocked()
}

func TestUsingUseless_Mocked(t *testing.T) {
	m := MyIntMocked{}
	m.UselessMocked = func() int {
		return 1
	}

	expected := 1
	actual := UsingUseless(&m)
	if actual != expected {
		t.Fatalf("expect %v got %v", expected, actual)
	}
}