package responses

import "twit/models"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginData struct {
	AccessToken string      `json:"access-token"`
	User        models.User `json:"user"`
}

type LoginUserResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    LoginData `json:"data"`
}
