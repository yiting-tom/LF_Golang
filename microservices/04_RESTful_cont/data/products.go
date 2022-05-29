package data

import (
	"encoding/json"
	"fmt"
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

// Proeucts is a collection of Product
type Products []*Product

// Define the product not found Error.
var ErrProductNotFound = fmt.Errorf("Product not found")

// Return the next available ID.
func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

// Find a Product in productList by id.
func findProductById(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

// Serialize the list to JSON.
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

// Deconde the JSON to Product.
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)

	return e.Decode(p)
}

// Return the mock of collection of products.
func GetProducts() Products {
	return productList
}

// Add a new Product into productList.
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// Update a Product in productList.
func UpdateProduct(id int, p *Product) error {
	if prevProduct, pos, err := findProductById(id); err != nil {
		return err
	} else {
		p.ID = prevProduct.ID
		productList[pos] = p
	}

	return nil
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
