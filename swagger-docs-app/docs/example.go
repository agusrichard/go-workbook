// Package classification Learn Swagger.
//
// Documentation of our Learn Swagger API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package docs

import "swagger-docs-app/api"

// swagger:route GET /hello hello-tag idOfHelloEndpoint
// Hello does some amazing stuff.
// responses:
//   200: helloResponse

// This text will appear as description of your response body.
// swagger:response helloResponse
type helloResponseWrapper struct {
	// in:body
	Body api.HelloResponse
}
