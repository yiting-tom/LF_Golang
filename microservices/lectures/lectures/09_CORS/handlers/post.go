package handlers

import (
	"09_cors/data"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new proudct
//
// responses:
//  200: productResponse
//  422: errorResponse
//  501: errorResponse

// Create handles POST reuqests to create a new product
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	// fetch the product from the request body
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Println("[DEBUG] create product", prod)
	data.AddProduct(prod)
}
