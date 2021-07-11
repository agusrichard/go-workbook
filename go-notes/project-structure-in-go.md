# Project Structure in Go

</br>

## List of Contents:
### 1. [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](#content-1)
### 2. [Go Project Structure Best Practices](#content-2)
### 3. [How do I Structure my Go Project?](#content-3)
### 4. [Structuring Applications in Go](#content-4)

</br>

---

## Contents:

### [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) <span id="content-1"></span>

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

</br>

---

### [Go Project Structure Best Practices](https://tutorialedge.net/golang/go-project-structure-best-practices/) <span id="content-2"></span>

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

</br>

---

### [How do I Structure my Go Project?](https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/) <span id="content-3"></span>

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

</br>

---

## [Structuring Applications in Go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091) <span id="content-4"></span>
- Don't use global variables
- You may decide to add a global database connection or a global configuration variable but these globals are a nightmare to use when writing unit tests
- Snippet
```go
type HelloHandler struct {
    db *sql.DB
}
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var name string
    // Execute the query.
    row := h.db.QueryRow(“SELECT myname FROM mytable”)
    if err := row.Scan(&name); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    // Write it back to the client.
    fmt.Fprintf(w, “hi %s!\n”, name)
}
```
- This is how to use wrapper https://gist.github.com/tsenart/5fc18c659814c078378d
```go
func helloHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    		var name string
    		// Execute the query.
    		row := db.QueryRow("SELECT myname FROM mytable")
    		if err := row.Scan(&name); err != nil {
        		http.Error(w, err.Error(), 500)
        		return
    		}
    		// Write it back to the client.
    		fmt.Fprintf(w, "hi %s!\n", name)
    	})
}

func withMetrics(l *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		next.ServeHTTP(w, r)
		l.Printf("%s %s took %s", r.Method, r.URL, time.Since(began))
	})
}
```
- Separate your binary from your application
- Library driven development
- Wrap types for application-specific context
```go
package myapp
import (
    "database/sql"
)
type DB struct {
    *sql.DB
}
type Tx struct {
    *sql.Tx
}
```
```go
// Open returns a DB reference for a data source.
func Open(dataSourceName string) (*DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    return &DB{db}, nil
}
// Begin starts an returns a new transaction.
func (db *DB) Begin() (*Tx, error) {
    tx, err := db.DB.Begin()
    if err != nil {
        return nil, err
    }
    return &Tx{tx}, nil
}
```
- Don’t go crazy with subpackages
- Using a single root package
- Organize the most important type at the top of the file and add types in decreasing importance towards the bottom of the file.
- If you’re writing Go projects the same way you write Ruby, Java, or Node.js projects then you’re probably going to be fighting with the language.

</br>

---

## References
- https://www.youtube.com/watch?v=oL6JBUk6tj0
- https://tutorialedge.net/golang/go-project-structure-best-practices/
- https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/
- https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091
- https://levelup.gitconnected.com/a-practical-approach-to-structuring-go-applications-7f77d7f9c189 (UNREAD)