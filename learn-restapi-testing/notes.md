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
```bigquery
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
```bigquery
-- main.go
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
```bigquery
-- main_test.go
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

References:
- https://golang.org/doc/code
- https://tutorialedge.net/golang/improving-your-tests-with-testify-go/