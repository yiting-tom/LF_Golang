package handlers

import (
	"03_restful/data"
	"log"
	"net/http"
)

// Products is a http.Handler.
type Products struct {
	l *log.Logger
}

// Creates a products handler with the given logger.
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// interface
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// If no method is satisfied return an error
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// Return the products from the mock data.
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fFetch the products from the db.
	productList := data.GetProducts()

	// Serialize the list to JSON.
	if err := productList.ToJSON(w); err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
