# Learning Notes for How to Test RESTAPI Application

## Notes:

**Go mod**:
- We don't have to declare the module path belonging to a repository.
- A module can be defined locally without belonging to a repository.
- `go install` command builds the module, producing an executable binary.
- The binaries are install to the bin subdirectory of the default GOPATH.
- The easiest way to make your module available for others to use is usually to make its module path match the URL for the repository.
- `go mod tidy` command adds missing module requirements for imported packages.
- `go clean -modcache`: remove all downloaded modules.

**Improving Your Go Tests and Mocks With Testify**
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

**Unit Test vs Integration Test: What's the Difference?**
- Unit test: Test the unit of code (component).
- Integration test: Individual units of a program are combined and tested as a group.
- Unit Testing tests only the functionality of the units themselves and may not catch integration errors, or other system-wide issues
- Integrating testing may detect errors when modules are integrated to build the overall system
- Unit test does not verify whether your code works with external dependencies correctly.
- Integration tests verify that your code works with external dependencies correctly.
- White Box Testing is software testing technique in which internal structure, design and coding of software are tested to verify flow of input-output and to improve design, usability and security
- Black Box Testing is a software testing method in which the functionalities of software applications are tested without having knowledge of internal code structure, implementation details and internal paths. Black Box Testing mainly focuses on input and output of software applications. (also known as behavioral testing)

**Unit Testing for REST APIs in Go**
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

**Develop REST API using Go and Test using various methods**
- REST is not a standard but a set of recommendations and constraints for RESTful web services.
    - Client-Server: SystemA makes an HTTP request to a URL hosted by SystemB, which returns a response.
    - Stateless: REST is stateless: the client request should contain all the information necessary to respond to a request.
    - Cacheable: A response should be defined as cacheable or not.
    - Layered: The requesting client need not know whether it’s communicating with the actual server, a proxy, or any other intermediary.
    - Uniform interface – REST is defined by four interface constraints: identification of resources; manipulation of resources through representations; self-descriptive messages; and, hypermedia as the engine of application state.
    - Code on demand (optional) – REST allows client functionality to be extended by downloading and executing code in the form of applets or scripts. This simplifies clients by reducing the number of features required to be pre-implemented.

**Structuring Tests in Go**
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

## References:
- https://golang.org/doc/code
- https://tutorialedge.net/golang/improving-your-tests-with-testify-go/
- https://www.guru99.com/unit-test-vs-integration-test.html#:~:text=Unit%20Testing%20test%20each%20part,see%20they%20are%20working%20fine.&text=Unit%20Testing%20is%20executed%20by,performed%20by%20the%20testing%20team.
- https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d
- https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c