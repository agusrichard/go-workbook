# Testing in Go

</br>

## List of Contents:
### 1. [Improving Your Go Tests and Mocks With Testify](#content-1)
### 2. [Unit Test vs Integration Test: What's the Difference?](#content-2)
### 3. [Unit Testing for REST APIs in Go](#content-3)
### 4. [Structuring Tests in Go](#content-4)
### 5. [Setup and Teardown using Go testing package](#content-5)
### 6. [Another patterns of testing in Go](#content-6)
### 7. [Go: tests with HTML coverage report](#content-7)


</br>

---

## Contents:

## [Improving Your Go Tests and Mocks With Testify](https://tutorialedge.net/golang/improving-your-tests-with-testify-go/) <span id="content-1"></span>

</br>

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

</br>

---


## [Unit Test vs Integration Test: What's the Difference?](https://www.guru99.com/unit-test-vs-integration-test.html#:~:text=Unit%20Testing%20test%20each%20part,see%20they%20are%20working%20fine.&text=Unit%20Testing%20is%20executed%20by,performed%20by%20the%20testing%20team.) <span id="content-2"></span>


- Unit test: Test the unit of code (component).
- Integration test: Individual units of a program are combined and tested as a group.
- Unit Testing tests only the functionality of the units themselves and may not catch integration errors, or other system-wide issues
- Integrating testing may detect errors when modules are integrated to build the overall system
- Unit test does not verify whether your code works with external dependencies correctly.
- Integration tests verify that your code works with external dependencies correctly.
- White Box Testing is software testing technique in which internal structure, design and coding of software are tested to verify flow of input-output and to improve design, usability and security
- Black Box Testing is a software testing method in which the functionalities of software applications are tested without having knowledge of internal code structure, implementation details and internal paths. Black Box Testing mainly focuses on input and output of software applications. (also known as behavioral testing)

</br>

---

## [Unit Testing for REST APIs in Go](https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d) <span id="content-3"></span>

</br >

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

</br>

---

## [Structuring Tests in Go](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c) <span id="content-4"></span>


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

</br>

---

## [Setup and Teardown using Go testing package](https://golang.org/pkg/testing/) <span id="content-5"></span>

You can set a setUp and tearDown function.

A setUp function prepares your environment to tests.
A tearDown function does a rollback.

```go
package main

import (
	"fmt"
	"os"
	"testing"
)

var number int

func setup() {
	number = 10
	fmt.Println("setup the test")
}

func teardown() {
	number = 99
	fmt.Println("teardown the test")
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func TestCalculate1(t *testing.T) {
	expected := 3
	actual := Calculate(1)
	if actual != expected {
		t.Error(fmt.Sprintf("expect %v got %v", expected, actual))
	}
}

func TestIntermediate(t *testing.T) {
	number += 10
	fmt.Println("here")
}

func TestCalculate2(t *testing.T) {
	fmt.Println("number", number)
	expected := 4
	actual := Calculate(2)
	if actual != expected {
		t.Error(fmt.Sprintf("expect %v got %v", expected, actual))
	}
}

func BenchmarkCalculate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(1)
	}
}
```

</br>

---

## [Another patterns of testing in Go](https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go) <span id="content-6"></span>

We can use this approach when we want to implement table driven tests

```shell
go test -run ''      # Run all tests.
go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
```

```go
// main_test.go
// running parallel tests
func TestGroupedParallel(t *testing.T) {
    for _, tc := range tests {
    tc := tc // capture range variable
        t.Run(tc.Name, func(t *testing.T) {
            t.Parallel()
            ...
        })
    }
}
```

The race detector kills the program if it exceeds 8192 concurrent goroutines, so use care when running parallel tests with the -race flag set.

```go
// main_test.go
package main

import (
	"fmt"
	"testing"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test")
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}

func TestCalculate(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{ 1, 3},
		{ 2, 4},
		{ 3, 5},
		{ 4, 6},
		{ 5, 7},
	}

	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range cases {
		t.Run(fmt.Sprintf("running test when n=%d", tc.n), func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)

			actual := Calculate(tc.n)
			if actual != tc.expected {
				t.Fatalf("expect %v got %v", tc.expected, actual)
			}
		})
	}
}

```

</br>

---

## [Go: tests with HTML coverage report](https://kenanbek.medium.com/go-tests-with-html-coverage-report-f977da09552d) <span id="content-6"></span>

To add cover reports `go test -cover ./...`

To get detail information of coverage
```shell
go test -coverprofile=coverage.out ./...  # save coverage results
go tool cover -func=coverage.out          # print results
go tool cover -html=coverage.out          # view coverage as html page
```

The coverage tool also includes three different coverage modes. You can select coverage mode by using -covermode option:

```shell
go test -covermode=count -coverprofile=coverage.out
```

Makefile for simplicity

```makefile
GO=go
GOCOVER=$(GO) tool cover
.PHONY: test/cover
test/cover:
    $(GOTEST) -v -coverprofile=coverage.out ./...
    $(GOCOVER) -func=coverage.out
    $(GOCOVER) -html=coverage.out
```

There are three different cover modes:
- set: did each statement run?
- count: how many times did each statement run?
- atomic: like count, but counts precisely in parallel programs

</br>

---

## References:
- https://golang.org/doc/code
- https://tutorialedge.net/golang/improving-your-tests-with-testify-go/
- https://www.guru99.com/unit-test-vs-integration-test.html#:~:text=Unit%20Testing%20test%20each%20part,see%20they%20are%20working%20fine.&text=Unit%20Testing%20is%20executed%20by,performed%20by%20the%20testing%20team.
- https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d
- https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c
- https://golang.org/pkg/testing/
- https://riptutorial.com/go/example/15183/testing-using-setup-and-teardown-function
- https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go
- https://kenanbek.medium.com/go-tests-with-html-coverage-report-f977da09552d