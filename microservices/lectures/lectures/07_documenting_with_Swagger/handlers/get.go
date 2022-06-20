package handlers

import (
	"07_documenting_with_swagger/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Return a list of products from the db.
// responses:
//  200: productsResponse

// ListAll handles GET requests and returns all list of products
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all products")

	ps := data.GetProducts()

	if err := data.ToJSON(ps, w); err != nil {
		p.l.Println("[ERROR] serializing product", err)
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a product from the db.
// responses:
//  200: productsResponse
//  404: errorResponse

// ListSingle handles GET requests and returns a single product
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get product", id)

	target, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", id, err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return

	default:
		p.l.Println("[ERROR] fetching product", id, err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	if err = data.ToJSON(target, w); err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
