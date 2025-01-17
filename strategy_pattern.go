package main

import "fmt"

type PaymentInterface interface {
	Pay()
}

type CreditCard struct{}

func (c *CreditCard) Pay() {
	fmt.Println("Payment using CreditCard")
}

type UPI struct{}

func (c *UPI) Pay() {
	fmt.Println("Payment using UPI")
}

// Maintains a reference to a strategy object and allows dynamically changing the strategy
type PaymentContext struct {
	payment PaymentInterface
}

func NewPayment() PaymentContext {
	return PaymentContext{}
}

func (p *PaymentContext) SetStrategy(paymentMode PaymentInterface) {
	p.payment = paymentMode
}

func (p *PaymentContext) Pay() {
	if p.payment == nil {
		fmt.Println("No payment strategy set")
		return
	}
	p.payment.Pay()
}

func main() {
	payment := NewPayment()
	payment.Pay()

	cc := &CreditCard{}
	upi := &UPI{}

	payment.SetStrategy(cc)
	payment.Pay()

	payment.SetStrategy(upi)
	payment.Pay()
}
