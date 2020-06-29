package handlers

import (
	"context"
	"net/http"
	"fmt"
	"Building-Microservices-With-go/product-api/data"
)

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

	    err := prod.FromJSON(r.Body)
	    if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(
				rw, 
				fmt.Sprintf("ERROR vaildating product: %s", err), 
				http.StatusBadRequest,
			)
			return
		}
		
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}