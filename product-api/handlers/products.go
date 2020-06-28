package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"Building-Microservices-With-go/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost { // curl localhost:9090 -d '{}'| jq
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut { // curl localhost:9090/1 -XPUT -d '{"name": "tea", "description": "a nice cup of tea"}'| jq
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		// p.l.Println(g) [[/1 1]]
		// p.l.Println(g[0]) [/1 1]

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		// p.l.Println("got id", id)
		// curl localhost:9090/1 -XPUT
		// got id 1

		p.updateProducts(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	// p.l.Printf("Prod: %#v", prod) 
	// curl localhost:9090 -d '{"id": 1, "name": "tea", "description": "nice"}'| jq 
	// Prod: &data.Product{ID:1, Name:"tea", Description:"nice", Price:0, SKU:"", CreatedOn:"", UpdatedOn:"", DeletedOn:""} 
	
	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r * http.Request) {
	p.l.Println("Handle Put Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}