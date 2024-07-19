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

// PaymentGateway routes to the corresponding payment method.
type PaymentGateway interface {
	ProcessPayment(amount float64) error
}

// PayPalGateway represents the PayPal configuration service.
type PayPalGateway struct{}

// ProcessPayment process the payment for PayPal service.
func (pg *PayPalGateway) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	// Simulate PayPal payment processing logic.
	return nil
}

// StripeGateway represents the Stripe configuration service.
type StripeGateway struct{}

// ProcessPayment process the payment for Stripe service.
func (sg *StripeGateway) ProcessPayment(amount float64) error {
	fmt.Printf("Processing Stripe payment of $%.2f\n", amount)
	// Simulate Stripe payment processing logic.
	return nil
}

// NewPaymentGateway sets the payment type.
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
