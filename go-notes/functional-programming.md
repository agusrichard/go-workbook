# Functional Programming in Go

</br>

## List of Contents:
## 1. [Functional programming in Go](#content-1)
## 2. [Functional Go](#content-2)

</br>

---

## Contents:

## [Functional programming in Go](https://blog.logrocket.com/functional-programming-in-go/) <span id="content-1"></span>

Code from learning session: https://github.com/agusrichard/go-workbook/tree/master/functional-programming/content-1

### Introduction

Software engineer and author Eric Elliot defined function programming as follows.

> Functional programming is the process of building software by composing pure functions, avoiding shared state, mutable data, and side-effects. Functional programming is declarative rather than imperative, and application state flows through pure functions. Contrast with object-oriented programming, where application state is usually shared and colocated with methods in objects.


### 4 important concepts to understand
- Pure functions and idempotence
  - A pure function always returns the same output if you give it the same input (Idempotence property).
  - Idempotence means that a function should always return the same output, independent of the number of calls.
- Side effects
  - A pure function can’t have any side effects.
  - For example, functional programming considers an API call to be a side effect. Why? Because an API call is considered an external environment that is not under your direct control.
  - Other common side effects include:
    - Data mutation
    - DOM manipulation
    - Requesting conflicting data
- Function composition
  - The basic idea of function composition is straightforward: you combine two pure functions to create a new function.
- Shared state and immutable data
  - The goal of functional programming is to create functions that do not hold a state.
  - Shared states, especially, can introduce side effects or mutability problems in your pure functions, rendering them nonpure.
  - The goal of functional programming is to make the state visible and explicit to eliminate any side effects.
  - A program uses immutable data structures to derive new data from using pure functions.
  - This way, there is no need for mutable data that may cause side effects.

### Rules for functional programming
- Guidelines to follow the functional programming paradigms as closely as possible:
  - No mutable data to avoid side effects
  - No state (or implicit state, such as a loop counter)
  - Do not modify variables once they are assigned a value
  - Avoid side effects, such as an API call
- One good “side effect” we often encounter in functional programming is strong modularization.
- Start by defining modules that group similar pure functions that you expect to need in the future.
- Next, start writing those small, stateless, independent functions to create your first modules.
- We are essentially creating black boxes.
- This enables you to build a strong base of tests, especially unit tests that verify the correctness of your pure functions.
- This step in the development process also involves writing integration tests to ensure proper integration of the two components.

### 5 Functional programming examples in Go

#### 1. Updating a string
```go
// Bad
name := "first name"
name := name + " last name"

// Good
firstname := "first"
lastname := "last"
fullname := firstname + " " + lastname
```
- A variable can't be modified within a function.

#### 2. Avoid updating arrays
- The objective of functional programming is to use immutable data to derive a new immutable data state through pure functions.
- Comparison: </br>
  ```go
  // Non functional programming
  names := [3]string{"Tom", "Ben"}
  // Add Lucas to the array
  names[2] = "Lucas"

  // Functional programming
  names := []string{"Tom", "Ben"}
  allNames := append(names, "Lucas")
  ```

#### 3. Avoid updating maps
- Comparison: </br>
  ```go
  // Nonfunctional programming
  fruits := map[string]int{"bananas": 11}
  // Buy five apples
  fruits["apples"] = 5

  // Functional programming
  fruits := map[string]int{"bananas": 11}
  newFruits := map[string]int{"apples": 5}

  allFruits := make(map[string]int, len(fruits) + len(newFruits))


  for k, v := range fruits {
      allFruits[k] = v
  }


  for k, v := range newFruits {
      allFruits[k] = v
  }
  ```
- By using functional programming, we have longer code.
- The performance of this snippet of is much worse than a simple mutable update of the map because we are looping through both maps.


### 4. Higher-order functions and currying
- Higher order functions would be handy to establish currying.
- Example: </br>
  ```go
  func add(x int) func(y int) int {
      return func(y int) int {
          return x + y
      }
  }

  func main() {
      // Create more variations
      add10 := add(10)
      add20 := add(20)

      // Currying
      fmt.Println(add10(1)) // 11
      fmt.Println(add20(1)) // 21
  }
  ```

### 5. Recursion
- Recursion is a software pattern that is commonly employed to circumvent the use of loops. Because loops always hold an internal state to know which round they’re at, we can’t use them under the functional programming paradigm.
- Comparison: </br>
  ```go
  // Non functional programming
  func factorial(fac int) int {
      result := 1
      for ; fac > 0; fac-- {
          result *= fac
      }
      return result
  }

  // Functional programming
  func calculateFactorial(fac int) int {
      if fac == 0 {
          return 1
      }
      return fac * calculateFactorial(fac - 1)
  }
  ```


### Conculsion
- Although Golang supports functional programming, it wasn’t designed for this purpose.
- Functional programming improves the readability of your code because functions are pure.
- Pure functions are easier to test since there is no internal state that can alter the output.

</br>

---

## [Functional Go](https://medium.com/@geisonfgfg/functional-go-bc116f4c96a4) <span id="content-2"></span>

Code from learning session: https://github.com/agusrichard/go-workbook/tree/master/functional-programming/content-2

### Introduction
- Functional Programming by Wikipedia:
  > “Functional programming is a programming paradigm that treats computation as the evaluation of mathematical functions and avoids state and mutable data”. In other words, functional programming promotes code with no side effects, no change of value in variables. It opposes imperative programming, which empathizes change of state”
- It means no mutable data and no state (implicit, hidden state)
- Once assigned, a variable does not change
- Functions are pure functions in the mathematical sense: their output depend only on their inputs, there is no “environment”
- The same result returned by functions called with the same inputs

### The advantages
- **Cleaner code**: “variables” are not modified once defined, so we don’t have to follow the change of state to comprehend what a function, a, method, a class, a whole project works.
- **Referential transparency**: Expressions can be replaced by their values. If we call a function with the same parameters, we know for sure the output will be the same (there is no state anywhere that would change it).


### Advantages enabled by referential transparency
- Memoization: Cache results for previous function calls
- Idempotence: Same results regardless of how many times you call a function
- Modularization: We build our project bottom- up. From black boxes then combine them into a whole project
- Ease of debugging: Functions are isolated, they only depend on their input and their output, so they are very easy to debug.
- Parallelization: Functions' calls are independent. So we can parallelize them. `result = func1(a, b) + func2(a, c)`, here we can parallelize func1 and func2.
- Concurrence:
  - With no shared data, concurrency gets a lot simpler. No locks, no race conditions, and no dead-locks.

### Examples:
- Don’t Update, Create — String </br>
  ```go
  // Non-functional programming
  name := “Geison”
  name := name + “ Flores”

  // Functional programming
  const firstname = “Geison”
  const lasname = “Flores”
  const name = firstname + “ “ + lastname
  ```
- Don’t Update, Create — Arrays </br>
  ```go
  // Non-functional programming
  years := [4]int{2001, 2002}
  years[2] = 2003
  years[3] = 2004
  years // [2001, 2002, 2003, 2004]

  // Functional programming
  years := [2]int{2001, 2002}
  allYears := append(years, 2003, [2]int{2004, 2005})
  ```
- Don’t Update, Create — Maps </br>
  ```go
  // Non-functional programming
  ages := map[string]int{“John”: 30}
  ages[“Mary”] = 28
  ages // {‘John’: 30, ‘Mary’: 28}

  // Functional programming
  ages1 := map[string]int{“John”: 30}
  ages2 := map[string]int{“Mary”: 28}
  func mergeMaps(mapA, mapB map[string]int) map[string]int {
  allAges := make(map[string]int, len(mapA) + len(mapB))
      for k, v := range mapA {
          allAges[k] = v
      }
      for k, v := range mapB {
          allAges[k] = v
      }
      return allAges
  }
  allAges := mergeMaps(ages1, ages2)
  ```
- Higher order functions </br>
  ```go
  func caller(f func(string) string) {
      result := f(“David”)
      fmt.Println(result)
  }
  f := func(s name) string {
      return “Hello “ + name
  }
  caller(f)
  ```
- Closure </br>
  ```go
  func add_x(x int) func(z int) int {
    return func(y int) int { // anonymous function
        return x + y
    }
  }
  add_5 := add_x(5)
  add_7 := add_x(7)
  add_5(10) // result 15
  add_7(10) // result 17
  ```

### Currying and Partial Functions
- Higher-order functions enable Currying, which the ability to take a function that accepts n parameters and turns it into a composition of n functions each of them takes 1 parameter.
- Snippet: </br>
  ```go
  func plus(x, y int) int {
      return x + y
  }
  func partialPlus(x int) func(int) int {
      return func(y int) int {
          return plus(x, y)
      }
  }
  func main() {
      plus_one := partialPlus(1)
      fmt.Println(plus_one(5)) //prints 6
  }
  ```

### Eager vs Lazy Evaluation
- Eager evaluation: expressions are calculated at the moment that variables are assigned to the function called
- Lazy evaluation: delays the evaluation of the expression until it is needed.
- Memory efficient: no memory used to store complete structures
- CPU efficient: no need to calculate the complete result before returning
- Laziness is not a requisite for FP, but it is a strategy that fits nicely on the paradigm(Haskell)
- Golang uses eager evaluation
- Golang arrays are not lazy, use channels and goroutines to emulate a generator when necessary


### Recursion
- Looping is not a functional concept, as it requires variables to be passed around to store the state of the loop at a given time
- Need to be remembered: a recursion function could involve expensive operations
- Purely functional languages have no imperative for-loops, so they use recursion a lot.
- If every recursion created a stack, it would blow up very soon
- Tail-call optimization (TCO) avoids creating a new stack when the last call in recursion is the function itself

### FP in OOP?
- Yes, it's possible.
- OOP is orthogonal to FP
- Typical OOP tends to emphasize the change of state in objects
- Typical OOP mixes the concepts of identity and state

### A Practical Example
- Imperative: </br>
  ```go
  func main() {
      n, numElements, s := 1, 0, 0
      for numElements < 10 {
          if n * n % 5 == 0 {
              s += n
              numElements++
          }
          n++
      }
      fmt.Println(s) //275
  }
  ```
- Functional: </br>
  ```go
  sum := func (memo interface{}, el interface{}) interface{} {
      return memo.(float64) + el.(float64)
  }
  pred := func (i interface{}) bool {
      return (i.(uint64) * i.(uint64)) % 5 == 0
  }
  createValues := func () []int {
      values := make([]int, 100)
      for num := 1; num <= 100; num++ {
          values = append(values, num)
      }
      return values
  }
  Reduce(Filter(pred, createValues), sum, uint64).(uint64)
  ```


</br>

---

## References:
- https://blog.logrocket.com/functional-programming-in-go/
- https://medium.com/@geisonfgfg/functional-go-bc116f4c96a4