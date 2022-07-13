package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/data"
)

// KeyProduct is the key used to store a product in the context
type KeyProduct struct{}

// Products handler
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts returns a new products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProduct is the error returned when the product is not valid
var ErrInvalidProduct = fmt.Errorf("invalid path, must be /products/{id}")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Message []string `json:"message"`
}

// getProductID from the URL path
// if cannot convert the id into an integer will panic
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id to an integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}
