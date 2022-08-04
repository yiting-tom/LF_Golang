package main

import "fmt"

// Money is a value object
type Money struct {
	Amount   int
	Currency string
}

// NewMoney creates a new Money
func NewMoney(amount int, currency string) (*Money, error) {
	if amount < 0 {
		return nil, fmt.Errorf("amount must be positive")
	}
	return &Money{Amount: amount, Currency: currency}, nil
}

// AddAmount adds amount to Money, and returns a new Money
func (m *Money) AddAmount(amount int) (*Money, error) {
	if amount < 0 {
		return nil, fmt.Errorf("amount must be positive")
	}
	newAmount := m.Amount + amount
	return NewMoney(newAmount, m.Currency)
}