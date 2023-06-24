package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareRequestID(t *testing.T) {
	t.Run("test middleware requestID", func(t *testing.T) {
		e := echo.New()
		e.Use(MiddlewareRequestID())

		req := httptest.NewRequest(echo.GET, "/", nil)
		res := httptest.NewRecorder()
		e.ServeHTTP(res, req)

		assert.NotEmpty(t, req.Header.Get(echo.HeaderXRequestID))
		assert.NotEmpty(t, res.Header().Get(echo.HeaderXRequestID))
	})
}
