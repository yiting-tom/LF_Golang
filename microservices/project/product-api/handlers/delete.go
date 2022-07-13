package handlers

import (
	"net/http"

	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/data"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/utils"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a product by ID
//
// responses:
// 200: noContentResponse
// 404: errorResponse
// 501: errorResponse

// Delete handles DELETE requests and removes a product from the database
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting product: ", id)

	err := data.DeleteProduct(id)
	switch err {
	case data.ErrProductNotFound:
		p.l.Println("[ERROR] deleting product", id, "not found")

		w.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&GenericError{Message: "Product not found"}, w)

	case nil:
		p.l.Println("[ERROR] deleting product", id, "failed: ", err)

		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericError{Message: err.Error()}, w)
		return nil
	}

	// unknown error
	w.WriteHeader(http.StatusNoContent)

	return nil
}
