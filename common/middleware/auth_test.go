package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	jwtEndpoint        = "https://api-staging.qoala.app/api/sessions/check"
	partnerKeyEndpoint = "https://api-staging.qoala.app/v2/sessions/partner/check-api-key"
)

var mockConfig = Config{JWTKeyEndpoint: jwtEndpoint, PartnerKeyEndpoint: partnerKeyEndpoint}

func TestJWTAuthWithInvalidTokens(t *testing.T) {
	t.Run("test empty header", func(t *testing.T) {
		middlewareFunc := []func(handler http.Handler) http.Handler{JWTAuthenticationV1(mockConfig),
			JWTAuthenticationV2(mockConfig)}
		for _, mf := range middlewareFunc {
			e := echo.New()
			e.GET("/", func(context echo.Context) error {
				return nil
			}, echo.WrapMiddleware(mf))

			req := httptest.NewRequest(echo.GET, "/", nil)
			req.Header.Set("Authorization", "Bearer ")
			res := httptest.NewRecorder()
			e.ServeHTTP(res, req)

			assert.Equal(t, res.Code, http.StatusUnauthorized)
		}
	})

	t.Run("test malformed jwt token", func(t *testing.T) {
		middlewareFunc := []func(handler http.Handler) http.Handler{JWTAuthenticationV1(mockConfig),
			JWTAuthenticationV2(mockConfig)}
		for _, mf := range middlewareFunc {
			e := echo.New()
			e.GET("/", func(context echo.Context) error {
				return nil
			}, echo.WrapMiddleware(mf))

			req := httptest.NewRequest(echo.GET, "/", nil)
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.uLk1FPtMrqh70SHpn1O5Mbx50cq3nKVuR78GSzfrIBE")
			res := httptest.NewRecorder()
			e.ServeHTTP(res, req)

			assert.Equal(t, res.Code, http.StatusUnauthorized)
		}

	})
}

func TestJWTAuthWithValidUserServiceResponse(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzQ2NjIzODMsImV4cCI6MTYzNDc0ODc4MywiYXVkIjoiYWRtIiwiaXNzIjoicW9hbGEqUU9BTEEiLCJzdWIiOiJ1RW9GU3dKVjE0Z0tKYkRLd0pVa3JkIn0.uLk1FPtMrqh70SHpn1O5Mbx50cq3nKVuR78GSzfrIBE"

	// go request test stubs that are needed for httpMock register responder to comply
	// with non standard http client package (gorequest etc)
	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var jsonResponse CheckUserV1Response
	jsonString := `{
			"data": {
				"user": {
					"id": 1234,
					"uuid": "e81c209b-3120-4b1a-s89z-f184ac133022",
					"firstName": "Qoala",
					"lastName": "Qoala",
					"fullName": "Qoala Qoala"
				}
			}
		}`
	err := json.Unmarshal([]byte(jsonString), &jsonResponse)
	if err != nil {
		t.Fatalf("error in unmarshalling json string for mock checkUserResponse")
	}

	httpmock.RegisterResponder("POST", jwtEndpoint,
		func(req *http.Request) (*http.Response, error) {
			checkUserRequest := CheckUserV1Request{Token: token}
			if err := json.NewDecoder(req.Body).Decode(&checkUserRequest); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			resp, err := httpmock.NewJsonResponse(200, jsonResponse)
			if err != nil {
				return httpmock.NewStringResponse(500, "Internal Server Error"), nil
			}
			return resp, nil
		},
	)
	middlewareFunc := []func(handler http.Handler) http.Handler{JWTAuthenticationV1(mockConfig),
		JWTAuthenticationV2(mockConfig)}
	for _, mf := range middlewareFunc {
		var ctx echo.Context
		e := echo.New()
		e.GET("/", func(context echo.Context) error {
			ctx = context
			return nil
		}, echo.WrapMiddleware(mf))

		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		res := httptest.NewRecorder()
		e.ServeHTTP(res, req)

		assert.Equal(t, res.Code, http.StatusOK)
		assert.NotNil(t, ctx.Request().Context().Value(ContextUserKey))
	}
}

func TestJWTAuthWithInvalidUserServiceResponse(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzQ2NjIzODMsImV4cCI6MTYzNDc0ODc4MywiYXVkIjoiYWRtIiwiaXNzIjoicW9hbGEqUU9BTEEiLCJzdWIiOiJ1RW9GU3dKVjE0Z0tKYkRLd0pVa3JkIn0.uLk1FPtMrqh70SHpn1O5Mbx50cq3nKVuR78GSzfrIBE"

	// currently CheckUserV1ErrorResponse == CheckUserV2Response in terms
	// of structure so we can just use one declaration for the test
	var jsonResponse CheckUserV1ErrorResponse
	jsonString := `{
			"message": "Unauthorized from user-service"
		}`
	err := json.Unmarshal([]byte(jsonString), &jsonResponse)
	if err != nil {
		t.Fatalf("error in unmarshalling json string for mock checkUserV1Response")
	}

	// go request test stubs that are needed for httpMock register responder to comply
	// with non standard http client package (gorequest etc)
	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", jwtEndpoint,
		func(req *http.Request) (*http.Response, error) {
			checkUserRequest := CheckUserV1Request{Token: token}
			if err := json.NewDecoder(req.Body).Decode(&checkUserRequest); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			resp, err := httpmock.NewJsonResponse(401, jsonResponse)
			if err != nil {
				return httpmock.NewStringResponse(500, "Internal Server Error"), nil
			}
			return resp, nil
		},
	)
	middlewareFunc := []func(handler http.Handler) http.Handler{JWTAuthenticationV1(mockConfig),
		JWTAuthenticationV2(mockConfig)}
	for _, mf := range middlewareFunc {
		e := echo.New()
		e.GET("/", func(context echo.Context) error {
			return nil
		}, echo.WrapMiddleware(mf))

		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res := httptest.NewRecorder()
		e.ServeHTTP(res, req)

		assert.Equal(t, res.Code, http.StatusUnauthorized)
	}
}

func TestPartnerKeyAuthWithEmptyKey(t *testing.T) {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return nil
	}, echo.WrapMiddleware(PartnerKeyAuthentication(mockConfig)))

	req := httptest.NewRequest(echo.GET, "/", nil)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusUnauthorized)
}

func TestPartnerKeyAuthWithValidKey(t *testing.T) {

	mockApiKey := "sa97230askdfoaij832098rkjepie0qw0iebfb"
	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()


	httpmock.RegisterResponder("GET", partnerKeyEndpoint,
		func(req *http.Request) (*http.Response, error) {
			req.Header.Set("x-api-key", mockApiKey)
			resp, err := httpmock.NewJsonResponse(200, "")
			if err != nil {
				return httpmock.NewStringResponse(500, "Internal Server Error"), nil
			}
			return resp, nil
		},
	)

	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return nil
	}, echo.WrapMiddleware(PartnerKeyAuthentication(mockConfig)))
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set("x-api-key", mockApiKey)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusOK)
}

func TestPartnerKeyAuthWithNotFoundKey(t *testing.T) {

	mockApiKey := "sa97230askdfoaij832098rkjepie0qw0iebfb"
	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()


	var jsonResponse PartnerKeyErrorResponse
	jsonString := `{
			"message": "Unauthorized from user-service"
		}`
	err := json.Unmarshal([]byte(jsonString), &jsonResponse)
	if err != nil {
		t.Fatalf("error in unmarshalling json string for mock checkUserV1Response")
	}

	httpmock.RegisterResponder("GET", partnerKeyEndpoint,
		func(req *http.Request) (*http.Response, error) {
			req.Header.Set("x-api-key", mockApiKey)
			resp, err := httpmock.NewJsonResponse(404, jsonResponse)
			if err != nil {
				return httpmock.NewStringResponse(500, "Internal Server Error"), nil
			}
			return resp, nil
		},
	)

	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return nil
	}, echo.WrapMiddleware(PartnerKeyAuthentication(mockConfig)))
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set("x-api-key", mockApiKey)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusUnauthorized)
}