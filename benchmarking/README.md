# Benchmarking in Go

## [An Introduction to Benchmarking Your Go Programs](https://tutorialedge.net/golang/benchmarking-your-go-programs/)

> Note - It’s important to note that performance tweaking should typically be done after the system is up and running.

> "Premature optimization is the root of all evil” - Donald Knuth

```go
// main.go
package main

import "fmt"

func Calculate(x int) int {
	result := x+2
	return result
}

func main() {
	fmt.Println("Hello World")
}
```

```go
// main_test.go
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
```

To run just benchmark case then what we need to do just `go test -run=Bench -bench=.`.
Therefore, to define a benchmark function we need to star the function definition with `Benchmark`.

## [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

For better or worse, modern CPUs rely heavily on active thermal management which can add noise to benchmark results.

- Benchmark functions start with Benchmark not Test.
- Benchmark functions are run several times by the testing package. The value of b.N will increase each time until the benchmark runner is satisfied with the stability of the benchmark. This has some important ramifications which we’ll investigate later in this article.
- Each benchmark must execute the code under test b.N times. The for loop in BenchmarkFib10 will be present in every benchmark function.

`go test -bench=.` will run benchmarking and testing
`go test -run=xxx -bench=.` will run benchmarking only (if our tests' definitions don't contain xxx)

There are some other things to observe in this benchmark run.

- Each benchmark is run for a minimum of 1 second by default. If the second has not elapsed when the Benchmark function returns, the value of b.N is increased in the sequence 1, 2, 5, 10, 20, 50, … and the function run again.
- The final BenchmarkFib40 only ran two times with the average was just under a second for each run. As the testing package uses a simple average (total time to run the benchmark function over b.N) this result is statistically weak. You can increase the minimum benchmark time using the -benchtime flag to produce a more accurate result.
```shell
go test -bench=Fib40 -benchtime=20s
```

### Traps for young player

Here are two examples of a faulty Fib benchmark.
```go
func BenchmarkFibWrong(b *testing.B) {
        for n := 0; n < b.N; n++ {
                Fib(n)
        }
}

func BenchmarkFibWrong2(b *testing.B) {
        Fib(b.N)
}
```

### A note on compiler optimisatioons
Before concluding I wanted to highlight that to be completely accurate, any benchmark should be careful to avoid compiler optimisations eliminating the function under test and artificially lowering the run time of the benchmark.

```go
var result int

func BenchmarkFibComplete(b *testing.B) {
        var r int
        for n := 0; n < b.N; n++ {
                // always record the result of Fib to prevent
                // the compiler eliminating the function call.
                r = Fib(10)
        }
        // always store the result to a package level variable
        // so the compiler cannot eliminate the Benchmark itself.
        result = r
}
```

Example:
```go
func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
~~~~
var result int

func BenchmarkFibComplete(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Fib(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
```

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

### [Another patterns of testing in Go](https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go)

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



## References:
- https://tutorialedge.net/golang/benchmarking-your-go-programs/
- https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
- https://golang.org/pkg/testing/
- https://riptutorial.com/go/example/15183/testing-using-setup-and-teardown-function
- https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go
- https://kenanbek.medium.com/go-tests-with-html-coverage-report-f977da09552d