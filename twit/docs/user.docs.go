// Package classification Twit Application.
//
// Documentation of our Twit Application.
//
//     Schemes: http, https
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:9090
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

import (
	"twit/models/requests"
	"twit/models/responses"
)

// swagger:route POST /auth/register Authentication registerUser
// Register User.
// responses:
//   200: registerResponse

// This text will appear as description of your response body.
// swagger:response registerResponse
type registerUserResponseWrapper struct {
	// in:body
	Body responses.Response
}

// swagger:parameters registerUser
type registerParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body requests.RegisterUserRequest
}

// swagger:route POST /auth/login Authentication loginUser
// Login User.
// responses:
//   200: loginResponse

// This text will appear as description of your response body.
// swagger:response loginResponse
type loginUserResponseWrapper struct {
	// in:body
	Body responses.LoginUserResponse
}

// swagger:parameters loginUser
type loginParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body requests.LoginUserRequest
}
