# Concurrency in Go

</br>

## List of Contents:
### 1. [Concurrency](#content-1)
### 2. [Concurrency With Golang Goroutines](#content-2)


</br>

---

## Contents:

## [An Introduction to Benchmarking Your Go Programs](https://www.golang-book.com/books/intro/10) <span id="content-1"></span>

### Introduction
- Making progress on more than one task simultaneously is known as concurrency.

### Goroutines
- A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine we use the keyword go followed by a function invocation
- Example:
  ```go
  package main

  import "fmt"

  func f(n int) {
    for i := 0; i < 10; i++ {
      fmt.Println(n, ":", i)
    }
  }

  func main() {
    go f(0)
    var input string
    fmt.Scanln(&input)
  }
  ```


**[⬆ back to top](#list-of-contents)**

</br>

---


## References:
- https://www.golang-book.com/books/intro/10