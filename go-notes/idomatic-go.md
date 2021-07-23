# Idiomatic Go

</br>

## List of Contents:
### 1. [The Zen of Go](#content-1)


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


## References:
- https://dave.cheney.net/2020/02/23/the-zen-of-go