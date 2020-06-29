package handlers

import (
	"net/http"
	"Building-Microservices-With-Go/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse
// ListAll handles GET requests and returns all current products

func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	rw.Header().Add("Content-Type", "application/json")

	prod := data.GetProducts()
	err := data.ToJSON(prod, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
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

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)
		
		rw.WriteHeader(http.StatusNotFound)
		return
	
	default:
		p.l.Println("[ERROR] fetching product", err)
		
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}