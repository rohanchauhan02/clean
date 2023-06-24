// Qoala Common Error Package Library
// This is used for generating error messages to clients/end-users
// and also to run type assertions on the errors in order to run
// some cases for e.g: alerts on server errors

// Usage:
// On caller side, a NewClientError or NewServerError is initialized and the
// caller has control over the error code, error message and
// status code of the error. ResponseBody() function will
// returns JSON-output error message

// A type is defined for ClientError and ServerError intentionally.
// This will allow for robust alerting behavior by checking the error types
// for example by using IsServerError in a middleware to decide whether
// such errors need to be alerted to Slack/Sentry

package error

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	Errors "github.com/pkg/errors"
)

type ServiceError interface {
	Error() string
	ResponseBody() (string, error)
}

// ServiceError is a general Service error that implements
// error (specifically ServiceError) which are errors
// returned to the users.
type ClientError struct {
	Code         int    `json:"code"`
	ErrorCode    string `json:"error_code"`
	Message      string `json:"message"`
	Status       string `json:"status"`
	WrappedError error  `json:"-"`
}

// ServerError is a general Service error that implements
// error (specifically ServiceError) which are errors
// returned to the users.
type ServerError struct {
	Code         int    `json:"code"`
	ErrorCode    string `json:"error_code"`
	Message      string `json:"message"`
	Status       string `json:"status"`
	WrappedError error  `json:"-"`
}

// Error returns error type as a string that returns the
// error message supplied along with a stacktrace
// of when Wrap was called
func (q *ClientError) Error() string {
	return fmt.Sprintf("%v", q.WrappedError)
}

// Error returns error type as a string that returns the
// error message supplied along with a stacktrace
// of when Wrap was called
func (q *ServerError) Error() string {
	return fmt.Sprintf("%v", q.WrappedError)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil. Usually used for type assertion
// on the error
func (q *ClientError) Unwrap() error {
	return Errors.Unwrap(q.WrappedError)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil. Usually used for type assertion
// on the error
func (q *ServerError) Unwrap() error {
	return Errors.Unwrap(q.WrappedError)
}

// ResponseBody returns a standardized JSON format for
// error response that can be supplied back to the client
func (q *ClientError) ResponseBody() (string, error) {
	resp, err := json.Marshal(q)
	if err != nil {
		return "", fmt.Errorf("error marshalling error struct to json")
	}
	return string(resp), nil
}

// ResponseBody returns a standardized JSON format for
// error response that can be supplied back to the client
func (q *ServerError) ResponseBody() (string, error) {
	resp, err := json.Marshal(q)
	if err != nil {
		return "", fmt.Errorf("error marshalling error struct to json")
	}
	return string(resp), nil
}

// NewClientError returns a wrapped error of type ClientError
// type assertion can be done on caller side
func NewClientError(err error, code int, errorCode string, errorMsg string, status string) (*ClientError, error) {
	if code > 499 {
		return nil, fmt.Errorf("Cannot construct NewClientError for http status code > 499. Use NewServerError instead")
	}
	if status == "" {
		status = "failed"
	}
	wrappedErr := Errors.Wrap(err, err.Error())
	return &ClientError{
		Code:         code,
		ErrorCode:    errorCode,
		Message:      errorMsg,
		Status:       status,
		WrappedError: wrappedErr,
	}, nil
}

// NewServer returns a wrapped error of type ServerError
// type assertion can be done on caller side
func NewServerError(err error, code int, errorCode string, errorMsg string, status string) (*ServerError, error) {
	if code < 500 {
		return nil, fmt.Errorf("Cannot construct NewServerError for http status code < 500. Use NewClientError instead")
	}
	if status == "" {
		status = "failed"
	}
	wrappedErr := Errors.Wrap(err, err.Error())
	return &ServerError{
		Code:         code,
		ErrorCode:    errorCode,
		Message:      errorMsg,
		Status:       status,
		WrappedError: wrappedErr,
	}, nil
}

//IsServerError will return whether the error type passed
//is of type ServerError. This is basically type assertion
//as a helper function. Typically this function is used in usecase
//when you want to decide to alerts types of errors.
func IsServerError(err error) bool {
	var ew *ServerError
	return Errors.As(err, &ew)
}

// New returnns new error message in standard pkg errors new
func New(msg string) error {
	return Errors.New(msg)
}

// Wrap returns a new error that adds context to the original error
func Wrap(code int, errorCode string, err error, msg string, status string) error {
	return Errors.Wrap(&ClientError{
		Code:      code,
		ErrorCode: errorCode,
		Message:   msg,
		Status:    status,
	}, err.Error())
}

var ErrorHandler = func(err error, c echo.Context) {
	statusCode := http.StatusInternalServerError
	switch err.(type) {
	case *ClientError:
		cErr := err.(*ClientError)
		statusCode = cErr.Code
	case *ServerError:
		sErr := err.(*ServerError)
		statusCode = sErr.Code
	default:
		err, _ = NewServerError(errors.New("internal server error"), http.StatusInternalServerError, "QES-ERR-001", "internal server error", "failed")
		statusCode = http.StatusInternalServerError
	}
	if !c.Response().Committed {
		_ = c.JSON(statusCode, err)
	}
}
