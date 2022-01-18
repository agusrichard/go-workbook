# Practical Go

</br>

## [List of Contents:](#list-of-contents)
### 1. [Guiding principles](#content-1)


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
- In summary:
  - When declaring a variable without initialisation, use the var syntax.
  - When declaring and explicitly initialising a variable, use :=.

> My advice in this situation is to follow the local style.


**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://dave.cheney.net/practical-go/presentations/qcon-china.html