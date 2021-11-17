# Inner Go

</br>

## List of Contents:
### 1. [Go mod](#content-1)
### 2. [What “accept interfaces, return structs” means in Go](#content-2)
### 3. [Preemptive Interface Anti-Pattern in Go](#content-3)
### 4. [Context in Golang!](#content-4)
### 5. [Interfaces in Golang](#content-5)
### 6. [A Guide On SQL Database Transactions In Go](#content-6)
### 7. [5 Useful Go Tricks and Tips You Should Know 🐹 🚀](#content-7)
### 8. [Try and Catch in Golang](#content-8)


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

## [Interfaces in Golang](https://medium.com/nerd-for-tech/interfaces-in-golang-f9df59b0b71e) <span id="content-5"></span>

### Why are we not defining interfaces with the type definition like in typical static languages like C++, Java etc ?
- In languages like C++, Java, one needs to specify that a type implements an interface like in the code given below:
  ```java
	interface Vehicle 
	{   
		// public and abstract  
		void drive(); 
	} 

	// A class that implements the interface. 
	class Car implements Vehicle 
	{ 
		// Implementing the capabilities of 
		// interface. 
		public void drive() 
		{ 
			System.out.println("Drive"); 
		} 
	}
  ```
- In such languages, defining the interface for the Object enables the compiler to form a dispatch tables for the objects pointing to the functions.
- For a developer, this means that the object implementing the interface does not need to explicitly say it implements it, as shown in the code below:
  ```go
	type Vehicle interface {
		Drive()
	}

	// Car implements the Vehicle interface simply by implementing method Drive.
	type Car struct{}

	func (c Car) Drive() {
		fmt.Println("Drive")
	}
  ```
- Thus any struct can satisfy an interface simply by implementing its method signatures. It offers several advantages like:
  - Makes it easier to use mocks instead of real objects in unit tests.
  - Helps enforce decoupling between parts of your codebase.

### Should this package export interface in combination with the exposed type implementing the interface?
- Don’t export any interfaces unless you have to.
- Take a case of golang error interface.
	```go
  type error interface {
      Error() string
  }
  ```
- It is a builtin interface in the standard library that standardises the error behaviour.

### Should this package return an interface rather than the concrete type?
- According to CodeReviewComments, Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values.
- However Effective go docs also complements it by saying that: if a type exists only to implement an interface and will never have exported methods beyond that interface, there is no need to export the type itself.
- But the question is how do you identify such scenarios? How do you know that the type will have no additional value in the future? In my experience, the answer is to “wait”.
- Do not start off by returning interfaces, but wait till your code evolves and you see the need for them. As Rob Pike says: “Don’t design with interfaces, discover them.”
- A good hint for exposing an interface is when you have multiple types in your package implementing the same method signature.

### In case of confusion, it is helpful to look for some red flags that can signals that you’re probably using interfaces wrong. Some are:
- Your interface is not decoupling an API from change.
  - Example:
    ```go
  package sendgrid

  type SendGrid interface {
    SendValidationEmail(email string) error
    SendPasswordChangeEmail(email string) error
  }

  // sendgrid is our Sendgrid implementation.
  type sendgrid struct {
    /* impl */
  }

  func NewSendGrid(host string) SendGrid {
    return &sendgrid{host}
  }

  func (s *sendgrid) SendValidationEmail(email string) error {
      /* impl */
  }

  func (s *sendgrid) SendPasswordChangeEmail(email string) error {
      /* impl */
  }
  ```
- Also imagine that you want to add a new method to this interface that’s used by lots of people, how do you add a new method to it without breaking their code?
  ```go
  package sendgrid

  // Remove the interface and change the concrete type to public
  type SendGrid struct {
    /* impl */
  }

  // Change the NewSendGrid to return pointer to SendGrid instead.

  func NewSendGrid(host string) *SendGrid {
    return &sendgrid{host}
  }

  func (s *sendgrid) SendValidationEmail(email string) error {
      /* impl */
  }

  func (s *sendgrid) SendPasswordChangeEmail(email string) error {
      /* impl */
  }
  ```
- Your interface has more than 1 or 2 methods.
  - Having too many methods for your interface reduces its usability. Taking example of fmt.Stringer interface, it has only one method signature, i.e
    ```go
  type Stringer interface {
      String() string
  }
  ```
- The bigger the interface, the weaker the abstraction.

**[⬆ back to top](#list-of-contents)**

</br>

---

## [A Guide On SQL Database Transactions In Go](https://www.sohamkamani.com/golang/sql-transactions/) <span id="content-6"></span>

### Introduction
- Transactions are very useful when you want to perform multiple operations on a database, but still treat them as a single unit.

### Transactions
- For example, what if someone adopted pets and bought food for them? We could then write two queries to do just that:
  ```sql
  INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat');
  INSERT INTO food (name, quantity) VALUES ('Dog Biscuit', 3), ('Cat Food', 5);
  ```
- Now lets think about what happens if the first query succeeds, but the second query fails: you now have data which shows that two new pets are adopted, but no food has been bought.
  - We cannot treat this as a success, since one of the queries failed
  - We cannot treat this as a failure, since the first query passed. If we consider this a failure and retry, we would have to insert Fido and Albert for a second time.
- To avoid situations like this, we want both the queries to pass or fail together. This is where SQL transactions come in.
- Executing queries:
  ```sql
  BEGIN;
  INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat');
  INSERT INTO food (name, quantity) VALUES ('Dog Biscuit', 3), ('Cat Food', 5);
  END;
  ```
- The BEGIN statement starts a new transaction
- Once the transaction has begun, SQL statements are executed one after the other, although they don’t reflect in the database just yet.
- The END statement commits the above transactions atomically
- Incase we want to abort the transaction in the middle, we could have used the ROLLBACK statement
- Here, “atomically” means both of the SQL statements are treated as a single unit - they pass or fail together

### Implementing Transactions in Go - Basic Transactions
- Example:
  ```go
  package main

  import (
    "context"
    "database/sql"
    "log"

    _ "github.com/lib/pq"
  )

  func main() {
    // Create a new connection to our database
    connStr := "user=soham dbname=pet_shop sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
      log.Fatal(err)
    }

    // Create a new context, and begin a transaction
    ctx := context.Background()
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
      log.Fatal(err)
    }
    // `tx` is an instance of `*sql.Tx` through which we can execute our queries

    // Here, the query is executed on the transaction instance, and not applied to the database yet
    _, err = tx.ExecContext(ctx, "INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat')")
    if err != nil {
      // Incase we find any error in the query execution, rollback the transaction
      tx.Rollback()
      return
    }

    // The next query is handled similarly
    _, err = tx.ExecContext(ctx, "INSERT INTO food (name, quantity) VALUES ('Dog Biscuit', 3), ('Cat Food', 5)")
    if err != nil {
      tx.Rollback()
      return
    }

    // Finally, if no errors are recieved from the queries, commit the transaction
    // this applies the above changes to our database
    err = tx.Commit()
    if err != nil {
      log.Fatal(err)
    }
  }
  ```

### Read-and-Update Transactions
- Example:
  ```go
  package main

  import (
    "context"
    "database/sql"
    "log"

    _ "github.com/lib/pq"
  )

  func main() {
    // Initialize a connection, and begin a transaction like before
    connStr := "user=soham dbname=pet_shop sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
      log.Fatal(err)
    }

    ctx := context.Background()
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
      log.Fatal(err)
    }

    _, err = tx.ExecContext(ctx, "INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat')")
    if err != nil {
      tx.Rollback()
      return
    }

    // Run a query to get a count of all cats
    row := tx.QueryRow("SELECT count(*) FROM pets WHERE species='cat'")
    var catCount int
    // Store the count in the `catCount` variable
    err = row.Scan(&catCount)
    if err != nil {
      tx.Rollback()
      return
    }

    // Now update the food table, increasing the quantity of cat food by 10x the number of cats
    _, err = tx.ExecContext(ctx, "UPDATE food SET quantity=quantity+$1 WHERE name='Cat Food'", 10*catCount)
    if err != nil {
      tx.Rollback()
      return
    }

    // Commit the change if all queries ran successfully
    err = tx.Commit()
    if err != nil {
      log.Fatal(err)
    }
  }
  ```
- It’s important to note why the read query is executed within the transaction: any read query outside the transaction doesn’t consider the values of an uncommitted transaction.
- This means that if our read query was outside the transaction, we would not consider the pets added in the first insert query.


**[⬆ back to top](#list-of-contents)**

</br>

---

## [5 Useful Go Tricks and Tips You Should Know 🐹 🚀](https://cgarciarosales97.medium.com/5-useful-go-tricks-and-tips-you-should-know-b8017d1f1833) <span id="content-7"></span>

### 1. Execution time of a code
- Snippet:
  ```go
  package main

  import (
    "log"
    "time"
  )

  func ExecTime(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
  }

  func main() {
    defer ExecTime(time.Now(), "main")
    time.Sleep(3 * time.Second)
  }
  ```
### 2. Marshall Json string with no defined Struct to a “Object”
- Snippet:
  ```go
  package main

  import (
    "encoding/json"
    "fmt"
    "time"
  )

  type LogRegister struct {
    ID            int
    RequestString string
    RequestJson   map[string]interface{}
    CreatedAt     time.Time
    UpdatedAt     time.Time
  }

  func (c *LogRegister) FormatStringToJson() {
    var request map[string]interface{}

    json.Unmarshal([]byte(c.RequestString), &request)

    c.RequestJson = request
  }

  func main() {
    logRegister := LogRegister{
      ID: 1,
      RequestString: `
      {
        "status":200,
        "message":"ok, now you can use me"
      }
      `,
      CreatedAt: time.Now(),
    }
    logRegister.FormatStringToJson()
    if logRegister.RequestJson["status"] != nil {
      fmt.Println(logRegister.RequestJson["status"])
    }
    if logRegister.RequestJson["message"] != nil {
      fmt.Println(logRegister.RequestJson["message"])
    }
  }
  ```

### 3. Get Test Coverage of your code
- Command:
  ```shell
  go test -coverprofile=coverage.out ./… && go tool cover -html=coverage.out && rm coverage.out
  ```

### 4. Search if a String is inside of a Slice
- Snippet:
  ```go
  package main

  func StringInSlice(a string, list []string) bool {
    for _, b := range list {
      if b == a {
        return true
      }
    }
    return false
  }

  func main() {
    permmitedFields := []string{"status", "message"}
    field := "createdAt"

    if StringInSlice(field, permmitedFields) {
      println("the field is authorized")
    }
    println("the field is not authorized")
  }
  ```

### 5. CPU Profiling
- Command:
  ```shell
  go test -cpuprofile=cpu.out ./… && go tool pprof cpu.out && rm cpu.out
  ```

## [Try and Catch in Golang](https://dzone.com/articles/try-and-catch-in-golang) <span id="content-8"></span>

- Snippet:
  ```go
  Block{
          Try: func() {
              fmt.Println("I tried")
              Throw("Oh,...sh...")
          },
          Catch: func(e Exception) {
              fmt.Printf("Caught %v\n", e)
          },
          Finally: func() {
              fmt.Println("Finally...")
          },
      }.Do()
  ```
- Snippet:
  ```go
  package main
  
  import (
      "fmt"
  )
  
  type Block struct {
      Try     func()
      Catch   func(Exception)
      Finally func()
  }
  
  type Exception interface{}
  
  func Throw(up Exception) {
      panic(up)
  }
  
  func (tcf Block) Do() {
      if tcf.Finally != nil {
  
          defer tcf.Finally()
      }
      if tcf.Catch != nil {
          defer func() {
              if r := recover(); r != nil {
                  tcf.Catch(r)
              }
          }()
      }
      tcf.Try()
  }
  
  func main() {
      fmt.Println("We started")
      Block{
          Try: func() {
              fmt.Println("I tried")
              Throw("Oh,...sh...")
          },
          Catch: func(e Exception) {
              fmt.Printf("Caught %v\n", e)
          },
          Finally: func() {
              fmt.Println("Finally...")
          },
      }.Do()
      fmt.Println("We went on")
  }
  ```

**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8
- https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a
- https://levelup.gitconnected.com/context-in-golang-98908f042a57
- https://medium.com/nerd-for-tech/interfaces-in-golang-f9df59b0b71e
- https://www.sohamkamani.com/golang/sql-transactions/
- https://cgarciarosales97.medium.com/5-useful-go-tricks-and-tips-you-should-know-b8017d1f1833