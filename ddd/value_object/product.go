package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Product
type Product struct {
	Name string
	Cost *Money
	l    *log.Logger
}

// NewProduct creates a new product
func NewProduct(name string, cost *Money, l *log.Logger) (*Product, error) {
	if name == "" {
		return nil, fmt.Errorf("name must not be empty")
	}
	return &Product{Name: name, Cost: cost, l: l}, nil
}

// AddCost adds cost to Money of the product
func (p *Product) AddCost(amount int) error {
	newCost, err := p.Cost.AddAmount(amount)
	if err != nil {
		p.l.Fatal(err)
		return err
	}
	p.Cost = newCost
	return nil
}
