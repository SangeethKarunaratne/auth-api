package errors

import "fmt"

type AdapterError struct {
	errType string
	code    int
	msg     string
	details string
}

func NewAdapterError(message string, code int, details string) error {

	return &AdapterError{
		errType: "AdapterError",
		code:    code,
		msg:     message,
		details: details,
	}
}

func (e *AdapterError) Error() string {
	return fmt.Sprintf("%s|%d|%s|%s", e.errType, e.code, e.msg, e.details)
}
