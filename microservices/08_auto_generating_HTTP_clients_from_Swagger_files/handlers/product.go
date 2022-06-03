package handlers

import (
	"08_auto_generating_http_clients_from_swagger_files/data"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyProduct is a key used fo the Product object in the context.
type KeyProduct struct{}

// Products handler for getting and updating products.
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts returns a new products handler with the given Logger and Validation.
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid.
var ErrInavlidProductPath = fmt.Errorf("invalid Path, path shoud be /products/{id}")

// GenericError is a generic error message returned by a server.
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages.
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the id in the url path.
// Panics if cannot convert the id into an integer
// (this should never happen as the router ensures this is a valid number)
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// conver the id into int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}
