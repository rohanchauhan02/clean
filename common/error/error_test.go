package error

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("test ok error as string", func(t *testing.T) {
		err := errors.New("mocked error")
		svcError := ClientError{
			Message:      "error string",
			WrappedError: err,
		}
		errStr := svcError.Error()
		assert.Equal(t, errStr, "mocked error")
	})
}

func TestNew(t *testing.T) {
	t.Run("test ok error new", func(t *testing.T) {
		err := New("error string")
		assert.NotNil(t, err, "success generate error")
		assert.Equal(t, err.Error(), "error string")
	})
}

func TestWrap(t *testing.T) {
	t.Run("test ok error wrapper", func(t *testing.T) {
		errSys := New("error string")
		err := Wrap(http.StatusOK, "invalid_email_address", errSys, "invalid email", "failed")
		assert.NotNil(t, err, "success wrap error")
	})
}
func TestErrorsResponseBody(t *testing.T) {
	t.Run("test valid responseBody", func(t *testing.T) {
		mockedErr := errors.New("mocked error")
		svcErr, initErr := NewClientError(mockedErr, http.StatusBadRequest, "SVC-ERR-001", "Bad Request", "")
		assert.Nil(t, initErr)
		resp, err := svcErr.ResponseBody()
		assert.Nil(t, err)
		assert.Equal(t,
			resp, "{\"code\":400,\"error_code\":\"SVC-ERR-001\",\"message\":\"Bad Request\",\"status\":\"\"}")
	})

	t.Run("test valid responseBody", func(t *testing.T) {
		mockedErr := errors.New("mocked error")
		svcErr, initErr := NewServerError(mockedErr, http.StatusInternalServerError, "SVC-ERR-001", "Internal Server Error", "")
		assert.Nil(t, initErr)
		resp, err := svcErr.ResponseBody()
		assert.Nil(t, err)
		assert.Equal(t,
			resp, "{\"code\":500,\"error_code\":\"SVC-ERR-001\",\"message\":\"Internal Server Error\",\"status\":\"\"}")
	})

	t.Run("test http status code 500 returns error on NewClientError", func(t *testing.T) {
		mockedErr := errors.New("mocked error")
		_, initErr := NewClientError(mockedErr, http.StatusInternalServerError, "SVC-ERR-001", "Bad Request", "")
		assert.Error(t, initErr)
	})

	t.Run("test http status code 400 returns error on NewServerErro", func(t *testing.T) {
		mockedErr := errors.New("mocked error")
		_, initErr := NewServerError(mockedErr, http.StatusBadRequest, "SVC-ERR-001", "Bad Request", "")
		assert.Error(t, initErr)
	})
}

func TestTypeAssertionsOnWrappedError(t *testing.T) {
	t.Run("test valid type on wrapped ClientError", func(t *testing.T) {
		mockedErr := errors.New("mocked error")
		svcErr, initErr := NewClientError(mockedErr, http.StatusBadRequest, "", "", "")
		assert.Nil(t, initErr)
		assert.NotNil(t, svcErr.Unwrap())
		assert.False(t, IsServerError(svcErr))
	})

	t.Run("test valid type on wrapped ServerError", func(t *testing.T) {
		mockedErr := errors.New("mocked error")
		svcErr, initErr := NewServerError(mockedErr, http.StatusBadGateway, "", "", "")
		assert.Nil(t, initErr)
		assert.NotNil(t, svcErr.Unwrap())
		assert.True(t, IsServerError(svcErr))
	})
}
