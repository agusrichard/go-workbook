package models

import "fmt"

type RequestError struct {
	StatusCode int64
	Err        error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf(r.Err.Error())
}
