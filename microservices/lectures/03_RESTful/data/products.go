package data

import (
	"encoding/json"
	"io"
)

// Structure fo an API product
type Product struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// Return a list of products
func GetProducts() Products {
	return productList
}

// productList is a mock list of products
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       2.50,
		CreatedOn:   "2020-01-01",
		UpdatedOn:   "2020-01-01",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee",
		Price:       3.00,
		CreatedOn:   "2020-01-01",
		UpdatedOn:   "2020-01-01",
		DeletedOn:   "",
	},
}
