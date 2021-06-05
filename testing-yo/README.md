# Benchmarking in Go

## [Setup and Teardown using Go testing package](https://golang.org/pkg/testing/)

You can set a setUp and tearDown function.

A setUp function prepares your environment to tests.
A tearDown function does a rollback.

```go
package main

import (
	"fmt"
	"os"
	"testing"
)

var number int

func setup() {
	number = 10
	fmt.Println("setup the test")
}

func teardown() {
	number = 99
	fmt.Println("teardown the test")
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func TestCalculate1(t *testing.T) {
	expected := 3
	actual := Calculate(1)
	if actual != expected {
		t.Error(fmt.Sprintf("expect %v got %v", expected, actual))
	}
}

func TestIntermediate(t *testing.T) {
	number += 10
	fmt.Println("here")
}

func TestCalculate2(t *testing.T) {
	fmt.Println("number", number)
	expected := 4
	actual := Calculate(2)
	if actual != expected {
		t.Error(fmt.Sprintf("expect %v got %v", expected, actual))
	}
}

func BenchmarkCalculate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(1)
	}
}
```

## [Another patterns of testing in Go](https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go)

We can use this approach when we want to implement table driven tests

```shell
go test -run ''      # Run all tests.
go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
```

```go
// main_test.go
// running parallel tests
func TestGroupedParallel(t *testing.T) {
    for _, tc := range tests {
    tc := tc // capture range variable
        t.Run(tc.Name, func(t *testing.T) {
            t.Parallel()
            ...
        })
    }
}
```

The race detector kills the program if it exceeds 8192 concurrent goroutines, so use care when running parallel tests with the -race flag set.

```go
// main_test.go
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

```

## [Go: tests with HTML coverage report](https://kenanbek.medium.com/go-tests-with-html-coverage-report-f977da09552d)

To add cover reports `go test -cover ./...`

To get detail information of coverage
```shell
go test -coverprofile=coverage.out ./...  # save coverage results
go tool cover -func=coverage.out          # print results
go tool cover -html=coverage.out          # view coverage as html page
```

The coverage tool also includes three different coverage modes. You can select coverage mode by using -covermode option:

```shell
go test -covermode=count -coverprofile=coverage.out
```

Makefile for simplicity

```makefile
GO=go
GOCOVER=$(GO) tool cover
.PHONY: test/cover
test/cover:
    $(GOTEST) -v -coverprofile=coverage.out ./...
    $(GOCOVER) -func=coverage.out
    $(GOCOVER) -html=coverage.out
```

There are three different cover modes:
- set: did each statement run?
- count: how many times did each statement run?
- atomic: like count, but counts precisely in parallel programs

## [How I mock unit tests in Golang](https://dev.to/chseki/how-i-mock-unit-tests-in-golang-3dcp)

```go
package usecases_test

import (
    "testing"

    "github.com/iamseki/dev-to/domain"
    "github.com/iamseki/dev-to/usecases"
)

type addEventFakeRepository struct {
    MockAddFn func(domain.Event) error
}

func (fake *addEventFakeRepository) Add(e domain.Event) error {
    return fake.MockAddFn(e)
}

func newAddEventFakeRepository() *addEventFakeRepository {
    return &addEventFakeRepository{
        MockAddFn: func(e domain.Event) error { return nil },
    }
}

func TestAddEventInMemorySucceed(t *testing.T) {
    r := newAddEventFakeRepository()
    sut := usecases.NewAddEventInMemory(r)

    err := sut.Save(domain.Event{})
    if err != nil {
        t.Error("Expect error to be nil but got:", err)
    }
}
```

## [How to mock? Go Way.](https://medium.com/@ankur_anand/how-to-mock-in-your-go-golang-tests-b9eee7d7c266)


## References:
- https://golang.org/pkg/testing/
- https://riptutorial.com/go/example/15183/testing-using-setup-and-teardown-function
- https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go
- https://kenanbek.medium.com/go-tests-with-html-coverage-report-f977da09552d