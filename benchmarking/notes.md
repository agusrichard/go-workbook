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



## References:
- https://tutorialedge.net/golang/benchmarking-your-go-programs/
- https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go