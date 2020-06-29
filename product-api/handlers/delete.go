package handlers

import (
	"net/http"
	"Building-Microservices-With-Go/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
//    201: noContentResponse
//  404: errorResponse
//  501: errorResponse

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getProductID(r)

	p.l.Println("Handle Delete Product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}