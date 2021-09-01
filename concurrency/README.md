# Concurrency in Go

</br>

## List of Contents:
### 1. [Concurrency](#content-1)


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
- This program consists of two goroutines. The first goroutine is implicit and is the main function itself. The second goroutine is created when we call go f(0).
- Normally when we invoke a function our program will execute all the statements in a function and then return to the next line following the invocation.
- With a goroutine we return immediately to the next line and don't wait for the function to complete.
- This is why the call to the Scanln function has been included; without it the program would exit before being given the opportunity to print all the numbers.
- oroutines are lightweight and we can easily create thousands of them.
- Example of create a bunch of goroutines:
  ```go
  func main() {
    for i := 0; i < 10; i++ {
      go f(i)
    }
    var input string
    fmt.Scanln(&input)
  }
  ```



**[⬆ back to top](#list-of-contents)**

</br>

---


## References:
- https://www.golang-book.com/books/intro/10