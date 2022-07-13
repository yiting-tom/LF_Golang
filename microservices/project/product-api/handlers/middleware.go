package handlers

import (
	"context"
	"net/http"

	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/data"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-api/utils"
)

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		p_new := &data.Product{}

		err := utils.FromJSON(p_new, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product: ", err)

			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		errs := p.v.Validate(p_new)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product: ", errs)

			w.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&ValidationError{Message: errs.Errors()}, w)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, p_new)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another
		// middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
