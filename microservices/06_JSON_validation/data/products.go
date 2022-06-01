package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"github.com/go-playground/validator"
)

// Structure fo an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func validateSKU(fl validator.FieldLevel) bool {
	// SKU is of format abc-abcd-abcde
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	// Exist more than one match or no match.
	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("SKU", validateSKU)

	return validate.Struct(p)
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
		SKU:         "abc123",
		CreatedOn:   "2020-01-01",
		UpdatedOn:   "2020-01-01",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee",
		Price:       3.00,
		SKU:         "def456",
		CreatedOn:   "2020-01-01",
		UpdatedOn:   "2020-01-01",
		DeletedOn:   "",
	},
}
