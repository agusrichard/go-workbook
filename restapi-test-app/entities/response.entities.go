package entities

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type AppError struct {
	Err error
	StatusCode int
}

type AppResult struct {
	Data interface{}
	Message string
	Err error
	StatusCode int
}