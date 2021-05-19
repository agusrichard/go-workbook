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
//     SecurityDefinitions:
//     Bearer:
//       type: apiKey
//       name: Authorization
//       in: header
//
// swagger:meta
package docs

import (
	"twit/models/requests"
	"twit/models/responses"
)

// ================================ REGISTER ================================

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

// ================================ LOGIN ================================

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

// ================================ GET PROFILE ================================

// swagger:route GET /user/profile Authentication getProfile
// Get Profile.
// responses:
//   200: getProfileResponse
// Security:
//   Bearer: []

// This text will appear as description of your response body.
// swagger:response getProfileResponse
type getProfileResponseWrapper struct {
	// in:body
	Body responses.LoginUserResponse
}
