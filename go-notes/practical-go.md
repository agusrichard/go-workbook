# Practical Go

</br>

## [List of Contents:](#list-of-contents)

### [1. Guiding principles](#content-1)

### [2. Identifiers](#content-2)

### [3. Comments](#content-3)

### [4. Package Design](#content-4)

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
- We can see that its a map, and it has something to do with the \*User type, that’s probably good. But usersMap is a map, and Go being a statically typed language won’t let us accidentally use it where a scalar variable is required, so the Map suffix is redundant.
- If users isn’t descriptive enough, then usersMap won’t be either.
- Snippet:

  ```go
  type Config struct {
    //
  }

  func WriteConfig(w io.Writer, config *Config)
  ```

- In this case consider conf or maybe c will do if the lifetime of the variable is short enough.
- If there is more that one \*Config in scope at any one time then calling them conf1 and conf2 is less descriptive than calling them original and updated as the latter are less likely to be mistaken for one another.
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

- You can't declare a variable to be nil, because nil does not have a type. Instead we have a choice, do we want the zero value for a slice?

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

## [5. Project Structure](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_project_structure) <span id="content-5"></span>

### Introduction

- Just like a package, each project should have a clear purpose. If your project is a library, it should provide one thing, say XML parsing, or logging. You should avoid combining multiple purposes into a single project, this will help avoid the dreaded common library.
- TIP: In my experience, the common repo ends up tightly coupled to its biggest consumer and that makes it hard to back-port fixes without upgrading both common and consumer in lock step, bringing in a lot of unrelated changes and API breakage along the way.

### 5.1. Consider fewer, larger packages

- In Go we have only two access modifiers, public and private, indicated by the capitalisation of the first letter of the identifier. If an identifier is public, it’s name starts with a capital letter, that identifier can be referenced by any other Go package.
- Every package, with the exception of cmd/ and internal/, should contain some source code.
- The advice I find myself repeating is to prefer fewer, larger packages. Your default position should be to not create a new package. That will lead to too many types being made public creating a wide, shallow, API surface for your package.
- TIP: If you’re coming from a Java or C# background, consider this rule of thumb. - A Java package is equivalent to a single .go source file. - A Go package is equivalent to a whole Maven module or .NET assembly.
- How do you know when you should break up a .go file into multiple ones? How do you know when you’ve gone to far and should consider consolidating .go file? Here are the guidelines I use:
  - Start each package with one .go file. Give that file the same name as the name of the folder. eg. package http should be placed in a file called http.go in a directory named http.
  - As your package grows you may decide to split apart the various responsibilities into different files. eg, messages.go contains the `Request and Response types, client.go contains the Client type, server.go contains the Server type.
  - If you find your files have similar import declarations, consider combining them. Alternatively, identify the differences between the import sets and move those
  - Different files should be responsible for different areas of the package. messages.go may be responsible for marshalling of HTTP requests and responses on and off the network, http.go may contain the low level network handling logic, client.go and server.go implement the HTTP business logic of request construction or routing, and so on.
  - TIP: Prefer nouns for source file names.
  - NOTE: The Go compiler compiles each package in parallel. Within a package the compiler compiles each function (methods are just fancy functions in Go) in parallel. Changing the layout of your code within a package should not affect compilation time.
- The go tool supports writing your testing package tests in two places. Assuming your package is called http2, you can write a http2_test.go file and use the package http2 declaration. Doing so will compile the code in http2_test.go as if it were part of the http2 package. This is known colloquially as an internal test.
- The go tool also supports a special package declaration, ending in test, ie., package http_test. This allows your test files to live alongside your code in the same package, however when those tests are compiled they are not part of your package’s code, they live in their own package. This allows you to write your tests as if you were another package calling into your code. This is known as an \_external test.
- I recommend using internal tests when writing unit tests for your package. This allows you to test each function or method directly, avoiding the bureaucracy of external testing.
- TIP: Avoid elaborate package hierarchies, resist the desire to apply taxonomy. With one exception, which we’ll talk about next, the hierarchy of Go packages has no meaning to the go tool. For example, the net/http package is not a child or sub-package of the net package.
- If your project contains multiple packages you may find you have some exported functions which are intended to be used by other packages in your project, but are not intended to be part of your project’s public API.
- If you find yourself in this situation the go tool recognises a special folder name—​not package name--, internal/ which can be used to place code which is public to your project, but private to other projects.
- To create such a package, place it in a directory named internal/ or in a sub-directory of a directory named internal/. When the go command sees an import of a package with internal in its path, it verifies that the package doing the import is within the tree rooted at the parent of the internal directory.
- For example, a package …​/a/b/c/internal/d/e/f can be imported only by code in the directory tree rooted at …​/a/b/c. It cannot be imported by code in …​/a/b/g or in any other repository.

### 5.2. Keep package main small as small as possible

- Your main function, and main package should do as little as possible. This is because main.main acts as a singleton; there can only be one main function in a program, including tests.
- Because main.main is a singleton there are a lot of assumptions built into the things that main.main will call that they will only be called during main.main or main.init, and only called once. This makes it hard to write tests for code written in main.main, thus you should aim to move as much of your business logic out of your main function and ideally out of your main package.
- TIP: func main() should parse flags, open connections to databases, loggers, and such, then hand off execution to a high level object.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [6. API Design6. API Design](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_api_design) <span id="content-6"></span>

### Introduction

- However when it comes to reviewing APIs during code review, I am less forgiving. This is because everything I’ve talked about so far can be fixed without breaking backward compatibility; they are, for the most part, implementation details.

### 6.1. Design APIs that are hard to misuse.

> APIs should be easy to use and hard to misuse. — Josh Bloch

- If an API is hard to use for simple things, then every invocation of the API will look complicated. When the actual invocation of the API is complicated it will be less obvious and more likely to be overlooked.
- Be wary of functions which take several parameters of the same type.
- A good example of a simple looking, but hard to use correctly API is one which takes two or more parameters of the same type. Let’s compare two function signatures:
  ```go
  func Max(a, b int) int
  func CopyFile(to, from string) error
  ```
- Max is commutative; the order of its parameters does not matter. The maximum of eight and ten is ten regardless of if I compare eight and ten or ten and eight.
- The same thing can't be applied to CopyFile, since it's not commutative. One solution:

  ```go
  type Source string

  func (src Source) CopyTo(dest string) error {
    return CopyFile(dest, string(src))
  }

  func main() {
    var from Source = "presentation.md"
    from.CopyTo("/tmp/backup")
  }
  ```

- In this way CopyFile is always called correctly—​this can be asserted with a unit test—​and can possibly be made private, further reducing the likelihood of misuse.
- TIP: APIs with multiple parameters of the same type are hard to use correctly.

### 6.2. Design APIs for their default use case

- The gist of this talk was you should design your APIs for the common use case. Said another way, your API should not require the caller to provide parameters which they don’t care about.
- Discourage the use of nil as a parameter.
- Here’s an example from the net/http package

  ```go
  package http

  // ListenAndServe listens on the TCP network address addr and then calls
  // Serve with handler to handle requests on incoming connections.
  // Accepted connections are configured to enable TCP keep-alives.
  //
  // The handler is typically nil, in which case the DefaultServeMux is used.
  //
  // ListenAndServe always returns a non-nil error.
  func ListenAndServe(addr string, handler Handler) error {
  ```

- ListenAndServe takes two parameters, a TCP address to listen for incoming connections, and http.Handler to handle the incoming HTTP request. Serve allows the second parameter to be nil, and notes that usually the caller will pass nil indicating that they want to use http.DefaultServeMux as the implicit parameter.
- How we can use ListenAndServe:
  ```go
  http.ListenAndServe("0.0.0.0:8080", nil)
  http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux)
  ```
- The author of http.ListenAndServe was trying to make the API user’s life easier in the common case, but possibly made the package harder to use safely.
- There is no difference in line count between using DefaultServeMux explicitly, or implicitly via nil.

  ```go
  const root = http.Dir("/htdocs")
  http.Handle("/", http.FileServer(root))
  http.ListenAndServe("0.0.0.0:8080", nil)

  // Verses
  const root = http.Dir("/htdocs")
  http.Handle("/", http.FileServer(root))
  http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux)
  ```

- Give serious consideration to how much time helper functions will save the programmer. Clear is better than concise.
- Avoid public APIs with test only parameters
  Avoid exposing APIs with values who only differ in test scope. Instead, use Public wrappers to hide those parameters, use test scoped helpers to set the property in test scope.
- Prefer var args to []T parameters
- Additionally, because the ids parameter is a slice, you can pass an empty slice or nil to the function and the compiler will be happy. This adds extra testing load because you should cover these cases in your testing.
- We'll have a long if checking:
  ```go
  if svc.MaxConnections > 0 || svc.MaxPendingRequests > 0 || svc.MaxRequests > 0 || svc.MaxRetries > 0 {
    // apply the non zero parameters
  }
  ```
- Instead of using that, just use args:
  ```go
  // anyPostive indicates if any value is greater than zero.
  func anyPositive(values ...int) bool {
    for _, v := range values {
      if v > 0 {
        return true
      }
    }
    return false
  }
  ```
- This enabled me to make the condition where the inner block will be executed clear to the reader:
  ```go
  if anyPositive(svc.MaxConnections, svc.MaxPendingRequests, svc.MaxRequests, svc.MaxRetries) {
          // apply the non zero parameters
  }
  ```
- However there is a problem with anyPositive, someone could accidentally invoke it like this
  ```go
  if anyPositive() { ... }
  ```
- In this case anyPositive would return false because it would execute zero iterations and immediately return false. This isn’t the worst thing in the world — that would be if anyPositive returned true when passed no arguments
- Nevertheless it would be be better if we could change the signature of anyPositive to enforce that the caller should pass at least one argument. We can do that by combining normal and vararg parameters like this:
  ```go
  // anyPostive indicates if any value is greater than zero.
  func anyPositive(first int, rest ...int) bool {
    if first > 0 {
      return true
    }
    for _, v := range rest {
      if v > 0 {
        return true
      }
    }
    return false
  }
  ```

### 6.3. Let functions define the behaviour they requires

- Example of saving a document:
  ```go
  // Save writes the contents of doc to the file f.
  func Save(f *os.File, doc *Document) error
  ```
- Save is also unpleasant to test, because it operates directly with files on disk. So, to verify its operation, the test would have to read the contents of the file after being written.
- Solution:
  ```go
  // Save writes the contents of doc to the supplied
  // ReadWriterCloser.
  func Save(rwc io.ReadWriteCloser, doc *Document) error
  ```
- Using io.ReadWriteCloser we can apply the interface segregation principle to redefine Save to take an interface that describes more general file shaped things.
- With this change, any type that implements the io.ReadWriteCloser interface can be substituted for the previous \*os.File.
- This makes Save both broader in its application, and clarifies to the caller of Save which methods of the \*os.File type are relevant to its operation.
- And as the author of Save I no longer have the option to call those unrelated methods on \*os.File as it is hidden behind the io.ReadWriteCloser interface.
- So we can narrow the specification for the interface we pass to Save to just writing and closing.
  ```go
  // Save writes the contents of doc to the supplied
  // WriteCloser.
  func Save(wc io.WriteCloser, doc *Document) error
  ```
- A better solution would be to redefine Save to take only an io.Writer, stripping it completely of the responsibility to do anything but write data to a stream.
  ```go
  // Save writes the contents of doc to the supplied
  // Writer.
  func Save(w io.Writer, doc *Document) error
  ```
- By applying the interface segregation principle to our Save function, the results has simultaneously been a function which is the most specific in terms of its requirements—​it only needs a thing that is writable—​and the most general in its function, we can now use Save to save our data to anything which implements io.Writer.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [7. Error handling](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_error_handling) <span id="content-7"></span>

### 7.1. Eliminate error handling by eliminating errors

- Let’s write a function to count the number of lines in a file.

  ```go
  func CountLines(r io.Reader) (int, error) {
    var (
      br    = bufio.NewReader(r)
      lines int
      err   error
    )

    for {
      _, err = br.ReadString('\n')
      lines++
      if err != nil {
        break
      }
    }

    if err != io.EOF {
      return 0, err
    }
    return lines, nil
  }
  ```

- We construct a bufio.Reader, and then sit in a loop calling the ReadString method, incrementing a counter until we reach the end of the file, then we return the number of lines read.
- Strange construction:
  ```go
      _, err = br.ReadString('\n')
      lines++
      if err != nil {
        break
      }
  ```
- We increment the count of lines before checking the error—​that looks odd.
- The reason we have to write it this way is ReadString will return an error if it encounters and end-of-file before hitting a newline character. This can happen if there is no final newline in the file.
- To try to fix this, we rearrange the logic to increment the line count, then see if we need to exit the loop.
- An improved version:

  ```go
  func CountLines(r io.Reader) (int, error) {
    sc := bufio.NewScanner(r)
    lines := 0

    for sc.Scan() {
      lines++
    }
    return lines, sc.Err()
  }
  ```

- This improved version switches from using bufio.Reader to bufio.Scanner.
- Under the hood bufio.Scanner uses bufio.Reader, but it adds a nice layer of abstraction which helps remove the error handling with obscured the operation of CountLines.
- The method, sc.Scan() returns true if the scanner has matched a line of text and has not encountered an error. So, the body of our for loop will be called only when there is a line of text in the scanner’s buffer. This means our revised CountLines correctly handles the case where there is no trailing newline, and also handles the case where the file was empty.
- Secondly, as sc.Scan returns false once an error is encountered, our for loop will exit when the end-of-file is reached or an error is encountered. The bufio.Scanner type memoises the first error it encountered and we can recover that error once we’ve exited the loop using the sc.Err() method.
- Lastly, sc.Err() takes care of handling io.EOF and will convert it to a nil if the end of file was reached without encountering another error.
- Earlier in this presentation We’ve seen examples dealing with opening, writing and closing files. The error handling is present, but not overwhelming as the operations can be encapsulated in helpers like ioutil.ReadFile and ioutil.WriteFile. However when dealing with low level network protocols it becomes necessary to build the response directly using I/O primitives the error handling can become repetitive. Consider this fragment of a HTTP server which is constructing the HTTP response.

  ```go
  type Header struct {
    Key, Value string
  }

  type Status struct {
    Code   int
    Reason string
  }

  func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
    _, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
    if err != nil {
      return err
    }

    for _, h := range headers {
      _, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
      if err != nil {
        return err
      }
    }

    if _, err := fmt.Fprint(w, "\r\n"); err != nil {
      return err
    }

    _, err = io.Copy(w, body)
    return err
  }
  ```

- First we construct the status line using fmt.Fprintf, and check the error. Then for each header we write the header key and value, checking the error each time. Lastly we terminate the header section with an additional \r\n, check the error, and copy the response body to the client. Finally, although we don’t need to check the error from io.Copy, we need to translate it from the two return value form that io.Copy returns into the single return value that WriteResponse returns.
- That’s a lot of repetitive work. But we can make it easier on ourselves by introducing a small wrapper type, errWriter.
- errWriter fulfils the io.Writer contract so it can be used to wrap an existing io.Writer. errWriter passes writes through to its underlying writer until an error is detected. From that point on, it discards any writes and returns the previous error.

  ```go
  type errWriter struct {
    io.Writer
    err error
  }

  func (e *errWriter) Write(buf []byte) (int, error) {
    if e.err != nil {
      return 0, e.err
    }
    var n int
    n, e.err = e.Writer.Write(buf)
    return n, nil
  }

  func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
    ew := &errWriter{Writer: w}
    fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

    for _, h := range headers {
      fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
    }

    fmt.Fprint(ew, "\r\n")
    io.Copy(ew, body)
    return ew.err
  }
  ```

- Applying errWriter to WriteResponse dramatically improves the clarity of the code. Each of the operations no longer needs to bracket itself with an error check. Reporting the error is moved to the end of the function by inspecting the ew.err field, avoiding the annoying translation from `io.Copy’s return values.

### 7.2. Only handle an error once

- Lastly, I want to mention that you should only handle errors once. Handling an error means inspecting the error value, and making a single decision.
  ```go
  // WriteAll writes the contents of buf to the supplied writer.
  func WriteAll(w io.Writer, buf []byte) {
          w.Write(buf)
  }
  ```
- But making more than one decision in response to a single error is also problematic. The following is code that I come across frequently.
  ```go
  func WriteAll(w io.Writer, buf []byte) error {
    _, err := w.Write(buf)
    if err != nil {
      log.Println("unable to write:", err) // annotated error goes to log file
      return err                           // unannotated error returned to caller
    }
    return nil
  }
  ```
  ```go
  func WriteConfig(w io.Writer, conf *Config) error {
    buf, err := json.Marshal(conf)
    if err != nil {
      log.Printf("could not marshal config: %v", err)
      return err
    }
    if err := WriteAll(w, buf); err != nil {
      log.Println("could not write config: %v", err)
      return err
    }
    return nil
  }
  ```
- The problem I see a lot is programmers forgetting to return from an error. As we talked about earlier, Go style is to use guard clauses, checking preconditions as the function progresses and returning early.
- Adding context to errors

  ```go
  func WriteConfig(w io.Writer, conf *Config) error {
    buf, err := json.Marshal(conf)
    if err != nil {
      return fmt.Errorf("could not marshal config: %v", err)
    }
    if err := WriteAll(w, buf); err != nil {
      return fmt.Errorf("could not write config: %v", err)
    }
    return nil
  }

  func WriteAll(w io.Writer, buf []byte) error {
    _, err := w.Write(buf)
    if err != nil {
      return fmt.Errorf("write failed: %v", err)
    }
    return nil
  }
  ```

- could not write config: write failed: input/output error
- Wrapping errors with github.com/pkg/errors
- The fmt.Errorf pattern works well for annotating the error message, but it does so at the cost of obscuring the type of the original error. I’ve argued that treating errors as opaque values is important to producing software which is loosely coupled, so the face that the type of the original error should not matter if the only thing you do with an error value is
  - Check that it is not nil.
  - Print or log it.
- Recovering original cause:
  ```go
  func main() {
    _, err := ReadConfig()
    if err != nil {
      fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
      fmt.Printf("stack trace:\n%+v\n", err)
      os.Exit(1)
    }
  }
  ```
  ```go
  original error: *os.PathError open /Users/dfc/.settings.xml: no such file or directory
  stack trace:
  open /Users/dfc/.settings.xml: no such file or directory
  open failed
  main.ReadFile
          /Users/dfc/devel/practical-go/src/errors/readfile2.go:16
  main.ReadConfig
          /Users/dfc/devel/practical-go/src/errors/readfile2.go:29
  main.main
          /Users/dfc/devel/practical-go/src/errors/readfile2.go:35
  runtime.main
          /Users/dfc/go/src/runtime/proc.go:201
  runtime.goexit
          /Users/dfc/go/src/runtime/asm_amd64.s:1333
  could not read config
  ```

**[⬆ back to top](#list-of-contents)**

</br>

---

## [8. Concurrency](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_concurrency) <span id="content-8"></span>

### 8.1. Keep yourself busy or do the work yourself

- What's wrong with this program:

  ```go
  package main

  import (
    "fmt"
    "log"
    "net/http"
  )

  func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintln(w, "Hello, GopherCon SG")
    })
    go func() {
      if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
      }
    }()

    for {
    }
  }
  ```

- The program does what we intended, it serves a simple web server. However it also does something else at the same time, it wastes CPU in an infinite loop. This is because the for{} on the last line of main is going to block the main goroutine because it doesn’t do any IO, wait on a lock, send or receive on a channel, or otherwise communicate with the scheduler.
- As the Go runtime is mostly cooperatively scheduled, this program is going to spin fruitlessly on a single CPU, and may eventually end up live-locked.
- How could we fix this? Here’s one suggestion.

  ```go
  package main

  import (
    "fmt"
    "log"
    "net/http"
    "runtime"
  )

  func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintln(w, "Hello, GopherCon SG")
    })
    go func() {
      if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
      }
    }()

    for {
      runtime.Gosched()
    }
  }
  ```

- Another solution:

  ```go
  package main

  import (
    "fmt"
    "log"
    "net/http"
  )

  func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintln(w, "Hello, GopherCon SG")
    })
    go func() {
      if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
      }
    }()

    select {}
  }
  ```

- An empty select statement will block forever. This is a useful property because now we’re not spinning a whole CPU just to call runtime.GoSched(). However, we’re only treating the symptom, not the cause.
- Best solution (lol)

  ```go
  package main

  import (
    "fmt"
    "log"
    "net/http"
  )

  func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintln(w, "Hello, GopherCon SG")
    })
    if err := http.ListenAndServe(":8080", nil); err != nil {
      log.Fatal(err)
    }
  }
  ```

- So this is my first piece of advice: if your goroutine cannot make progress until it gets the result from another, oftentimes it is simpler to just do the work yourself rather than to delegate it.

### 8.2. Leave concurrency to the caller

- What is the difference between these two APIs?
  ```go
  // ListDirectory returns the contents of dir.
  func ListDirectory(dir string) ([]string, error)
  // ListDirectory returns a channel over which
  // directory entries will be published. When the list
  // of entries is exhausted, the channel will be closed.
  func ListDirectory(dir string) chan string
  ```
- Firstly, the obvious differences; the first example reads a directory into a slice then returns the whole slice, or an error if something went wrong. This happens synchronously, the caller of ListDirectory blocks until all directory entries have been read. Depending on how large the directory, this could take a long time, and could potentially allocate a lot of memory building up the slide of directory entry names.
- Lets look at the second example. This is a little more Go like, ListDirectory returns a channel over which directory entries will be passed. When the channel is closed, that is your indication that there are no more directory entries. As the population of the channel happens after ListDirectory returns, ListDirectory is probably starting a goroutine to populate the channel.
- Its not necessary for the second version to actually use a Go routine; it could allocate a channel sufficient to hold all the directory entries without blocking, fill the channel, close it, then return the channel to the caller. But this is unlikely, as this would have the same problems with consuming a large amount of memory to buffer all the results in a channel.
- The solution to the problems of both implementations is to use a callback, a function that is called in the context of each directory entry as it is executed.
  ```go
  func ListDirectory(dir string, fn func(string))
  ```

### 8.3. Never start a goroutine without knowning when it will stop

- Example:

  ```go
  package main

  import (
    "fmt"
    "net/http"
    _ "net/http/pprof"
  )

  func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
      fmt.Fprintln(resp, "Hello, QCon!")
    })
    go http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux) // debug
    http.ListenAndServe("0.0.0.0:8080", mux)                       // app traffic
  }
  ```

- Solution step 1:

  ```go
  func serveApp() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
      fmt.Fprintln(resp, "Hello, QCon!")
    })
    http.ListenAndServe("0.0.0.0:8080", mux)
  }

  func serveDebug() {
    http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
  }

  func main() {
    go serveDebug()
    serveApp()
  }
  ```

- But there are some operability problems with this program. If serveApp returns then main.main will return causing the program to shutdown and be restarted by whatever process manager you’re using.
- However, serveDebug is run in a separate goroutine and if it returns just that goroutine will exit while the rest of the program continues on. Your operations staff will not be happy to find that they cannot get the statistics out of your application when they want too because the /debug handler stopped working a long time ago.
- What we want to ensure is that if any of the goroutines responsible for serving this application stop, we shut down the application.

  ```go
  func serveApp() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
      fmt.Fprintln(resp, "Hello, QCon!")
    })
    if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
      log.Fatal(err)
    }
  }

  func serveDebug() {
    if err := http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux); err != nil {
      log.Fatal(err)
    }
  }

  func main() {
    go serveDebug()
    go serveApp()
    select {}
  }
  ```

- This approach has a number of problems:
  - If ListenAndServer returns with a nil error, log.Fatal won’t be called and the HTTP service on that port will shut down without stopping the application.
  - log.Fatal calls os.Exit which will unconditionally exit the program; defers won’t be called, other goroutines won’t be notified to shut down, the program will just stop. This makes it difficult to write tests for those functions.
- What we’d really like is to pass any error that occurs back to the originator of the goroutine so that it can know why the goroutine stopped, can shut down the process cleanly.

  ```go
  func serveApp() error {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
      fmt.Fprintln(resp, "Hello, QCon!")
    })
    return http.ListenAndServe("0.0.0.0:8080", mux)
  }

  func serveDebug() error {
    return http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
  }

  func main() {
    done := make(chan error, 2)
    go func() {
      done <- serveDebug()
    }()
    go func() {
      done <- serveApp()
    }()

    for i := 0; i < cap(done); i++ {
      if err := <-done; err != nil {
        fmt.Println("error: %v", err)
      }
    }
  }
  ```

- Now we have a way to wait for each goroutine to exit cleanly and log any error they encounter. All that is needed is a way to forward the shutdown signal from the first goroutine that exits to the others.
- It turns out that asking a http.Server to shut down is a little involved, so I’ve spun that logic out into a helper function. The serve helper takes an address and http.Handler, similar to http.ListenAndServe, and also a stop channel which we use to trigger the Shutdown method.

  ```go
  func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
    s := http.Server{
      Addr:    addr,
      Handler: handler,
    }

    go func() {
      <-stop // wait for stop signal
      s.Shutdown(context.Background())
    }()

    return s.ListenAndServe()
  }

  func serveApp(stop <-chan struct{}) error {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
      fmt.Fprintln(resp, "Hello, QCon!")
    })
    return serve("0.0.0.0:8080", mux, stop)
  }

  func serveDebug(stop <-chan struct{}) error {
    return serve("127.0.0.1:8001", http.DefaultServeMux, stop)
  }

  func main() {
    done := make(chan error, 2)
    stop := make(chan struct{})
    go func() {
      done <- serveDebug(stop)
    }()
    go func() {
      done <- serveApp(stop)
    }()

    var stopped bool
    for i := 0; i < cap(done); i++ {
      if err := <-done; err != nil {
        fmt.Println("error: %v", err)
      }
      if !stopped {
        stopped = true
        close(stop)
      }
    }
  }
  ```

**[⬆ back to top](#list-of-contents)**

</br>

---

## References:

- https://dave.cheney.net/practical-go/presentations/qcon-china.html
