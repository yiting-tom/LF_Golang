package data

import (
	"fmt"
)

// ErrProductNotFound is an error
// raiswd when a product cant not be found in the db.
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure fo an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	//  min: 1
	ID int `json:"id"`

	// the name for the product
	//
	// required: true
	// min length: 255
	Name string `json:"name" validate:"required"`

	// the description for the product
	//
	// required: false
	// max length: 10240
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0
	Price float32 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// min length: 1
	SKU string `json:"sku" validate:"required,skuValidate"`
}

// Proeucts defines a slice of Product
type Products []*Product

// GetProducts return all products from the database.
func GetProducts() Products {
	return productList
}

// GetProductByID return a product with specifc ID from db .
// If not found return ErrProductNotFound error.
func GetProductByID(id int) (*Product, error) {
	idx := findIndexByProductID(id)

	if id == -1 {
		return nil, ErrProductNotFound
	}
	return productList[idx], nil
}

// AddProduct adds a new product into the db.
func AddProduct(p Product) {
	p.ID = getNextID()
	// update the product in the db.
	productList = append(productList, &p)
}

// UpdateProduct replaces a product in the db with the given product.
// If no product with the given id exists, an ErrProductNotFound error is returned.
func UpdateProduct(p Product) error {
	idx := findIndexByProductID(p.ID)
	if idx == -1 {
		return ErrProductNotFound
	}

	// update the product in the db.
	productList[idx] = &p
	return nil
}

// DeleteProduct replaces a product in the db with the given product.
// If no product with the given id exists, an ErrProductNotFound error is returned.
func DeleteProduct(id int) error {
	idx := findIndexByProductID(id)
	if idx == -1 {
		return ErrProductNotFound
	}

	// update the product in the db.
	productList[idx] = productList[len(productList)-1] // copy last ele to idx
	productList = productList[:len(productList)-1]     // truncate slice

	return nil
}

// findIndexByProductID finds the index of product in the database.
// If not found return -1.
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

// getNextID returns the next id for the product for adding to db.
func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

// productList is a mock list of products
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       2.50,
		SKU:         "aaa-aaa-aaa",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee",
		Price:       3.00,
		SKU:         "bbb-bbb-bbb",
	},
}
