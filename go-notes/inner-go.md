# Inner Go

</br>

## List of Contents:
### 1. [Go mod](#content-1)
### 2. [What “accept interfaces, return structs” means in Go](#content-2)
### 3. [Preemptive Interface Anti-Pattern in Go](#content-3)
### 4. [Context in Golang!](#content-4)


</br>

---

## Contents:

## [Go mod](https://golang.org/doc/code) <span id="content-1"></span>

</br>

- We don't have to declare the module path belonging to a repository.
- A module can be defined locally without belonging to a repository.
- `go install` command builds the module, producing an executable binary.
- The binaries are install to the bin subdirectory of the default GOPATH.
- The easiest way to make your module available for others to use is usually to make its module path match the URL for the repository.
- `go mod tidy` command adds missing module requirements for imported packages.
- `go clean -modcache`: remove all downloaded modules.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [What “accept interfaces, return structs” means in Go](https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8) <span id="content-2"></span>


- > All problems in computer science can be solved by another level of indirection, except of course for the problem of too many indirections
  > >David J. Wheeler
- Interfaces abstract away from structures in Go
- Tt doesn’t make sense to create this complexity until it’s needed
- > Always [abstract] things when you actually need them, never when you just foresee that you need them.
- You can control the return values of a function, but you can't control the input type.
- That's why it's better for us to accept interface, instead of concrete types.
- Another aspect of simplification is removing unnecessary detail.
- If you don't need some recipes to make something, then don't list it on your need-list.
- Check this snippet:
```go
type Database struct{ }
func (d *Database) AddUser(s string) {...}
func (d *Database) RemoveUser(s string) {...}
func NewUser(d *Database, firstName string, lastName string) {
  d.AddUser(firstName + lastName)
}
```
- On the above code, we define database to have 2 methods. But on NewUser database job is just to add new user. No need to add RemoveUser
- This is probably the better way:
```go
type DatabaseWriter interface {
  AddUser(string)
}
func NewUser(d DatabaseWriter, firstName string, lastName string) {
  d.AddUser(firstName + lastName)
}
```

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Preemptive Interface Anti-Pattern in Go](https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a) <span id="content-3"></span>
- Interfaces are a way to describe behavior
- Preempttive interfaces are when developer codes to an interface before an actual need arises.
- Example:
```go
type Auth interface {
  GetUser() (User, error)
}
type authImpl struct {
  // ...
}
func NewAuth() Auth {
  return &authImpl
}
```
- You have to change the code if you use this
```java
// auth.java
public class Auth {
  public boolean canAction() {
    // ...
  }
}
// logic.java
public class Logic {
  public void takeAction(Auth a) {
    // ...
  }
}
```
- For example, you want to take any objects in takeAction as long as it has canAction inside it. How it would be?
- Better code in java
```java
// auth.java
public interface Auth {
  public boolean canAction()
}
// authimpl.java
class AuthImpl implements Auth {
}
// logic.java
public class Logic {
  public void takeAction(Auth a) {
    // ...
  }
}
```
- (Personal notes) It'e better to pass a pointer of a struct rather than the struct itself. It makes sure that we check for its nullity.
- Go uses implicit interface, which means concrete objects (structs) don't need to explicitly defined that they are using this interface. It's different from explicit interface like in Java.
- Usually you don't need preemptive interface in go.
- Go is at its most powerful when interface definitions are small.
- In the standard library, most interface definitions are a single method
- Accepting interfaces gives your API the greatest flexibility and returning structs allows the people reading your code to quickly navigate to the correct function
- Unnecessary abstraction creates unnecessary complication. Don’t over complicate code until it’s needed.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Context in Golang!](https://levelup.gitconnected.com/context-in-golang-98908f042a57) <span id="content-4"></span>


### Introduction
- Applications in golang use Contexts for controlling and managing very important aspects of reliable applications, such as cancellation and data sharing in concurrent programming.

### Context with value
- One of the most common uses for contexts is to share data, or use request scoped values. When you have multiple functions and you want to share data between them, you can do so using contexts.
- The easiest way to do that is to use the function `context.WithValue`.
- You can think about the internal implementation as if the context contained a map inside of it, so you can add and retrieve values by key.
- Example:
  ```go
  package main

  import (
  	"context"
  	"fmt"
  )

  func main() {
  	ctx := context.Background()
  	ctx = addValue(ctx)
  	readValue(ctx)
  }

  func addValue(ctx context.Context) context.Context {
  	return context.WithValue(ctx, "key", "test-value")
  }

  func readValue(ctx context.Context) {
  	val := ctx.Value("key")
  	fmt.Println(val)
  }
  ```
- One important aspect of the design behind context package is that everything returns a new `context.Context` struct.
- Using this technique you can pass along the context.Context to concurrent functions and as long as you properly manage the context you are passing on, it’s good way to share scoped values between those concurrent functions (meaning that each context will keep their own values on its scope).

### Middlewares
- The type http.Request contains a context which can carry scoped values throughout the HTTP pipeline.
- Example:
  ```go
  package main

  import (
  	"context"
  	"log"
  	"net/http"

  	"github.com/google/uuid"
  	"github.com/gorilla/mux"
  )

  func main() {
  	router := mux.NewRouter()
  	router.Use(guidMiddleware)
  	router.HandleFunc("/ishealthy", handleIsHealthy).Methods(http.MethodGet)
  	http.ListenAndServe(":8080", router)
  }

  func handleIsHealthy(w http.ResponseWriter, r *http.Request) {
  	w.WriteHeader(http.StatusOK)
  	uuid := r.Context().Value("uuid")
  	log.Printf("[%v] Returning 200 - Healthy", uuid)
  	w.Write([]byte("Healthy"))
  }

  func guidMiddleware(next http.Handler) http.Handler {
  	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  		uuid := uuid.New()
  		r = r.WithContext(context.WithValue(r.Context(), "uuid", uuid))
  		next.ServeHTTP(w, r)
  	})
  }
  ```

### Context Cancellation
- It’s a good practice to propagate the cancellation signal when you receive one. 
- Let’s say you have a function where you start tens of goroutines. That main function waits for all goroutines to finish or a cancellation signal before proceeding. If you receive the cancellation signal you may want to propagate it to all your goroutines, so you don’t waste compute resources. If you share the same context among all goroutines you can easily do that.
- To create a context with cancellation you only have to call the function context.WithCancel(ctx) passing your context as parameter. This will return a new context and a cancel function. To cancel that context you only need to call the cancel function.
- Example:
  ```go
  package main

  import (
  	"context"
  	"fmt"
  	"io/ioutil"
  	"net/http"
  	neturl "net/url"
  	"time"
  )

  func queryWithHedgedRequestsWithContext(urls []string) string {
  	ch := make(chan string, len(urls))
  	ctx, cancel := context.WithCancel(context.Background())
  	defer cancel()
  	for _, url := range urls {
  		go func(u string, c chan string) {
  			c <- executeQueryWithContext(u, ctx)
  		}(url, ch)

  		select {
  		case r := <-ch:
  			cancel()
  			return r
  		case <-time.After(21 * time.Millisecond):
  		}
  	}

  	return <-ch
  }

  func executeQueryWithContext(url string, ctx context.Context) string {
  	start := time.Now()
  	parsedURL, _ := neturl.Parse(url)
  	req := &http.Request{URL: parsedURL}
  	req = req.WithContext(ctx)

  	response, err := http.DefaultClient.Do(req)

  	if err != nil {
  		fmt.Println(err.Error())
  		return err.Error()
  	}

  	defer response.Body.Close()
  	body, _ := ioutil.ReadAll(response.Body)
  	fmt.Printf("Request time: %d ms from url%s\n", time.Since(start).Nanoseconds()/time.Millisecond.Nanoseconds(), url)
  	return fmt.Sprintf("%s from %s", body, url)
  }
  ```
- Each request is fired in a separate goroutine. The context is passed to all requests that are fired. The only thing that is being done with the context is that it gets propagated to the HTTP client. This allows a graceful cancellation of the request and underlying connection when the cancel function is called.
- This is a very common patter for functions that accept a context.Context as argument, they either actively act on the context (like checking if it was cancelled) or they pass it to an underlying function that deals with it (in this case the Do function that receives the context through the http.Request).


### Context Timeout
- It is a good practice to always defer the cancel function when it is available to avoid context leaking.
- Syntax:
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
  ```




**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8
- https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a
- https://levelup.gitconnected.com/context-in-golang-98908f042a57