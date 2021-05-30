# Project Structure in GO

## Contents:

### [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

Good Structure Goals:
- Consistent
- Easy to understand, navigate, reason about (**make sense**)
- Easy to change
- Loosely coupled
- Easy to test
- As simple as possible, but no simplistic
- Design reflects exactly how the software works
- Structure reflects the design exactly

> Good code structure will make our life easier as a programmer

#### Flat Structure
![img.png](images/img.png)

If our app is simple, start with flat structure. But everything is global, this is certainly a drawback.

#### Group by function (layered architecture)

- Presentation/User interface
- Business logic
- External dependencies/infrastructures

Cons:
- We will easily be trapped in circular imports

![img_1.png](images/img_1.png)

#### Group by module

![img_2.png](images/img_2.png)

Cons:
- Funny naming
- Naming clash

#### Group by context (Domain Driven Development)

- Establish domain and business logic
- Define bounded contexts, limit our code into its own context
- Categorising the building blocks

![img_3.png](images/img_3.png)

![img_4.png](images/img_4.png)

![img_5.png](images/img_5.png)

![img_6.png](images/img_6.png)

This is probably an overkill for small app
> Don't start with a complex structure!

#### Testing
- Keep the _test.go files next to the main files
- Use a shared mock subpackage

#### Naming
- Choose package name to communicate what it provides rather than to what it contains
- Avoid generic names like util
- Follow the usual go convections
- Avoid stutter (e.g strings.Reader not strings.StringReader)

#### What is in the main file
keep main file short, it should only do one thing (e.g running the server)
![img_7.png](images/img_7.png)

---

### [Go Project Structure Best Practices](https://tutorialedge.net/golang/go-project-structure-best-practices/)

Some people in Go community follow the well known [golang-standards/project-layout](https://github.com/golang-standards/project-layout). But after the introduction of go modules, this pattern starts to present challenges. 

#### Small Applications (Flat Structure)
Small applications are better to use flat structure, no need to overkill the structure by going complex at the start.
This structure has pros of simplicity. Developers don't have to think too hard to put which file on which folder.
They just have to focus on create the app itself and deliver its funcionalities.

Benefits:
- Suitable for microservices
- Suitable for small tools and libraries

#### Medium/Large Size Applications - Modularization

![img_8.png](images/img_8.png)

### [How do I Structure my Go Project?](https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/)

This pattern usually used for big project when there are several main.go files with its own binary.

- /cmd </br>
  Contains the main application entry point files
- /internal </br>
  Holds the private library code used in the service
- /pkg
  Contains code which is OK for other services to consume, this may include API clients, or utility functions which may be handy for other projects
  
Goals we should consider:
- Keep things consistent
- Kepp things as simple as possible, but no simpler
- Loosely coupled sections of the service or application
- Aim to ensure it is easy to navigate our way around


## References
- https://www.youtube.com/watch?v=oL6JBUk6tj0
- https://tutorialedge.net/golang/go-project-structure-best-practices/
- https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/