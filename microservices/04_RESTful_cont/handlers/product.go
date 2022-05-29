package handlers

import (
	"04_restful/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	// handle the request for a list of products.
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// handle the request for create a new product.
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		re := regexp.MustCompile(`/([0-9]+)`)
		group := re.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		if id, err := strconv.Atoi(idString); err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		} else {
			p.l.Println("Got id:", id)
			p.updateProduct(id, w, r)
		}
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

// Add a new product into mock data.
func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handler POST Products")

	// Create a new &Product.
	newProduct := &data.Product{}

	// Parse the JSON from request's Body to the newProduct.
	if err := newProduct.FromJSON(r.Body); err != nil {
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	// Add the newProduct to the mock data.
	data.AddProduct(newProduct)
}

// Update a product with specific id.
func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handler PUT Products")

	// Create a new &Product.
	newProduct := &data.Product{}

	// Parse the JSON from request's Body to the newProduct.
	if err := newProduct.FromJSON(r.Body); err != nil {
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

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
