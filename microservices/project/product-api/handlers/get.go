package handlers

import (
	"net/http"

	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/data"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/utils"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")
	p.l.Println("[DEBUG] fetching all products")

	prods := data.GetProducts()

	err := utils.ToJSON(prods, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}

	return nil
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {

	case nil:
		// With no error

	case data.ErrProductNotFound:
		// Let the user know the product couldn't be found.
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&GenericError{Message: err.Error()}, rw)
		return

	default:
		// unexpected error
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = utils.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
