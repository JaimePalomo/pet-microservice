package rest_errors

import "fmt"

type restError struct {
	ErrStatus  int    `json:"status"`
	ErrMessage string `json:"message"`
	ErrError   string `json:"error"`
}

type RestError interface {
	Status() int
	Message() string
	Error() string
}

func New(status int, message string) RestError {
	return &restError{
		ErrStatus:  status,
		ErrMessage: message,
		ErrError:   fmt.Sprintf("status: %d - message: %s", status, message),
	}
}

func (r restError) Status() int {
	return r.ErrStatus
}

func (r restError) Message() string {
	return r.ErrMessage
}

func (r restError) Error() string {
	return r.ErrError
}
