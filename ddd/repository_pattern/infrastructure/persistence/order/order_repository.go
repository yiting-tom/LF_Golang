package order

import (
	"database/sql"

	"github.com/yiting-tom/LF_Golang/ddd/repository_pattern/domain/order"
)

// SqlOrderRepository implements the aOrderRepository interface
type SqlOrderRepository struct {
    db *sql.DB
    // and other fields
}

// NewSqlOrderRepository returns a new SqlOrderRepository
func NewSqlOrderRepository(db *sql.DB) *SqlOrderRepository {
    return &SqlOrderRepository{
        db: db,
        // and other fields
    }
}

// GetById returns the order with the given order.OrderId
func (r *SqlOrderRepository) GetById(id order.OrderId) (*order.Order, error) {
    // do what should be done to get the order
    return nil, nil
}

// Save saves the order
func (r *SqlOrderRepository) Save(order *order.Order) error {
    // do what should be done to save the order
    return nil
}

// Remove removes the Order with the given order.OrderId
func (r *SqlOrderRepository) Remove(id order.OrderId) error {
    // do what should be done to remove the order
    return nil
}