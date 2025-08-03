package main

import (
	"errors"
	"fmt"
)

type PaymentInterface interface {
	Pay(amount float64) error
}

type CreditCard struct {
	CardNumber string
}

func (c *CreditCard) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	fmt.Printf("Paid $%.2f using Credit Card ****%s\n",
		amount, c.CardNumber[len(c.CardNumber)-4:])
	return nil
}

type UPI struct {
	UPIId string
}

func (u *UPI) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	fmt.Printf("Paid $%.2f using UPI ID: %s\n", amount, u.UPIId)
	return nil
}

type PaymentContext struct {
	strategy PaymentInterface
}

func NewPaymentContext() *PaymentContext {
	return &PaymentContext{}
}

func (p *PaymentContext) SetStrategy(strategy PaymentInterface) error {
	if strategy == nil {
		return errors.New("strategy cannot be nil")
	}
	p.strategy = strategy
	return nil
}

func (p *PaymentContext) Pay(amount float64) error {
	if p.strategy == nil {
		return errors.New("no payment strategy set")
	}
	return p.strategy.Pay(amount)
}

func main() {
	payment := NewPaymentContext()

	// Try payment without strategy
	err := payment.Pay(100.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Credit Card payment
	cc := &CreditCard{CardNumber: "1234567890123456"}
	payment.SetStrategy(cc)
	payment.Pay(150.50)

	// UPI payment
	upi := &UPI{UPIId: "user@paytm"}
	payment.SetStrategy(upi)
	payment.Pay(75.25)
}
