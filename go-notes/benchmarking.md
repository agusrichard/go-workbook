# Benchmarking in Go

</br>

## List of Contents:
### 1. [An Introduction to Benchmarking Your Go Programs](#content-1)
### 2. [How to write benchmarks in Go](#content-2)


</br>

---

## Contents:

## [An Introduction to Benchmarking Your Go Programs](https://tutorialedge.net/golang/benchmarking-your-go-programs/) <span id="content-1"></span>

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

</br>

---

## [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go) <span id="content-2"></span>

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


</br>

---


## References:
- https://tutorialedge.net/golang/benchmarking-your-go-programs/
- https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go