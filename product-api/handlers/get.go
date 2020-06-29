package handlers

import (
	"net/http"
	"Building-Microservices-With-go/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200: productsResponse

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}