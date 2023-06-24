package util

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGorequestHTTPClient(t *testing.T) {

	t.Run("test go request http call", func(t *testing.T) {
		res, body, err := GorequestHTTPClient(http.MethodGet, "https://google.com", nil, nil, false, 0)
		assert.Nil(t, err, "test ok gorequest http")
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.NotNil(t, body, "test ok resp body")
	})

	t.Run("test go request http call with headers", func(t *testing.T) {
		headers := map[string]string{
			"Accept": "text/html",
		}
		res, body, err := GorequestHTTPClient(http.MethodGet, "https://google.com", headers, nil, false, 0)
		assert.Nil(t, err, "test ok gorequest http")
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.NotNil(t, body, "test ok resp body")
	})

	t.Run("test go request http call with multipart", func(t *testing.T) {
		res, body, err := GorequestHTTPClient(http.MethodGet, "https://google.com", nil, nil, true, 0)
		assert.Nil(t, err, "test ok gorequest http")
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.NotNil(t, body, "test ok resp body")
	})

	t.Run("test go request http call with payload body", func(t *testing.T) {
		payload := map[string]string{
			"name": "qoala",
		}
		res, _, err := GorequestHTTPClient(http.MethodPost, "https://google.com", nil, payload, false, 0)
		assert.Nil(t, err, "test ok gorequest http")
		assert.Equal(t, res.StatusCode, http.StatusMethodNotAllowed)
	})

	t.Run("test go request http call with payload body", func(t *testing.T) {
		_, _, err := GorequestHTTPClient(http.MethodGet, "", nil, nil, false, 0)
		assert.NotNil(t, err, "test nok gorequest http error")
	})
}

func TestGorequestSuperAgent(t *testing.T) {
	t.Run("test go request http call", func(t *testing.T) {
		request := GorequestSuperAgent(http.MethodGet, "https://example.com", nil, nil, false)

		resp, body, requestErrors := request.End()

		if requestErrors != nil {
			t.Errorf("failed construct super agent, err: %s", requestErrors[0])
		}
		defer resp.Body.Close()

		assert.Equal(t, resp.StatusCode, http.StatusOK)
		assert.NotNil(t, body, "test ok resp body")
	})
}
