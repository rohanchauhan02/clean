package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo"
	gormHistory "github.com/rohanchauhan02/clean/common/gorm-history"
	"github.com/rohanchauhan02/clean/common/util"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func PartnerKeyAuthentication(config Config) func(http.Handler) http.Handler {
	if config.PartnerKeyEndpoint == "" {
		panic("qoala common jwt authentication middleware needs auth endpoint")
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("x-api-key")
			if token == "" {
				unauthorizedResponse(errors.New(ApiKeyEmpty), "QC-CLT-KEYA-V1-001", ErrMessageUnauthorizedToken, w)
				return
			}

			userV1RequestHeaders := map[string]string{
				"Content-Type": "application/json",
				"x-request-id": r.Header.Get("X-Request-ID"),
				"x-api-key":    token,
			}
			res, body, resError := util.GorequestHTTPClient(http.MethodGet, config.PartnerKeyEndpoint,
				userV1RequestHeaders, nil, false, 100*time.Millisecond)

			if resError != nil {
				badRequestResponse(resError, "QC-CLT-KEYA-V1-002", resError.Error(), w)
				return
			}

			if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
				var errResponse PartnerKeyErrorResponse
				err := json.Unmarshal([]byte(body), &errResponse)
				if err != nil {
					internalServerResponse(err, "QC-SVR-KEYA-V1-001", UnmarshallUserServiceError, w)
					return
				}
				unauthorizedResponse(errors.New(errResponse.Message), "QC-CLT-KEYA-V1-003", ErrMessageUnauthorizedToken, w)
				return
			}

			partnerCode := res.Header.Get("x-partner-code")
			userData := PartnerKeyResponse{
				PartnerCode: partnerCode,
			}

			historySourceData := gormHistory.Source{
				RequestID: r.Header.Get("X-Request-ID"),
			}

			ctx := context.WithValue(r.Context(), ContextUserKey, userData)
			ctx = context.WithValue(ctx, ContextHistorySourceKey, historySourceData)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func PartnerKeyAuthenticationV2(config Config) echo.MiddlewareFunc {
	if config.PartnerKeyEndpoint == "" {
		panic("qoala common jwt authentication middleware needs auth endpoint")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ac := c.(*util.CustomApplicationContext)
			start := time.Now()
			ctx := c.Request().Context()
			if ac.CustomContext != nil {
				ctx = ac.CustomContext
			}
			span, spanCtx := tracer.StartSpanFromContext(ctx, "PartnerKeyAuthenticationV2",
				tracer.SpanType("function"),
			)
			ac.CustomContext = spanCtx
			defer span.Finish()
			token := c.Request().Header.Get("x-api-key")
			if token == "" {
				c.Response().Status = http.StatusUnauthorized
				return errors.New(ApiKeyEmpty)
			}

			userV1RequestHeaders := map[string]string{
				"Content-Type": "application/json",
				"x-request-id": c.Request().Header.Get("X-Request-ID"),
				"x-api-key":    token,
			}
			res, body, resError := util.GorequestHTTPClient(http.MethodGet, config.PartnerKeyEndpoint,
				userV1RequestHeaders, nil, false, config.Timeout)

			if resError != nil {
				c.Response().Status = http.StatusBadRequest
				return resError
			}

			if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
				var errResponse PartnerKeyErrorResponse
				err := json.Unmarshal([]byte(body), &errResponse)
				if err != nil {
					c.Response().Status = http.StatusInternalServerError
					return err
				} else {
					c.Response().Status = http.StatusUnauthorized
					return errors.New(errResponse.Message)
				}
			}

			partnerCode := res.Header.Get("x-partner-code")
			userData := PartnerKeyResponse{
				PartnerCode: partnerCode,
			}

			historySourceData := gormHistory.Source{
				RequestID: c.Request().Header.Get("X-Request-ID"),
			}
			ac.Context.Set(ContextUserKey, userData)
			ac.Context.Set(ContextHistorySourceKey, historySourceData)
			end := time.Now()

			ac.DatadogClient.SendDurationMetric("PartnerKeyAuthenticationV2", start, end, config.Caller)

			if err = next(c); err != nil {
				return err
			}
			return
		}
	}
}
