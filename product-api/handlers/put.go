package handlers

import (
	"net/http"
	"Building-Microservices-With-go/product-api/data"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

func (p *Products) UpdateProducts(rw http.ResponseWriter, r * http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product) // 型アサーション
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}