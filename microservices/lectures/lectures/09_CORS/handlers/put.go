package handlers

import (
	"09_cors/data"
	"net/http"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
//  200: productResponse
//  404: errorResponse
//  422: errorResponse

// Update handles PUT requests to update an existing product
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
	// fetch the product from the request body
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] update product id", prod.ID)

	if err := data.UpdateProduct(prod); err != nil {
		p.l.Println("[ERROR] product not found", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in db"}, w)
		return
	}

	// write the no content success header
	w.WriteHeader(http.StatusNoContent)
}
