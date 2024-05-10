package main

import (
	"errors"
	"fmt"
)

type PaymentGatewayType int

const (
	PayPalGateway PaymentGatewayType = iota
	StripeGateway
)

type PaymentGateway interface {
	ProcessPayment(amount float64) error
}

type PayPalGateway struct{}

func (pg *PayPalGateway) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	// Simulate PayPal payment processing logic.
	return nil
}

type StripeGateway struct{}

func (sg *StripeGateway) ProcessPayment(amount float64) error {
	fmt.Printf("Processing Stripe payment of $%.2f\n", amount)
	// Simulate Stripe payment processing logic.
	return nil
}

func NewPaymentGateway(gwType PaymentGatewayType) (PaymentGateway, error) {
	switch gwType {
	case PayPalGateway:
		return &PayPalGateway{}, nil
	case StripeGateway:
		return &StripeGateway{}, nil
	default:
		return nil, errors.New("unsupported payment gateway type")
	}
}

func main() {
	payPalGateway, _ := NewPaymentGateway(PayPalGateway)
	payPalGateway.ProcessPayment(100.00)
	stripeGateway, _ := NewPaymentGateway(StripeGateway)
	stripeGateway.ProcessPayment(150.50)
}
