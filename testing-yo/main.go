package main

import "fmt"

type MyInt struct {
	number int
}

type MyIntI interface {
	Useless() int
}

func (m *MyInt) Useless() int {
	return m.number
}

func UsingUseless(m MyIntI) int {
	return m.Useless()
}

type notification struct {}

type Notification interface {
	SendPaymentNotification(amount int) string
}

func (n *notification) SendPaymentNotification(amount int) string {
	return fmt.Sprintf("you have made the payment with an amount of %d", amount)
}

type payment struct {
	notification Notification
}

type Payment interface {
	MakePayment(amount int) bool
}

func (p *payment) MakePayment(amount int) bool {
	fmt.Printf("make payment with an amount of %d\n", amount)
	fmt.Println(p.notification.SendPaymentNotification(amount))
	return true
}

func main() {
	fmt.Println("Hello World")
}
