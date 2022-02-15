# Practical Go

</br>

## [List of Contents:](#list-of-contents)
### [1. Guiding principles](#content-1)
### [2. Identifiers](#content-2)
### [3. Comments](#content-3)
### [4. Package Deisgn](#content-4)


</br>

---

## Contents:

## [1. Guiding principles](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_guiding_principles) <span id="content-1"></span>
### Introduction

> Software engineering is what happens to programming when you add time and other programmers. -- Russ Cox

- Guiding principles:
  - Simplicity
  - Readability
  - Productivity
- Performance and concurrency are important attributes, but not as important as simplicity, readability, and productivity.

### 1.1. Simplicity
> Simplicity is prerequisite for reliability. — Edsger W. Dijkstra

- We’ve all worked on programs where you’re scared to make a change because you’re worried it’ll break another part of the program; a part you don’t understand and don’t know how to fix. This is complexity.
> There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies, and the other way is to make it so complicated that there are no obvious deficiencies. The first method is far more difficult. — C. A. R. Hoare

- Complexity turns reliable software in unreliable software. Complexity is what kills software projects. Therefore simplicity is the highest goal of Go. Whatever programs we write, we should be able to agree that they are simple.

### 1.2. Readability
> Readability is essential for maintainability. — Mark Reinhold JVM language summit 2018

> Programs must be written for people to read, and only incidentally for machines to execute. — Hal Abelson and Gerald Sussman Structure and Interpretation of Computer Programs

- Readability is important because all software, not just Go programs, is written by humans to be read by other humans. The fact that software is also consumed by machines is secondary.
- Code is read many more times than it is written. A single piece of code will, over its lifetime, be read hundreds, maybe thousands of times.

> The most important skill for a programmer is the ability to effectively communicate ideas. — Gastón Jorquera

- Readability is key to being able to understand what the program is doing. If you can’t understand what a program is doing, how can you hope to maintain it? If software cannot be maintained, then it will be rewritten; and that could be the last time your company will invest in Go.
- If it will be used and contributed by so many people, it's important to make it maintainable.
- The first step towards writing maintainable code is making sure the code is readable.

### 1.3. Productivity
> Design is the art of arranging code to work today, and be changeable forever. — Sandi Metz

- Developer productivity is a sprawling topic but it boils down to this; how much time do you spend doing useful work verses waiting for your tools or hopelessly lost in a foreign code-base. Go programmers should feel that they can get a lot done with Go.
- More fundamental to the question of developer productivity, Go programmers realise that code is written to be read and so place the act of reading code above the act of writing it.
- Productivity is what the Go team mean when they say the language must scale.


**[⬆ back to top](#list-of-contents)**

</br>

---

## [2. Identifiers](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_identifiers) <span id="content-2"></span>

### Introduction
- An identifier is a fancy word for a name; the name of a variable, the name of a function, the name of a method, the name of a type, the name of a package, and so on.
> Poor naming is symptomatic of poor design. — Dave Cheney

### 2.1. Choose identifiers for clarity, not brevity
> Obvious code is important. What you can do in one line you should do in three. — Ukiah Smith

- Go is not a language that optimises for clever one liners. Go is not a language which optimises for the least number of lines in a program. We’re not optimising for the size of the source code on disk, nor how long it takes to type the program into an editor.
> Good naming is like a good joke. If you have to explain it, it’s not funny. — Dave Cheney

- Let’s talk about the qualities of a good name:
  - A good name is concise. A good name need not be the shortest it can possibly be, but a good name should waste no space on things which are extraneous. Good names have a high signal to noise ratio.
  - A good name is descriptive. A good name should describe the application of a variable or constant, not their contents. A good name should describe the result of a function, or behaviour of a method, not their implementation. A good name should describe the purpose of a package, not its contents. The more accurately a name describes the thing it identifies, the better the name.
  - A good name is should be predictable. You should be able to infer the way a symbol will be used from its name alone. This is a function of choosing descriptive names, but it also about following tradition. This is what Go programmers talk about when they say idiomatic.

### 2.2. Identifier length
- Sometimes people criticise the Go style for recommending short variable names. As Rob Pike said, "Go programmers want the right length identifiers".
> The greater the distance between a name’s declaration and its uses, the longer the name should be. — Andrew Gerrand 

- From this we can draw some guidelines:
  - Short variable names work well when the distance between their declaration and last use is short.
  - Long variable names need to justify themselves; the longer they are the more value they need to provide. Lengthy bureaucratic names carry a low amount of signal compared to their weight on the page.
  - Don’t include the name of your type in the name of your variable.
  - Constants should describe the value they hold, not how that value is used.
  - Prefer single letter variables for loops and branches, single words for parameters and return values, multiple words for functions and package level declarations
  - Prefer single words for methods, interfaces, and packages.
  - Remember that the name of a package is part of the name the caller uses to to refer to it, so make use of that.
- Snippet:
  ```go
  type Person struct {
    Name string
    Age  int
  }

  // AverageAge returns the average age of people.
  func AverageAge(people []Person) int {
    if len(people) == 0 {
      return 0
    }

    var count, sum int
    for _, p := range people {
      sum += p.Age
      count += 1
    }

    return sum / count
  }
  ```
- Use blank lines to break up the flow of a function in the same way you use paragraphs to break up the flow of a document. In AverageAge we have three operations occurring in sequence. The first is the precondition, checking that we don’t divide by zero if people is empty, the second is the accumulation of the sum and count, and the final is the computation of the average.

### 2.2.1. Context is key
- Snippet:
  ```go
  for index := 0; index < len(s); index++ {
    //
  }

  for i := 0; i < len(s); i++ {
    //
  }
  ```
- Don’t mix and match long and short formal parameters in the same declaration.

### 2.3. Don’t name your variables for their types
- You shouldn’t name your variables after their types for the same reason you don’t name your pets "dog" and "cat".
- Snippet:
  ```go
  var usersMap map[string]*User
  ```
- We can see that its a map, and it has something to do with the *User type, that’s probably good. But usersMap is a map, and Go being a statically typed language won’t let us accidentally use it where a scalar variable is required, so the Map suffix is redundant.
- If users isn’t descriptive enough, then usersMap won’t be either.
- Snippet:
  ```go
  type Config struct {
    //
  }

  func WriteConfig(w io.Writer, config *Config)
  ```
- In this case consider conf or maybe c will do if the lifetime of the variable is short enough.
- If there is more that one *Config in scope at any one time then calling them conf1 and conf2 is less descriptive than calling them original and updated as the latter are less likely to be mistaken for one another.
- Don’t let package names steal good variable names.

### 2.4. Use a consistent naming style
- Another property of a good name is it should be predictable. The reader should be able to understand the use of a name when they encounter it for the first time. When they encounter a common name, they should be able to assume it has not changed meanings since the last time they saw it.
- For example, if your code passes around a database handle, make sure each time the parameter appears, it has the same name. Rather than a combination of d *sql.DB, dbase *sql.DB, DB *sql.DB, and database *sql.DB, instead consolidate on something like;
- Go style dictates that receivers have a single letter name, or acronyms derived from their type. You may find that the name of your receiver sometimes conflicts with name of a parameter in a method. In this case, consider making the parameter name slightly longer, and don’t forget to use this new parameter name consistently.
- Finally, certain single letter variables have traditionally been associated with loops and counting. For example, i, j, and k are commonly the loop induction variable for simple for loops. n is commonly associated with a counter or accumulator. v is a common shorthand for a value in a generic encoding function, k is commonly used for the key of a map, and s is often used as shorthand for parameters of type string.

### 2.5. Use a consistent declaration style
- How to declare variables in Go:
  - `var x int = 1`
  - `var x = 1`
  - `var x int; x = 1`
  - `var x = int(1)`
  - `x := 1`
- When declaring, but not initialising, a variable, use var. When declaring a variable that will be explicitly initialised later in the function, use the var keyword.
  ```go
  var players int    // 0

  var things []Thing // an empty slice of Things

  var thing Thing    // empty Thing struct
  json.Unmarshall(reader, &thing)
  ```
- When declaring and initialising, use :=. When declaring and initialising the variable at the same time, that is to say we’re not letting the variable be implicitly initialised to its zero value, I recommend using the short variable declaration form. 
  ```go
  var players int = 0

  var things []Thing = nil

  var thing *Thing = new(Thing)
  json.Unmarshall(reader, thing)
  ```
- You can't declare a variable to be nil, because nil does not have a type.  Instead we have a choice, do we want the zero value for a slice?
  ```go
  // Zero value for a slice
  var things []Thing

  // A slice with zero value
  var things = make([]Thing, 0)
  ```
- Therefore it's better to use short declaration form:
  ```go
  things := make([]Thing, 0)
  ```
- A pointer of a Thing:
  ```go
  // Pointer of a Thing
  thing := new(Thing)

  // It's better to use compact literal struct initialiser form
  thing := &Thing{}
  ```
- In summary:
  - When declaring a variable without initialisation, use the var syntax.
  - When declaring and explicitly initialising a variable, use :=.

> My advice in this situation is to follow the local style.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [3. Comments](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_comments) <span id="content-3"></span>

> Good code has lots of comments, bad code requires lots of comments. — Dave Thomas and Andrew Hunt. The Pragmatic Programmer

- Each comments should do one—​and only one—​of three things:
  - The comment should explain what the thing does.
  - The comment should explain how the thing does what it does.
  - The comment should explain why the thing is why it is.
- The first form is ideal for commentary on public symbols:
  ```go
  // Open opens the named file for reading.
  // If successful, methods on the returned file can be used for reading.
  ```
- The second form is ideal for commentary inside a method:
  ```go
  // queue all dependant actions
  var results []chan error
  for _, dep := range a.Deps {
          results = append(results, execute(seen, dep))
  }
  ```
- The third form, the why , is unique as it does not displace the first two, but at the same time it’s not a replacement for the what, or the how. The why style of commentary exists to explain the external factors that drove the code you read on the page. Frequently those factors rarely make sense taken out of context, the comment exists to provide that context.
  ```go
  return &v2.Cluster_CommonLbConfig{
    // Disable HealthyPanicThreshold
      HealthyPanicThreshold: &envoy_type.Percent{
        Value: 0,
      },
  }
  ```
- In this example it may not be immediately clear what the effect of setting HealthyPanicThreshold to zero percent will do. The comment is needed to clarify that the value of 0 will disable the panic threshold behaviour.

### 3.1. Comments on variables and constants should describe their contents not their purpose
- When you add a comment to a variable or constant, that comment should describe the variables contents, not the variables purpose.
  ```go
  const randomNumber = 6 // determined from an unbiased die
  ```
- More examples:
  ```go
  const (
      StatusContinue           = 100 // RFC 7231, 6.2.1
      StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
      StatusProcessing         = 102 // RFC 2518, 10.1

      StatusOK                 = 200 // RFC 7231, 6.3.1
  ```
- For variables without an initial value, the comment should describe who is responsible for initialising this variable.
  ```go
  // sizeCalculationDisabled indicates whether it is safe
  // to calculate Types' widths and alignments. See dowidth.
  var sizeCalculationDisabled bool
  ```
- This is a tip from Kate Gregory. Sometimes you’ll find a better name for a variable hiding in a comment. Now the comment is redundant and can be removed.
  ```go
  // registry of SQL drivers. But registry of what?
  var registry = make(map[string]*sql.Driver)

  // By renaming the variable to sqlDrivers its now clear that the purpose of this variable is to hold SQL drivers. Now the comment is redundant and can be removed.
  var sqlDrivers = make(map[string]*sql.Driver)
  ```

### 3.2. Always document public symbols
- Here are two rules from the Google Style guide
  - Any public function that is not both obvious and short must be commented.
  - Any function in a library must be commented regardless of length or complexity
- Example:
  ```go
  package ioutil

  // ReadAll reads from r until an error or EOF and returns the data it read.
  // A successful call returns err == nil, not err == EOF. Because ReadAll is
  // defined to read from src until EOF, it does not treat an EOF from Read
  // as an error to be reported.
  func ReadAll(r io.Reader) ([]byte, error)
  ```
- There is one exception to this rule; you don’t need to document methods that implement an interface. Specifically don’t do this:
  ```go
  // Read implements the io.Reader interface
  func (r *FileReader) Read(buf []byte) (int, error)
  ```
- Example from the io package:
  ```go
  // LimitReader returns a Reader that reads from r
  // but stops with EOF after n bytes.
  // The underlying implementation is a *LimitedReader.
  func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }

  // A LimitedReader reads from R but limits the amount of
  // data returned to just N bytes. Each call to Read
  // updates N to reflect the new amount remaining.
  // Read returns EOF when N <= 0 or when the underlying R returns EOF.
  type LimitedReader struct {
    R Reader // underlying reader
    N int64  // max bytes remaining
  }

  func (l *LimitedReader) Read(p []byte) (n int, err error) {
    if l.N <= 0 {
      return 0, EOF
    }
    if int64(len(p)) > l.N {
      p = p[0:l.N]
    }
    n, err = l.R.Read(p)
    l.N -= int64(n)
    return
  }
  ```
- Note that the LimitedReader declaration is directly preceded by the function that uses it, and the declaration of LimitedReader.Read follows the declaration of LimitedReader itself. Even though LimitedReader.Read has no documentation itself, its clear from that it is an implementation of io.Reader.
- TIP: Before you write the function, write the comment describing the function. If you find it hard to write the comment, then it’s a sign that the code you’re about to write is going to be hard to understand.

### 3.2.1. Don’t comment bad code, rewrite it
> Don’t comment bad code — rewrite it — Brian Kernighan

- Comments highlighting the grossness of a particular piece of code are not sufficient. If you encounter one of these comments, you should raise an issue as a reminder to refactor it later. It is okay to live with technical debt, as long as the amount of debt is known.
- The tradition in the standard library is to annotate a TODO style comment with the username of the person who noticed it.
  ```
  // TODO(dfc) this is O(N^2), find a faster way to do this.
  ```
- The username is not a promise that that person has committed to fixing the issue, but they may be the best person to ask when the time comes to address it. Other projects annotate TODOs with a date or an issue number.
- Don't give a make up to bad code by adding comments. Instead make sure your code is clean then you can add make ups to your code by adding comments.

### 3.2.2. Rather than commenting a block of code, refactor it
> Good code is its own best documentation. As you’re about to add a comment, ask yourself, 'How can I improve the code so that this comment isn’t needed?' Improve the code and then document it to make it even clearer. — Steve McConnell

- Functions should do one thing only. If you find yourself commenting a piece of code because it is unrelated to the rest of the function, consider extracting it into a function of its own.
- In addition to being easier to comprehend, smaller functions are easier to test in isolation. Once you’ve isolated the orthogonal code into its own function, its name may be all the documentation required.
  

**[⬆ back to top](#list-of-contents)**

</br>

---

## [4. Package Design](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_package_design) <span id="content-4"></span>

### Introduction
> Write shy code - modules that don’t reveal anything unnecessary to other modules and that don’t rely on other modules' implementations.
- Each Go package is in effect it’s own small Go program. Just as the implementation of a function or method is unimportant to the caller, the implementation of the functions, methods and types that comprise your package’s public API—​its behaviour—​is unimportant for the caller.
  
### 4.1. A good package starts with its name
- Writing a good Go package starts with the package’s name. Think of your package’s name as an elevator pitch to describe what it does using just one word.
- Name your package for what it provides, not what it contains.
- Just as I talked about names for variables in the previous section, the name of a package is very important. The rule of thumb I follow is not, "what types should I put in this package?". Instead the question I ask "what does service does package provide?" Normally the answer to that question is not "this package provides the X type", but "this package let’s you speak HTTP".
- Within your project, each package name should be unique.
- If you find you have two packages which need the same name, it is likely either;
  - The name of the package is too generic.
  - The package overlaps another package of a similar name. In this case either you should review your design, or consider merging the packages.

### 4.2. Avoid package names like base, common, or util
> [A little] duplication is far cheaper than the wrong abstraction. — Sandy Metz
- Use plurals for naming utility packages. For example the strings for string handling utilities.
- Packages with names like base or common are often found when functionality common to two or more implementations, or common types for a client and server, has been refactored into a separate package. I believe the solution to this is to reduce the number of packages, to combine the client, server, and common code into a single package named after the function of the package.
- For example, the net/http package does not have client and server sub packages, instead it has a client.go and server.go file, each holding their respective types, and a transport.go file for the common message transport code.
- An identifier’s name includes its package name.
It’s important to remember that the name of an identifier includes the name of its package.
  - The Get function from the net/http package becomes http.Get when referenced by another package.
  - The Reader type from the strings package becomes strings.Reader when imported into other packages.
  - The Error interface from the net package is clearly related to network errors.

### 4.3. Return early rather than nesting deeply
- This is achieved by using guard clauses; conditional blocks with assert preconditions upon entering a function. Here is an example from the bytes package,
- Example:
  ```go
  func (b *Buffer) UnreadRune() error {
    if b.lastRead <= opInvalid {
      return errors.New("bytes.Buffer: UnreadRune: previous operation was not a successful ReadRune")
    }
    if b.off >= int(b.lastRead) {
      b.off -= int(b.lastRead)
    }
    b.lastRead = opInvalid
    return nil
  }
  ```
- Compare the above function with this:
  ```go
  func (b *Buffer) UnreadRune() error {
    if b.lastRead > opInvalid {
      if b.off >= int(b.lastRead) {
        b.off -= int(b.lastRead)
      }
      b.lastRead = opInvalid
      return nil
    }
    return errors.New("bytes.Buffer: UnreadRune: previous operation was not a successful ReadRune")
  }
  ```

### 4.4. Make the zero value useful
- Every variable declaration, assuming no explicit initialiser is provided, will be automatically initialised to a value that matches the contents of zeroed memory. This is the values zero value.
- Consider the sync.Mutex type. sync.Mutex contains two unexported integer fields, representing the mutex’s internal state. Thanks to the zero value those fields will be set to will be set to 0 whenever a sync.Mutex is declared. sync.Mutex has been deliberately coded to take advantage of this property, making the type usable without explicit initialisation.
- Example:
  ```go
  type MyInt struct {
    mu  sync.Mutex
    val int
  }

  func main() {
    var i MyInt

    // i.mu is usable without explicit initialisation.
    i.mu.Lock()
    i.val++
    i.mu.Unlock()
  }
  ```
- Snippet:
  ```go
  func main() {
    // s := make([]string, 0)
    // s := []string{}
    var s []string

    s = append(s, "Hello")
    s = append(s, "world")
    fmt.Println(strings.Join(s, " "))
  }
  ```
- A useful, albeit surprising, property of uninitialised pointer variables—​nil pointers—​is you can call methods on types that have a nil value. This can be used to provide default values simply.
  ```go
  type Config struct {
    path string
  }

  func (c *Config) Path() string {
    if c == nil {
      return "/usr/home"
    }
    return c.path
  }

  func main() {
    var c1 *Config
    var c2 = &Config{
      path: "/export",
    }
    fmt.Println(c1.Path(), c2.Path())
  }
  ```
- The key to writing maintainable programs is that they should be loosely coupled—​a change to one package should have a low probability of affecting another package that does not directly depend on the first.
- There are two excellent ways to achieve loose coupling in Go
  - Use interfaces to describe the behaviour your functions or methods require.
  - Avoid the use of global state.
- If you want to reduce the coupling a global variable creates,
  - Move the relevant variables as fields on structs that need them.
  - Use interfaces to reduce the coupling between the behaviour and the implementation of that behaviour.



**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://dave.cheney.net/practical-go/presentations/qcon-china.html