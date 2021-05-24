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