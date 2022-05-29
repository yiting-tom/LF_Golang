package handlers

import (
	"06_json_validation/data"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products is a http.Handler.
type Products struct {
	l *log.Logger
}

// Creates a products handler with the given logger.
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Return the products from the mock data.
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// Fetch the products from the db.
	productList := data.GetProducts()

	// Serialize the list to JSON.
	if err := productList.ToJSON(w); err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// Add a new product into mock data.
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	// Get the new product from the context of middleware.
	newProduct := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Println("Get new Product:", newProduct)

	// Add the newProduct to the mock data.
	data.AddProduct(newProduct)
}

// Update a product with specific id.
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the params from the request.
	vars := mux.Vars(r)

	// Get the id from the params.
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to parse id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Products:", id)

	// Get the new product from the context of middleware.
	newProduct := r.Context().Value(KeyProduct{}).(*data.Product)

	// Update the product with the specific id.
	if err := data.UpdateProduct(id, newProduct); err != nil {

		if err == data.ErrProductNotFound { // If the product is not found, return 404.
			http.Error(w, "Product not found", http.StatusNotFound)

		} else { // If any other error, return 500.
			http.Error(w, "Unable to update product", http.StatusInternalServerError)
		}

		return
	}
}

type KeyProduct struct {
}

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new &Product.
		newProduct := &data.Product{}

		// Try to parse this request's body to the newProduct.
		if err := newProduct.FromJSON(r.Body); err != nil {
			http.Error(w, "Unable to parse json", http.StatusBadRequest)
			return
		}

		// Create a new context with value {key: KeyProduct{}, val: newProduct}.
		ctx := context.WithValue(r.Context(), KeyProduct{}, newProduct)

		// Attach the new context to the request.
		r = r.WithContext(ctx)

		// Call the next handler,
		// which can be another middleware in this chain,
		// or the final handler.
		next.ServeHTTP(w, r)
	})
}
