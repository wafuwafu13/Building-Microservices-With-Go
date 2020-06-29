// Package classification Petstore API.
//
// Documentation for Product API
//
// Terms Of Service:
//
//     Schemes: http
//     BasePath: /
//     Versiton: 1.0.0
//    
//     Consumes:
//     - application/json
//    
//     Produces:
//     - application/json
// swagger:meta
package handlers

import (
	"Building-Microservices-With-go/product-api/data"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in system
	// in:body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// swagger:response noContent
type productsNoContent struct {

}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}
