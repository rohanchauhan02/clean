package models

import (
	"fmt"
)

// QoalaError is a general qoala error response struct
type QoalaError struct {
	Code      int
	ErrorCode string
	Message   string
	Status    string
}

func (e QoalaError) Error() string {
	return fmt.Sprintf("Error Occurred. CODE: %d; ERROR_CODE: %s; MESSAGE: %s", e.Code, e.ErrorCode, e.Message)
}

func NewQoalaError(code int, errorCode string, message string, status string) QoalaError {
	return QoalaError{
		Code:      code,
		ErrorCode: errorCode,
		Message:   message,
		Status:    status,
	}
}
