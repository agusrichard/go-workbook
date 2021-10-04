# Concurrency in Go

</br>

## List of Contents:
### 1. [Concurrency](#content-1)
### 2. [Concurrency With Golang Goroutines](#content-2)
### 3. [Go Channels Tutorial](#content-3)
### 4. [Go WaitGroup Tutorial](#content-4)
### 5. [Go Mutex Tutorial](#content-5)
### 6. [Deep dive on goroutine leaks and best practices to avoid them](#content-6)


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


## [Go WaitGroup Tutorial](https://tutorialedge.net/golang/go-waitgroup-tutorial/) <span id="content-4"></span>

### Understanding WaitGroups
- When you start writing applications in Go that leverage goroutines, you start hitting scenarios where you need to block the execution of certain parts of your code base, until these goroutines have successfully executed.
- Example:
  ```go
  package main

  import "fmt"

  func myFunc() {
      fmt.Println("Inside my goroutine")
  }

  func main() {
      fmt.Println("Hello World")
      go myFunc()
      fmt.Println("Finished Execution")
  }
  ```
- The goroutine from the above code doesn't have a chance to run before the main function invocation terminates.
- `WaitGroups` essentially allow us to tackle this problem by blocking until any goroutines within that WaitGroup have successfully executed.
- We first call .Add(1) on our WaitGroup to set the number of goroutines we want to wait for, and subsequently, we call .Done() within any goroutine to signal the end of its' execution.

### A Simple Example
- Example:
  ```go
  package main

  import (
      "fmt"
      "sync"
  )

  func myFunc(waitgroup *sync.WaitGroup) {
      fmt.Println("Inside my goroutine")
      waitgroup.Done()
  }

  func main() {
      fmt.Println("Hello World")

      var waitgroup sync.WaitGroup
      waitgroup.Add(1)
      go myFunc(&waitgroup)
      waitgroup.Wait()

      fmt.Println("Finished Execution")
  }
  ```
- As you can see, we’ve instantiated a new sync.WaitGroup and then called the .Add(1) method, before attempting to execute our goroutine.
- We’ve updated the function to take in a pointer to our existing sync.WaitGroup and then called the .Done() method once we have successfully finished our task.
- Finally, on line 19, we call waitgroup.Wait() to block the execution of our main() function until the goroutines in the waitgroup have successfully completed.


### Anonymous Functions
- Example:
  ```go
  package main

  import (
      "fmt"
      "sync"
  )

  func main() {
      fmt.Println("Hello World")

      var waitgroup sync.WaitGroup
      waitgroup.Add(1)
      go func() {
          fmt.Println("Inside my goroutine")
          waitgroup.Done()
      }()
      waitgroup.Wait()

      fmt.Println("Finished Execution")
  }
  ```

### Real Example
- Example:
  ```go
  package main

  import (
      "fmt"
      "log"
      "net/http"
      "sync"
  )

  var urls = []string{
      "https://google.com",
      "https://tutorialedge.net",
      "https://twitter.com",
  }

  func fetch(url string, wg *sync.WaitGroup) (string, error) {
      resp, err := http.Get(url)
      if err != nil {
          fmt.Println(err)
          return "", err
      }
      wg.Done()
      fmt.Println(resp.Status)
      return resp.Status, nil
  }

  func homePage(w http.ResponseWriter, r *http.Request) {
      fmt.Println("HomePage Endpoint Hit")
      var wg sync.WaitGroup

      for _, url := range urls {
          wg.Add(1)
          go fetch(url, &wg)
      }

      wg.Wait()
      fmt.Println("Returning Response")
      fmt.Fprintf(w, "Responses")
  }

  func handleRequests() {
      http.HandleFunc("/", homePage)
      log.Fatal(http.ListenAndServe(":8081", nil))
  }

  func main() {
      handleRequests()
  }

  ```

### My Example
- One:
  ```go
  func main() {
  	fmt.Println("Working on WaitGroup")

  	var wg sync.WaitGroup
  	start := time.Now()
  	defer func() {
  		dur := time.Since(start)
  		fmt.Println("Dur", dur)
  	}()

  	for i := 0; i < 100; i++ {
  		wg.Add(1)
  		go func() {
  			dur := time.Duration(rand.Intn(1000)) * time.Millisecond
  			time.Sleep(dur)
  			fmt.Println("Hello", dur)
  			wg.Done()
  		}()
  	}

  	wg.Wait()
  	fmt.Println("Done on WaitGroup")
  }
  ```
- Two:
  ```go
  func main() {
  	fmt.Println("Working on WaitGroup")

  	var wg sync.WaitGroup
  	start := time.Now()
  	defer func() {
  		dur := time.Since(start)
  		fmt.Println("Dur", dur)
  	}()

  	num := make(chan int, 100)
  	for i := 0; i < 100; i++ {
  		wg.Add(1)
  		go func(i int) {
  			fmt.Println("Sending a num")
  			randInt := rand.Intn(1000)
  			dur := time.Duration(randInt) * time.Millisecond
  			time.Sleep(dur)

  			num <- i
  			fmt.Println("Done sending a num")
  		}(i)
  	}

  	var nums []int
  	go func() {
  		for n := range num {
  			wg.Done()
  			nums = append(nums, n)
  		}
  	}()

  	wg.Wait()

  	fmt.Println("nums", nums)
  	fmt.Println("Done on WaitGroup")
  }
  ```
- Three
  ```go
  func main() {
  	fmt.Println("Start working on WaitGroup")
  	start := time.Now()
  	defer func() {
  		fmt.Println("Done working on WaitGroup")
  		fmt.Println("Duration", time.Since(start))
  	}()

  	name := make(chan string)
  	defer close(name)
  	go func(name chan<- string) {
  		fmt.Println("Sending a name")
  		dur := time.Duration(rand.Intn(1000)) * time.Millisecond
  		time.Sleep(dur)
  		name <- "John"
  		fmt.Println("Done sending a name")
  	}(name)

  	age := make(chan int)
  	defer close(age)
  	go func(age chan<- int) {
  		fmt.Println("Sending a age")
  		dur := time.Duration(rand.Intn(1000)) * time.Millisecond
  		time.Sleep(dur)
  		age <- 23
  		fmt.Println("Done sending a age")
  	}(age)

  	nums := make(chan []int)
  	defer close(nums)
  	go func(nums chan<- []int) {
  		var list []int
  		var wg sync.WaitGroup

  		num := make(chan int, 100)
  		for i := 0; i < 100; i++ {
  			wg.Add(1)
  			go func(i int, wg *sync.WaitGroup) {
  				fmt.Println("Sending a num")
  				dur := time.Duration(rand.Intn(5000)) * time.Millisecond
  				time.Sleep(dur)
  				num <- i
  				fmt.Println("Done sending a num", i)
  			}(i, &wg)
  		}

  		go func(nums chan<- []int) {
  			for n := range num {
  				list = append(list, n)
  				wg.Done()
  			}
  		}(nums)

  		wg.Wait()

  		nums <- list
  	}(nums)

  	fmt.Println("name", <-name)
  	fmt.Println("age", <-age)
  	fmt.Println("nums", <-nums)
  }
  ```


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Go Mutex Tutorial](https://tutorialedge.net/golang/go-mutex-tutorial/) <span id="content-5"></span>

### The Theory
- A Mutex, or a mutual exclusion is a mechanism that allows us to prevent concurrent processes from entering a critical section of data whilst it’s already being executed by a given process.
- Let’s think about an example where we have a bank balance and a system that both deposits and withdraws sums of money from that bank balance. Within a single threaded, synchronous program, this would be incredibly easy. We could effectively guarantee that it would work as intended every time with a small battery of unit tests
- However, if we were to start introducing multiple threads, or multiple goroutines in Go’s case, we may start to see issues within our code.

### A Simple Example
- Example:
  ```go
  package main

  import (
      "fmt"
      "sync"
  )

  var (
      mutex   sync.Mutex
      balance int
  )

  func init() {
      balance = 1000
  }

  func deposit(value int, wg *sync.WaitGroup) {
      mutex.Lock()
      fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
      balance += value
      mutex.Unlock()
      wg.Done()
  }

  func withdraw(value int, wg *sync.WaitGroup) {
      mutex.Lock()
      fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
      balance -= value
      mutex.Unlock()
      wg.Done()
  }

  func main() {
      fmt.Println("Go Mutex Example")

  	var wg sync.WaitGroup
  	wg.Add(2)
      go withdraw(700, &wg)
      go deposit(500, &wg)
      wg.Wait()

      fmt.Printf("New Balance %d\n", balance)
  }
  ```
- So, let’s break down what we have done here. Within both our deposit() and our withdraw() functions, we have specified the first step should be to acquire the mutex using the mutex.Lock() method.
- Each of our functions will block until it successfully acquires the Lock. Once successful, it will then proceed to enter it’s critical section in which it reads and subsequently updates the account’s balance. Once each function has performed it’s task, it then proceeds to release the lock by calling the mutex.Unlock() method.

### Avoiding Deadlock
- Deadlock is a scenario within our code where nothing can progress due to every goroutine continually blocking when trying to attain a lock.
- If you are developing goroutines that require this lock and they can terminate in a number of different ways, then ensure that regardless of how your goroutine terminates, it always calls the Unlock() method.
- Calling Lock() Twice

### Semaphore vs Mutex
- Everything you can achieve with a Mutex can be done with a channel in Go if the size of the channel is set to 1.
- However, the use case for what is known as a binary semaphore - a semaphore/channel of size 1 - is so common in the real world that it made sense to implement this exclusively in the form of a mutex.

### My example:
- Example:
  ```go
  func appendToSlice(s *[]int, wg *sync.WaitGroup, mu *sync.Mutex) {
  	mu.Lock()
  	fmt.Println("Appending the slice")
  	*s = append(*s, 10)
  	mu.Unlock()
  	wg.Done()
  }

  func popTheSlice(s *[]int, wg *sync.WaitGroup, mu *sync.Mutex) {
  	mu.Lock()
  	fmt.Println("Popping the slice")
  	*s = (*s)[:len(*s)-1]
  	mu.Unlock()
  	wg.Done()
  }

  func main() {
  	var wg sync.WaitGroup
  	var mu sync.Mutex

  	s := []int{1, 2, 3, 4, 5}
  	wg.Add(2)
  	go appendToSlice(&s, &wg, &mu)
  	go popTheSlice(&s, &wg, &mu)
  	fmt.Println("I don't know what I am doing in here")
  	fmt.Println("Yeah... Me too!")

  	wg.Wait()
  	fmt.Println("s", s)
  }
  ```


**[⬆ back to top](#list-of-contents)**

</br>

---

## [Deep dive on goroutine leaks and best practices to avoid them](https://mourya-g9.medium.com/deep-dive-on-goroutine-leaks-and-best-practices-to-avoid-them-a35021383f64) <span id="content-6"></span>

### Introduction
- Using goroutines and channels in production env without proper context on how they behave can cause some serious effects.
- Well, we faced one such impact where we had a leakage in goroutines that resulted in the application server bloating over time by consuming abundant CPU & frequent GC pauses affecting the SLA of multiple APIs.
- A goroutine leak is where the client spawns a goroutine to do some async task and writes some data to a channel once the task is done but
  - There is no listener consuming from that channel to which the data is being written.
    ```go
    func newgoroutine(dataChan chan <dataType>) {
        data := makeNetworkCall()
        dataChan <- data
        return
    }
    func main() {
        dataChan := make(chan <dataType>)
        go newgoroutine(dataChan)
        // Some application processing (but forgot to consume data from the channel (dataChan))
        return 
    }
    ```
  - In the above scenario, the code completes execution succesfully as if there is no issue at all. But what happens here is that, there will be a dangling goroutine that resides in memory eating up the CPU & RAM. 
  - The major reason for that is because of line 3 where we are writing data into a channel but as per go principles, an unbuffered channel blocks write to channel until consumer consumes the message from that channel.
  - So in this case the return on line number 4 will never get executed and the newgoroutine function gets stuck throughtout the application lifetime as there is no consumer for this channel.
  - There is some conditional logic between the goroutine start and channel listener.
    ```go
    func newgoroutine(dataChan chan <dataType>) {
        data := makeNetworkCall()
        dataChan <- data
        return
    }
    func main() {
        dataChan := make(chan <dataType>)
        go newgoroutine(dataChan)
        // Some application processing 
        if processingError != nil {
              return
        }
        data := <- dataChan
        // Do something with data
        return 
    }
    ```
  - We had a consumer consuming the data from the dataChan but from the time we spawned the goroutine and before we started consuming the data from the channel, there is a ton of application code that resides which can quit the main function on some processing error | DB error | Nil pointer exceptions | Panics due to which the consumer of the data channel never gets executed. 
  - The forgotten sender
    ```go
  func newgoroutine(dataChan chan <dataType>) {
      // Consume data from dataChan 
      data := <- dataChan
      // Do some processing on the data
      return
  }
  func main() {
      dataChan := make(chan <dataType>)
      go newgoroutine(dataChan)
      data, err := makeNetworkCall()
      if err != nil {
          return 
      }
      dataChan <- data // This piece of code is never executed in error case of networkCall which makes newgoroutine dangling
      // Do something with data
      return 
  }
    ```

### Approaches
- Approach -> We identify every error condition from the time we started the goroutine till we consume from the channel where we exit and place a receiver before every return statement just to unblock the spawned goroutine.
- Pitfall -> We have to find all edge cases manually and in the future, if we have to handle one more error condition, we need to remember what all channels we need to consume data from before returning. Buggy solution.
- Approach -> Instead of placing a receiver at every error case, why not have a defer function that can receive the data from the channel.
- Pitfall -> In case of success the data will be read from the channel after processing the static rules. So if we start to receive data from the channel at defer function this blocks the main goroutine in case of success. Faulty solution.
- The major problem here is we aren’t sure whether the receiver flow will be executed or not due to our application processing. 
- Well, the simple solution is to create a buffered channel with cap 1. With this, the sender is never blocked to write the data once even if there is no consumer spawned or the spawned consumer code is not reached.
- Pitfalls -> Absolutely zero. This works exactly like unbuffered channels but provides us an extra capability where sender is not blocked to send the data once and the consumer can consume it at any point and the spawned goroutine won’t be waiting for the consumer.

### The approach to find goroutine leaks
- When the server starts, disable Garbage Collector using debug.SetGCPercent(-1)
- Now run every flow in the code where a Go routine is used(Dev Env).
- At the entry point of each API, print the no of running goroutines before starting & after executing the API
- Now if a service returns a different count of Goroutines before & after, then there is a leak in that flow.


**[⬆ back to top](#list-of-contents)**

</br>

---
## References:
- https://www.golang-book.com/books/intro/10
- https://tutorialedge.net/golang/concurrency-with-golang-goroutines/
- https://tutorialedge.net/golang/go-channels-tutorial/
- https://tutorialedge.net/golang/go-waitgroup-tutorial/
- https://tutorialedge.net/golang/go-mutex-tutorial/
- https://mourya-g9.medium.com/deep-dive-on-goroutine-leaks-and-best-practices-to-avoid-them-a35021383f64