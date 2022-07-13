package handlers

import (
	"net/http"

	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/data"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/utils"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
