# Idiomatic Go

</br>

## List of Contents:
### 1. [The Zen of Go](#content-1)
### 2. [Golang UK Conference 2016 - Mat Ryer - Idiomatic Go Tricks](#content-2)
### 3. [Idiomatic Go](#content-3)


</br>

---

## Contents:

## [The Zen of Go](https://dave.cheney.net/2020/02/23/the-zen-of-go) <span id="content-1"></span>



### Idiomatic Go
- To say that something is idiomatic is to say that it follows the style of the time.
- If something is not idiomatic, it is not following the prevailing style. It is unfashionable.
- idiom (noun): a group of words established by usage as having a meaning not deducible from those of the individual words.

### Proverbs
- proverb (noun): a short, well-known pithy saying, stating a general truth or piece of advice.
- The goal of the Go Proverbs are to reveal a deeper truth about the design of the language, but how useful is advice like the empty interface says nothing to a novice from a language that doesn’t have structural typing?

### Engineering Values
- An image about engineering culture: </br>
  ![engineering-culture](https://dave.cheney.net/wp-content/uploads/2020/02/Lucovsky.001.jpeg)
- The central idea is values guide decisions in an unknown design space. 
  
### Go’s values
- This process of knowledge transfer is not optional. Without new blood and new ideas, our community become myopic and wither.

### The values of other languages\
- Discourse in our community is often fractious, so deriving a set of values from first principles would be a formidable challenge.
- Consensus is critical, but exponentially more difficult as the number of contributors to the discussion increases.

### A good package starts with a good name
- “Namespaces are one honking great idea–let’s do more of those!” The Zen of Python, Item 19
- In Go each package should have a purpose, and the best way to know a package’s purpose is by its name—a noun. A package’s name describes what it provides.
- Every Go package should have a single purpose.
- “Design is the art of arranging code to work today, and be changeable forever.” Sandi Metz
- What we do as programmers is manage change. When we do that well we call it design, or architecture. When we do it badly we call it technical debt, or legacy code.
- However, for the majority of programs, designing something to be flexible up front is over engineering.
- A good package starts with choosing a good name. Think of your package’s name as an elevator pitch, using just one word, to describe what it provides.

### Simplicity matters
- “Simple is better than complex.” The Zen of Python, Item 3
- “There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies, and the other way is to make it so complicated that there are no obvious deficiencies. The first method is far more difficult.” C. A. R. Hoare, The Emperor’s Old Clothes, 1980 Turing Award Lecture
- Simple does not mean easy, we know that. Often it is more work to make something simple to use, than easy to build.
- “Simplicity is prerequisite for reliability.” Edsger W Dijkstra, EWD498, 18 June 1975
- Simple doesn’t mean crude, it means readable and maintainable. Simple doesn’t mean unsophisticated, it means reliable, relatable, and understandable.
- “Controlling complexity is the essence of computer programming.” Brian W. Kernighan, Software Tools (1976)
- Simple code is preferable to clever code.

### Avoid package level state
- “Explicit is better than implicit.” The Zen of Python, Item 2
- The more valuable, in my opinon, place to be explicit are to do with coupling and with state.
- Coupling is a measure of the amount one thing depends on another. If two things are tightly coupled, they move together. An action that affects one is directly reflected in another. Imagine a train, each carriage joined–ironically the correct word is coupled–together; where the engine goes, the carriages follow.
- Cohesion measures how well two things naturally belong together.
- Avoid package level state. Reduce coupling and spooky action at a distance by providing the dependencies a type needs as fields on that type rather than using package variables.

### Plan for failure, not success
- “Errors should never pass silently.” The Zen of Python, Item 10
- Unchecked exceptions are clearly an unsafe model to program in.
- Go programmers believe that robust programs are composed from pieces that handle the failure cases before they handle the happy path.
- “I think that error handling should be explicit, this should be a core value of the language.” Peter Bourgon, GoTime #91
- I think so much of the success of Go is due to the explicit way errors are handled.
- Key to this is the cultural value of handling each and every error explicitly.

### Return early rather than nesting deeply
- “Flat is better than nested.” The Zen of Python, Item 5
- In my experience the more a programmer tries to subdivide and taxonimise their Go codebase the more they risk hitting the dead end that is package import loops.
- Simply put, avoid control flow that requires deep indentation.
- “Line of sight is a straight line along which an observer has unobstructed vision.”May Ryer, Code: Align the happy path to the left edge
- Light of sight coding means things like:
  - Using guard clauses to return early if a precondition is not met.
  - Placing the successful return statement at the end of the function rather than inside a conditional block.
  - Reducing the overall indentation level of the function by extracting functions and methods.
- Rather than nesting deeply, keep the successful path of the function close to the left hand side of your screen.
- “In the face of ambiguity, refuse the temptation to guess.” The Zen of Python, Item 12
- “APIs should be easy to use and hard to misuse.” Josh Bloch
- Don’t complicate your code because of outdated dogma, and, if you think something is slow, first prove it with a benchmark.

### Before you launch a goroutine, know when it will stop
- Goroutines are cheap. Because of the runtime’s ability to multiplex goroutines onto a small pool of threads (which you don’t have to manage), hundreds of thousands, millions of goroutines are easily accommodated.
- While that goroutine is alive, the lock is held, the network connection remains open, the buffer retained and the receivers of the channel will continue to wait for more data.
- The simplest way to free those resources is to tie them to the lifetime of the goroutine–when the goroutine exits, the resource has been freed. 


### Write tests to lock in the behaviour of your package’s API
- Your tests are the contract about what your software does and does not do. \
- If there is a unit test for each input permutation, you have defined the contract for what the code will do in code, not documentation.
- Tests lock in api behaviour. Any change that adds, modifies or removes a public api must include changes to its tests.


### Moderation is a virtue
- I had the same experience with embedding. Initially I mistook it for inheritance.
- Then later I recreated the fragile base class problem by composing complicated types, which already had several responsibilities, into more complicated mega types.
- If you can, don’t reach for a goroutine, or a channel, or embed a struct, anonymous functions, going overboard with packages, interfaces for everything, instead prefer simpler approach rather than the clever approach.

### Maintainability counts
- “Readability Counts.” The Zen of Python, Item 7
- Go use words like simplicity, readability, clarity, productivity, but ultimately they are all synonyms for one word–maintainability.
- Rather, we want to optimise our code to be clear to the reader. Because its the reader who’s going to have to maintain this code.
- If you’re writing a program for yourself, maybe it only has to run once, or you’re the only person who’ll ever see it, then do what ever works for you.
- But if this is a piece of software that more than one person will contribute to, or that will be used by people over a long enough time that requirements, features, or the environment it runs in may change, then your goal must be for your program to be maintainable.



</br>

---


## [Golang UK Conference 2016 - Mat Ryer - Idiomatic Go Tricks](https://www.youtube.com/watch?v=yeetIgNeIkc) <span id="content-2"></span>

- Idiomatic </br>
  adjective: Using, containing, or denoting expressions that are natural to a native speaker
- `defer` to do something when the function returns
- Line of sight:
  - definition: "a straight line along which an observer has unobstracted vision"
  - Happy path is aligned to the left
  - Error handling and edge cases indented
- Example of bad line of sight: </br>
  ```go
  func DoSomething() error {
    val, err := GetSomething()
    if err == nil {
      ... do something

    } else {
      return err
    }

    defer val.Close()

    result, err := GetMeAnother()
    if err == nil {
      for _, v := range result {
        item, err := GetOne()
        if err == nil {
          ...
        } else {
          ...
        }
      }
    } else {
      ...
    }
  }
  ```
- Line of sight tips:
  - Make happy return that last statement if possible
  - Next time we write else, consider flipping the logic </br>
    ```go
    // Before
    if something.OK() {
      ...
    } else {
      return false, err
    }

    // After
    if !something.OK() {
      return false, err
    }

    ...
    return true, nil
    ```
- Single method interfaces
  - Example: </br>
    ```go
    type Reader interface {
      Read(p []byte)(n int, err error)
    }
    ```
  - Interface consisting of only one method
  - Simpler = more powerful and useful
  - Easy to implement
  - Used throughout the standard library
- Log blocks
  - Better way: </br>
    ```go
    func foo() error {
      log.Println("--------")
      defer log.Println("--------")

      ...
    }
    ```
- Regarding unit testing
  - Return teardown function when setup the test

- Good timing
  - Let's write a function to measure the time a function is running </br>
    ```go
    func Timeit(name string) func() {
      t := time.Now()
      log.Println("Started")
      return func() {
        d := time.Now()
        log.Println(name, "took", d)
      }
    }


    func MyFunc() error {
      stop := Timeit()
      defer stop()

      ...
    }
    ```

- Discover interface
  - If we have several functions that have the same function signature call and can be called by another function, then it's better to use that function who is calling an interface as an argument.
  - Example: </br>
    ```go
    type Sizer interface {
      Size() int64
    }

    func Fits(capacity int64, y Sizer) bool {
      return capacity > y.Size()
    }

    type Sizers []Sizer

    func (s Sizers) Size() int64 {
      var total int64
      for _, sizer := range s {
        total += sizer.Size()
      }

      return total
    }
    ```
- Using simple mocks:
  - Just write a simple mock using a struct </br>
    ```go
    type MailSender interface {
      Send()
      SendFrom()
    }
    type MockedMailSender struct {
      SendFunc func()
      SendFromFunc func()
    }
    ```
- Mocking other people's structs
  - Let's say somebody writes a struct with some methods in it. Since, there is no interface to mock, then just make our own.

- Empty struct implementations
  - Empty struct{} to group methods together
  - Methods don't capture the receiver

- Be obvious not clever
  - Example:  </br>
    ```go
    // Better don't do this
    func Something() error {
      defer Timeit()
    }

    // Better to do this
    func Something() error {
      stop := Timeit() 
      defer stop()
    }
    ```

- How to become a native speaker
  - Read the standard library
  - Write obvious code (not clever)
  - Don't surprise your users
  - Seek simplicity
  - Learn from others
  - Participate in open-source projects
  - Ask for reviews and accept critisims
  - Help others when you spot something (and be kind)


</br>

---

## [Idiomatic Go](https://about.sourcegraph.com/go/idiomatic-go/) <span id="content-3"></span>

### What is idiomatic Go?
- Orthogonality: In programming terms "orthogonality" means that pieces are independent from each other. Changes to one part, a type, package, program, etc, have minimal to no effect on other parts.
- Simplicity
  - Go has no classes; methods can be added to any type.
  - It has no inheritance; interfaces are implicitly satisfied.
  - Types stand alone by themselves; they just are and have no hierarchy.
  - Methods aren’t special; they’re just functions.
- Readability: Go reduces clutter and noise. There are no header files, unsurprising syntax, and everything can only be declared once per block.

### Anti-patterns

### Tiny Package Syndrome
- The tiny package syndrome anti-pattern usually appears with a directory structure that looks something like this: </br>
  ```go
  context/                context (cont...)/
    cqlsession/             requestid/
      cqlsession.go           requestid.go
    dao/                    sinkctx/
      dao.go                  sink.go
    api/                    starttime/
      api.go                  starttime.go
    logtoken/               time/
      logtoken.go             time.go
    metricsapi/             tx/
      metricsapi.go           tx.go
    outlet/                 user/
      outlet.go               user.go
    producers/              version/
      producers.go            version.go
  ```
- The better way: </br>
  ```go
   context/
    cqlsession.go
    dao.go
    api.go
    logtoken.go
    metricsapi.go
    outlet.go
    producers.go
    requestid.go
    sink.go
    starttime.go
    time.go
    tx.go
    user.go
    version.go
  ```
- Key takeaways:
  - Group related functionality in the same package
  - Be weary of lots of untested micro-packages under a single directory.


### Premature Exportation
- Developers have a drive to taxonimize things and split them into the smallest possible pieces. This excessive over taxonification, in order to derive the smallest, DRYest piece of code can result in many small packages. 
- The use of an `internal` directory can be used to signal intent that these packages are not meant for external consumption.
- Key takeaways:
  - Don’t export types, variables, function and constants until there is a need to do so.
  - The DRYest, smallest, most segmented packages lead to the need to export everything.
  - Keep packages that are not meant for external consumption in an internal folder.

### Package `util`
- A common pattern found in projects that do lot of group-by-kind is to have some miscellaneous pieces left over. Then we put them in `util`
- Notice this piece of code. It's better for us just to create package `generator` or `generate` </br>
  ```go
  func GenerateRandomBytes(n int) ([]byte, error) { ... }
  util.GenerateRandomBytes(10)

  func GenerateRandomString(n int) (string, error) { ... }
  util.GenerateRandomString(10)

  func Cert(hostname string) (string, string, error) { ... }
  util.Cert("foozle.com")
  ```
- Key takeaways:
  - Package names have semantic meaning.
  - Package names should describe the purpose of the package, not it’s contents.
  - The only part of the package import path that matters is the right most name.
  - `util` says nothing about the purpose of a package beyond a grouping of bits.


### Config Structs:
- It's common to see `Config` or `Options` with a lot of members </br>
  ```go
  type Config struct {
    MaxLineLength                       int
    BackBuff                            int
    BatchSize                           int
    NumOutlets                          int
    InputFormat                         int
    MaxAttempts                         int
    Prival                              string
    Procid                              string
    Hostname                            string
    Appname                             string
    Msgid                               string
    // a total of 28 members
  ```
- The problem is if we are just using several of these members, why should we put them in one config struct? Instead, pass in the only information each function uses. Then we can get a better sense of what the function needs.
- Key takeaways:
  - Config structs increase coupling
  - Config structs Obfuscate the API of the functions or types that accept them, and hides complexity.
  - Only pass in the information a function needs.

### Pointer All The Things
- Pointers are about ownership. When you pass a pointer to a function you are delegating ownership of the pointee to that function.
- If the pointer value “escapes” the function that created it, it may be moved to the heap, along with the value it points to. If you think it will be faster, benchmark it.
- Key takeaways:
  - Pointers are about the ownership of data.
  - They aren’t necessarily faster. Use benchmarks to prove that the additional overhead is worth it.
  - Remember one to the Go proverbs: Don’t communicate by sharing memory, share memory by communicating.

### context.Value
- Use context.Value() and context.WithValue() sparingly. They create undocumented, side channel APIs.
- If you still feel the need to use them, document the values that may be extracted from the context and the purpose for those values.
- Use WithTimeout(), WithDeadline() and WithCancel() as they are a great abstraction around cancelation.


### Asynchronous APIs
- Example: </br>
  ```go
  func Logs() <-chan logs {
      c := make(chan logs)
      go func() { // Receive Logs }
      return c
  }

  func main() {
      for l := range Logs() {
          // Do stuff with each l
      }
  }
  ```
- Issues about the above example:
  - How is the goroutine going to be shutdown?
  - How is the size of the channel controlled?
  - Closing the channel received from Logs() is likely to result in a panic
  - How do errors from the goroutine get communicated?
  - What if a synchronous API is needed?
- Key takeaways:
  - Provide synchronous APIs.
  - It is possible to make a synchronous API async, but much harder, if not impossible, to do the inverse. Leave the concurrency to someone else.
  - As a general rule, it should be uncommon to expose channels in an API. Use them internally instead.
  - For inspiration see the stdlib http package. This package has a fairly large API surface and does quite a lot. It uses channels internally, but only exposes one in the CloseNotifier type.

### If-then-else
- Example: </br>
  ```go
  func things(x int) someType {
  	if x > 2 {
  		return 100
  	} else {
  		return 200
  	}
  }

  // Can be simplifiend
  func things(x int) someType {
	if x > 2 {
  		return 100
  	}
  	return 200
  }
  ```
- Keep the expected "happy-path" of a function is de-dented to the left, instead of weaving it through various if-else blocks.
- Key takeaways:
  - Handle unexpected cases and errors early and return often.
  - Keep common or happy paths de-dented.
  - When it’s not possible refactor and/or redesign.


### Panic in a Lib
- Return errors, don’t panic.
- Only panic when an error can’t be handled directly or the handling can’t be delegated to the upstream caller.
- This should only happen when the program cannot make any forward progress.

### Interface All The Things
- Small, focused interfaces are the key to writing powerful and flexible go code. 
- These large interfaces tend to be defined up-front. They aren’t discovered across implementations at a later date. Interfaces should be discovered from existing types and extracted out of them.
- Don't overengineer by introducing interface early on
- The `io.Reader` and `io.Writer` interfaces weren't designed up front, they were discovered later.
- Key takeaways:
  - The bigger the interface, the weaker the abstraction.
  - Rethinking the abstraction and pivoting away from what is being done to how it’s being done can help. Though sometimes the inverse is true.
  - There are other ways to test things then mocking an interface. The httptest.Server type’s handlers and network servers are pretty easy to write in Go.

### interface{}
- The empty interface says nothing. Notice that i does not tell us anything.</br>
  ```go
  func Voila(i interface{}) { ... }

  ```

</br>

---


## References:
- https://dave.cheney.net/2020/02/23/the-zen-of-go
- https://www.youtube.com/watch?v=yeetIgNeIkc
- https://about.sourcegraph.com/go/idiomatic-go/