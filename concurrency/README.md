# Concurrency in Go

</br>

## List of Contents:
### 1. [Concurrency](#content-1)
### 2. [Concurrency With Golang Goroutines](#content-2)
### 3. [Go Channels Tutorial](#content-3)


</br>

---

## Contents:

## [Concurrency](https://www.golang-book.com/books/intro/10) <span id="content-1"></span>

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

### Channels
- Channels provide a way for two goroutines to communicate with one another and synchronize their execution.
- Example:
  ```go
  package main

  import (
    "fmt"
    "time"
  )

  func pinger(c chan string) {
    for i := 0; ; i++ {
      c <- "ping"
    }
  }

  func printer(c chan string) {
    for {
      msg := <- c
      fmt.Println(msg)
      time.Sleep(time.Second * 1)
    }
  }

  func main() {
    var c chan string = make(chan string)

    go pinger(c)
    go printer(c)

    var input string
    fmt.Scanln(&input)
  }
  ```
- This program will print “ping” forever (hit enter to stop it).
- A channel type is represented with the keyword chan followed by the type of the things that are passed on the channel (in this case we are passing strings). 
- The `<-` (left arrow) operator is used to send and receive messages on the channel. c `<-` "ping" means send "ping". `msg := <-` c means receive a message and store it in msg.
- The fmt line could also have been written like this: `fmt.Println(<-c)` in which case we could remove the previous line.
- Using a channel like this synchronizes the two goroutines. When pinger attempts to send a message on the channel it will wait until printer is ready to receive the message. (this is known as blocking)


### Channel Direction
- We can specify a direction on a channel type thus restricting it to either sending or receiving.
- For example pinger's function signature can be changed to this:
  ```go
  func pinger(c chan<- string)
  ```
- Similarly we can change printer to this:
  ```go
  func printer(c <-chan string)
  ```
- A channel that doesn't have these restrictions is known as bi-directional. A bi-directional channel can be passed to a function that takes send-only or receive-only channels, but the reverse is not true.


### Select
- Example:
  ```go
  func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
      for {
        c1 <- "from 1"
        time.Sleep(time.Second * 2)
      }
    }()

    go func() {
      for {
        c2 <- "from 2"
        time.Sleep(time.Second * 3)
      }
    }()

    go func() {
      for {
        select {
        case msg1 := <- c1:
          fmt.Println(msg1)
        case msg2 := <- c2:
          fmt.Println(msg2)
        }
      }
    }()

    var input string
    fmt.Scanln(&input)
  }
  ```
- This program prints “from 1” every 2 seconds and “from 2” every 3 seconds.
- `select` picks the first channel that is ready and receives from it (or sends to it). If more than one of the channels are ready then it randomly picks which one to receive from. If none of the channels are ready, the statement blocks until one becomes available.
- Example:
  ```go
  select {
  case msg1 := <- c1:
    fmt.Println("Message 1", msg1)
  case msg2 := <- c2:
    fmt.Println("Message 2", msg2)
  case <- time.After(time.Second):
    fmt.Println("timeout")
  }
  ```
- `time.After` creates a channel and after the given duration will send the current time on it. (we weren't interested in the time so we didn't store it in a variable)
- The default case happens immediately if none of the channels are ready.
- Buffered Channels Example:
  ```go
  c := make(chan int, 1)
  ```
- This creates a buffered channel with a capacity of 1. Normally channels are synchronous; both sides of the channel will wait until the other side is ready. A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Concurrency With Golang Goroutines](https://tutorialedge.net/golang/concurrency-with-golang-goroutines/) <span id="content-1"></span>


### What Are Goroutines?
- Goroutines are incredibly lightweight “threads” managed by the go runtime. They enable us to create asynchronous parallel programs that can execute some tasks far quicker than if they were written in a sequential manner.
- Goroutines are far smaller that threads, they typically take around 2kB of stack space to initialize compared to a thread which takes 1Mb.
- Creating a thousand goroutines would typically require one or two OS threads at most, whereas if we were to do the same thing in java it would require 1,000 full threads each taking a minimum of 1Mb of Heap space.
- It’s incredibly in-expensive to create and destroy new goroutines due to their size and the efficient way that go handles them.


### A Simple Sequential Program
- Example:
  ```go
  package main


  import (
      "fmt"
      "time"
  )


  // a very simple function that we'll
  // make asynchronous later on
  func compute(value int) {
      for i := 0; i < value; i++ {
          time.Sleep(time.Second)
          fmt.Println(i)
      }
  }

  func main() {
      fmt.Println("Goroutine Tutorial")

      // sequential execution of our compute function
      compute(10)
      compute(10)

      // we scan fmt for input and print that to our console
      var input string
      fmt.Scanln(&input)

  }
  ``` 

### Making our Program Asynchronous
- If we aren’t fussed about the order in which our program prints out the values 0 to n then we can speed this program up by using goroutines and making it asynchronous.
- Example:
  ```go
  package main


  import (
      "fmt"
      "time"
  )

  // notice we've not changed anything in this function
  // when compared to our previous sequential program
  func compute(value int) {
      for i := 0; i < value; i++ {
          time.Sleep(time.Second)
          fmt.Println(i)
      }
  }

  func main() {
      fmt.Println("Goroutine Tutorial")

      // notice how we've added the 'go' keyword
      // in front of both our compute function calls
      go compute(10)
      go compute(10)

      var input string
      fmt.Scanln(&input)
  }
  ```

### Anonymous Goroutine Functions
- Example:
  ```go
  package main

  import "fmt"

  func main() {
      // we make our anonymous function concurrent using `go`
      go func() {
          fmt.Println("Executing my Concurrent anonymous function")
      }()
      // we have to once again block until our anonymous goroutine
      // has finished or our main() function will complete without
      // printing anything
      fmt.Scanln()
  }
  ```


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Go Channels Tutorial](https://tutorialedge.net/golang/go-channels-tutorial/) <span id="content-3"></span>

### Introduction
- Channels are pipes that link between goroutines within your Go based applications that allow communication and subsequently the passing of values to and from variables.

### A Simple Example
- Example:
  ```go
  package main

  import (
      "fmt"
      "math/rand"
  )

  func CalculateValue(values chan int) {
      value := rand.Intn(10)
      fmt.Println("Calculated Random Value: {}", value)
      values <- value
  }

  func main() {
      fmt.Println("Go Channel Tutorial")

      values := make(chan int)
      defer close(values)

      go CalculateValue(values)

      value := <-values
      fmt.Println(value)
  }
  ```
- In our `main()` function, we called `values := make(chan int)`, this call effectively created our new channel so that we could subsequently use it within our CalculateValue goroutine.
- After we created out channel, we then called defer close(values) which deferred the closing of our channel until the end of our main() function’s execution. This is typically considered best practice to ensure that we tidy up after ourselves.
- After our call to defer, we go on to kick off our single goroutine: `CalculateValue(values)` passing in our newly created values channel as its parameter. Within our `CalculateValue` function, we calculate a single random value between 1-10, print this out and then send this value to our values channel by calling `values <- value`.
- Jumping back into our main() function, we then call value := <-values which receives a value from our values channel.
- Notice how when we execute this program, it doesn’t immediately terminate. This is because the act of sending to and receiving from a channel are blocking. Our `main()` function blocks until it receives a value from our channel.
- With traditional unbuffered channels, whenever one goroutine sends a value to this channel, that goroutine will subsequently block until the value is received from the channel.


### Unbuffered Channels
- Example:
  ```go
  package main

  import (
      "fmt"
      "math/rand"
      "time"
  )

  func CalculateValue(c chan int) {
      value := rand.Intn(10)
      fmt.Println("Calculated Random Value: {}", value)
      time.Sleep(1000 * time.Millisecond)
      c <- value
      fmt.Println("Only Executes after another goroutine performs a receive on the channel")
  }

  func main() {
      fmt.Println("Go Channel Tutorial")

      valueChannel := make(chan int)
      defer close(valueChannel)

      go CalculateValue(valueChannel)
      go CalculateValue(valueChannel)

      values := <-valueChannel
      fmt.Println(values)
  }
  ```

### Buffered Channels
- These buffered channels are essentially queues of a given size that can be used for cross-goroutine communication. In order to create a buffered channel as opposed to an unbuffered channel, we supply a capacity argument to our make command:
  ```go
  bufferedChannel := make(chan int, 3)
  ```
- By changing this to a buffered channel, our send operation, c <- value only blocks within our goroutines should the channel be full.
- 

## References:
- https://www.golang-book.com/books/intro/10
- https://tutorialedge.net/golang/concurrency-with-golang-goroutines/
- https://tutorialedge.net/golang/go-channels-tutorial/