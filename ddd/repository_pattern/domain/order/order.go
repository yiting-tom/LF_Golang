package order

import "github.com/google/uuid"

// OrderId is the unique identifier of an Order and it should be a Value Object
type OrderId struct {
	uuid.UUID
    // and other fields
}

// Order is an Entity
type Order struct {
    id OrderId
    // and other fields
}

// OrderRepository is the interface for the OrderRepository
type OrderRepository interface {
    GetById(id OrderId) (*Order, error)
    Save(order *Order) error
    Remove(id OrderId) error
}