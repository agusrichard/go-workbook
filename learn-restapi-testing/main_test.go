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