package mockyo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockedObject struct {
	mock.Mock
}

func (m *MockedObject) DoAmazingStuff(input int) (bool, error) {
	args := m.Called(input)
	return args.Bool(0), args.Error(1)
}

func WhatIsThisFunction(m *MockedObject) {
	fmt.Println("WhatIsThisFunction")
	m.DoAmazingStuff(123)
}

func TestAmazingStuff(t *testing.T) {
	testObj := new(MockedObject)
	testObj.On("DoAmazingStuff", 123).Return(true, nil)
	WhatIsThisFunction(testObj)
	testObj.AssertExpectations(t)
}
