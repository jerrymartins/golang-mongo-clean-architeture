package error

import (
	"errors"
	"fmt"
)

type RequestError struct {
	StatusCodeNumber int
	Err              error
}

func (r *RequestError) Error() string {
	return fmt.Sprint(r.Err)
}

func (r *RequestError) StatusCode() int {
	return r.StatusCodeNumber
}

func ConfigureRequestError(message string, statusCode int) RequestError {
	return RequestError{
		StatusCodeNumber: statusCode,
		Err:              errors.New(message),
	}
}
