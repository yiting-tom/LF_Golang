package handlers

import (
	"09_cors/data"
	"context"
	"net/http"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new &Product.
		prod := &data.Product{}

		// Try to parse this request's body to the prod.
		if err := data.FromJSON(prod, r.Body); err != nil {
			p.l.Println("[ERROR] Unable to parse json:", err)

			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		// Validate the prod.
		if errs := p.v.Validate(prod); len(errs) != 0 {
			p.l.Println("[ERROR] Validation failed:", errs)

			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return
		}

		// Create a new context with value {key: KeyProduct{}, val: newProduct}.
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		// Attach the new context to the request.
		r = r.WithContext(ctx)

		// Call the next handler,
		// which can be another middleware in this chain,
		// or the final handler.
		next.ServeHTTP(w, r)
	})
}

// MiddlewareAddContentType adds the Content-Type header to the response
func (p *Products) MiddlewareAddContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
