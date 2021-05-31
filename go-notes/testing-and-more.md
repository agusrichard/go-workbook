# Learning Notes for How to Test RESTAPI Application

## Notes:

**[Go mod](https://golang.org/doc/code)**:
- We don't have to declare the module path belonging to a repository.
- A module can be defined locally without belonging to a repository.
- `go install` command builds the module, producing an executable binary.
- The binaries are install to the bin subdirectory of the default GOPATH.
- The easiest way to make your module available for others to use is usually to make its module path match the URL for the repository.
- `go mod tidy` command adds missing module requirements for imported packages.
- `go clean -modcache`: remove all downloaded modules.

**[Improving Your Go Tests and Mocks With Testify](https://tutorialedge.net/golang/improving-your-tests-with-testify-go/)**
- Assertion example on `main_test.go`
```go
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, Multiply(5, 3), 15)
}
```
- Mocking example
```go
// main.go
package main

import "fmt"

type messageService struct {}

type MessageService interface {
	SendChargeNotification(value int) error
}

func InitializeMessageService() MessageService {
	return &messageService{}
}

func (service *messageService) SendChargeNotification(value int) error {
	fmt.Println("Send notification!")
	return nil
}

type myService struct {
	messageService MessageService
}

type MyService interface {
	ChargeCustomer(value int) error
}

func InitializeMyService(service MessageService) MyService {
	fmt.Println("service", service)
	return &myService{service}
}

func (service *myService) ChargeCustomer(value int) error {
	err := service.messageService.SendChargeNotification(value)
	fmt.Printf("Charge customer %d\n", value)
	return err
}

func main() {
	fmt.Println("Hello")
	serviceOne := InitializeMessageService()
	serviceTwo := InitializeMyService(serviceOne)
	serviceTwo.ChargeCustomer(100)
}
```
```go
//main_test.go
package main

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type smsServiceMockT struct {
	mock.Mock
}

func (m *smsServiceMockT) SendChargeNotification(value int) error {
	fmt.Println("Send notification!")
	args := m.Called(value)
	return args.Error(0)
}

func TestMyService_ChargeCustomer(t *testing.T) {
	serviceOne := new(smsServiceMockT)
	serviceTwo := InitializeMyService(serviceOne)

	serviceOne.On("SendChargeNotification", 100).Return(nil)
	serviceTwo.ChargeCustomer(100)

	serviceOne.AssertExpectations(t)
}
```

**[Unit Test vs Integration Test: What's the Difference?](https://www.guru99.com/unit-test-vs-integration-test.html#:~:text=Unit%20Testing%20test%20each%20part,see%20they%20are%20working%20fine.&text=Unit%20Testing%20is%20executed%20by,performed%20by%20the%20testing%20team.)**
- Unit test: Test the unit of code (component).
- Integration test: Individual units of a program are combined and tested as a group.
- Unit Testing tests only the functionality of the units themselves and may not catch integration errors, or other system-wide issues
- Integrating testing may detect errors when modules are integrated to build the overall system
- Unit test does not verify whether your code works with external dependencies correctly.
- Integration tests verify that your code works with external dependencies correctly.
- White Box Testing is software testing technique in which internal structure, design and coding of software are tested to verify flow of input-output and to improve design, usability and security
- Black Box Testing is a software testing method in which the functionalities of software applications are tested without having knowledge of internal code structure, implementation details and internal paths. Black Box Testing mainly focuses on input and output of software applications. (also known as behavioral testing)

**[Unit Testing for REST APIs in Go](https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d)**
- Snippet 1
```go
func TestGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/entries", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEntries)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb@g.com","phone_number":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"FirstNameSample","last_name":"LastNameSample","email_address":"lr@gmail.com","phone_number":"1111111111"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
```
- Snippet 2
```go
func TestGetEntryByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEntryByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
```

**[Structuring Tests in Go](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c)**
- Tests should be two things: self-contained and easily reproducible
- Self-contained means changing one part of our test suite does not drastically affect another part.
- Reproducible means someone doesn’t have to go through multiple steps to get their test suite running the same as mine.
- Go has a perfectly good testing framework built in. Frameworks are also one more barrier to entry for other developers contributing to your code.
- In the same folder, you'll have a file and file_test, you can named the package myapp and myapp_test. Then use dot import.
```go
package myapp_test
import (
    "testing"
    . "github.com/benbjohnson/myapp"
)
func TestUser_Save(t *testing.T) {
    u := &User{Name: "Susy Queue"}
    ok(t, u.Save())
}
```
- Interface can make our code complex, but also makes it difficult to test.
- Use inline interfaces & simple mocks.
```go
package yo
type Client struct {}
// Send sends a "yo" to someone.
func (c *Client) Send(recipient string) error
// Yos retrieves a list of my yo's.
func (c *Client) Yos() ([]*Yo, error)
```
```go
package myapp
type MyApplication struct {
    YoClient interface {
        Send(string) error
    }
}
func (a *MyApplication) Yo(recipient string) error {
    return a.YoClient.Send(recipient)
}
```
```go
package main
func main() {
    c := yo.NewClient()
    a := myapp.MyApplication{}
    a.YoClient = c
    ...
}
```
- The caller should create the interface instead of the callee providing an interface.
```go
package myapp_test
// TestYoClient provides mockable implementation of yo.Client.
type TestYoClient struct {
    SendFunc func(string) error
}
func (c *TestYoClient) Send(recipient string) error {
    return c.SendFunc(recipient)
}
func TestMyApplication_SendYo(t *testing.T) {
    c := &TestYoClient{}
    a := &MyApplication{YoClient: c}
    // Mock our send function to capture the argument.
    var recipient string
    c.SendFunc = func(s string) error {
        recipient = s
        return nil
    }
    // Send the yo and verify the recipient.
    err := a.Yo("susy")
    ok(t, err)
    equals(t, "susy", recipient)
}

```

**[What “accept interfaces, return structs” means in Go](https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8)**
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

**[Preemptive Interface Anti-Pattern in Go](https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a)**
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

**[Structuring Applications in Go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)**
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

**[Error handling and Go](https://blog.golang.org/error-handling-and-go)**
- Standard usage:
```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```
- If you want to create custom error 1
```go
// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```
- Error message that probably you want to consider 
```go
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // implementation
}
```
- Formatting error message, this returns error type
```go
fmt.Errorf("math: square root of negative number %g", f)
```
- Another example of custom error
```go
type SyntaxError struct {
    msg    string // description of error
    Offset int64  // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string { return e.msg }
```
- Simplify repetitive error handling
```go
func init() {
    http.HandleFunc("/view", viewRecord)
}

func viewRecord(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    if err := viewTemplate.Execute(w, record); err != nil {
        http.Error(w, err.Error(), 500)
    }
}

type appHandler func(http.ResponseWriter, *http.Request) error

func viewRecord(w http.ResponseWriter, r *http.Request) error {
  c := appengine.NewContext(r)
  key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
  record := new(Record)
  if err := datastore.Get(c, key, record); err != nil {
      return err
  }
  return viewTemplate.Execute(w, record)
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if err := fn(w, r); err != nil {
    http.Error(w, err.Error(), 500)
  }
}
```

## References:
- https://golang.org/doc/code
- https://tutorialedge.net/golang/improving-your-tests-with-testify-go/
- https://www.guru99.com/unit-test-vs-integration-test.html#:~:text=Unit%20Testing%20test%20each%20part,see%20they%20are%20working%20fine.&text=Unit%20Testing%20is%20executed%20by,performed%20by%20the%20testing%20team.
- https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d
- https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c
- https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8
- https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a
- https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091
- https://blog.golang.org/error-handling-and-go