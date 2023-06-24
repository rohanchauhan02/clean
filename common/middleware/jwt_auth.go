package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	gormHistory "github.com/rohanchauhan02/clean/common/gorm-history"
	log "github.com/rohanchauhan02/common/logs"
	"github.com/rohanchauhan02/clean/common/util"
)

var (
	logger = log.NewCommonLog()
)

func JWTAuthenticationV1(config Config) func(http.Handler) http.Handler {
	if config.JWTKeyEndpoint == "" {
		panic("qoala common jwt authentication middleware needs auth endpoint")
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				unauthorizedResponse(errors.New(AuthorizationHeaderEmpty), "QC-CLT-UNATH-V1-001", ErrMessageUnauthorizedToken, w)
				return
			}
			jwtTokenArr := strings.Split(token, " ")
			if len(jwtTokenArr) != 2 {
				unauthorizedResponse(errors.New(MalformedBearerTokenFormatting), "QC-CLT-UNATH-V1-002", ErrMessageUnauthorizedToken, w)
				return
			}

			jwtToken := jwtTokenArr[1]
			_, err := jwt.Parse(jwtToken, nil)
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&(jwt.ValidationErrorMalformed) != 0 {
					unauthorizedResponse(errors.New(MalformedJWTToken), "QC-CLT-UNATH-V1-003", ErrMessageUnauthorizedToken, w)
					return
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					unauthorizedResponse(errors.New(MalformedExpiredToken), "QC-CLT-UNATH-V1-004", ErrMessageUnauthorizedToken, w)
					return
				}
			}

			userV1RequestDTO := CheckUserV1Request{
				Token: jwtToken,
			}
			userV1RequestHeaders := map[string]string{
				"Content-Type": "application/json",
				"x-request-id": r.Header.Get("X-Request-ID"),
			}
			res, body, resError := util.GorequestHTTPClient(http.MethodPost, config.JWTKeyEndpoint,
				userV1RequestHeaders, userV1RequestDTO, false, 0)

			if resError != nil {
				badRequestResponse(resError, "QC-CLT-UNATH-V1-005", resError.Error(), w)
				return
			}

			if res.StatusCode != http.StatusOK {
				var errResponse CheckUserV1ErrorResponse
				err = json.Unmarshal([]byte(body), &errResponse)
				if err != nil {
					internalServerResponse(err, "QC-SVR-UNATH-V1-001", UnmarshallUserServiceError, w)
					return
				}
				unauthorizedResponse(errors.New(errResponse.Message), "QC-CLT-UNATH-V1-006", ErrMessageUnauthorizedToken, w)
				return
			}

			userData := CheckUserV1Response{}
			err = json.Unmarshal([]byte(body), &userData)
			if err != nil {
				internalServerResponse(err, "QC-SVR-UNATH-V1-002", UnmarshallUserServiceError, w)
			}

			historyUserData := gormHistory.User{
				ID:       userData.Data.User.UUID,
				Type:     "USER",
				FullName: userData.Data.User.FullName,
				Email:    userData.Data.User.Email,
			}

			historySourceData := gormHistory.Source{
				RequestID: r.Header.Get("X-Request-ID"),
			}

			ctx := context.WithValue(r.Context(), ContextUserKey, userData)
			ctx = context.WithValue(ctx, ContextHistoryUserKey, historyUserData)
			ctx = context.WithValue(ctx, ContextHistorySourceKey, historySourceData)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func JWTAuthenticationV2(config Config) func(http.Handler) http.Handler {
	if config.JWTKeyEndpoint == "" {
		panic("qoala common jwt authentication middleware needs auth endpoint")
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				unauthorizedResponse(errors.New(AuthorizationHeaderEmpty), "QC-CLT-UNATH-V2-001", ErrMessageUnauthorizedToken, w)
				return
			}
			jwtTokenArr := strings.Split(token, " ")
			if len(jwtTokenArr) != 2 {
				unauthorizedResponse(errors.New(MalformedBearerTokenFormatting), "QC-CLT-UNATH-V2-002", ErrMessageUnauthorizedToken, w)
				return
			}

			jwtToken := jwtTokenArr[1]
			_, err := jwt.Parse(jwtToken, nil)
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&(jwt.ValidationErrorMalformed) != 0 {
					unauthorizedResponse(errors.New(MalformedJWTToken), "QC-CLT-UNATH-V2-003", ErrMessageUnauthorizedToken, w)
					return
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					unauthorizedResponse(errors.New(MalformedExpiredToken), "QC-CLT-UNATH-V2-004", ErrMessageUnauthorizedToken, w)
					return
				}
			}

			userV1RequestDTO := CheckUserV2Request{
				Token: jwtToken,
			}
			userV1RequestHeaders := map[string]string{
				"Content-Type": "application/json",
				"x-request-id": r.Header.Get("X-Request-ID"),
			}
			res, body, resError := util.GorequestHTTPClient(http.MethodPost, config.JWTKeyEndpoint,
				userV1RequestHeaders, userV1RequestDTO, false, 0)

			if resError != nil {
				badRequestResponse(resError, "QC-CLT-UNATH-V2-005", resError.Error(), w)
				return
			}

			if res.StatusCode != http.StatusOK {
				var errResponse CheckUserV2ErrorResponse
				err = json.Unmarshal([]byte(body), &errResponse)
				if err != nil {
					internalServerResponse(err, "QC-SVR-UNATH-V2-001", UnmarshallUserServiceError, w)
				}
				unauthorizedResponse(errors.New(errResponse.Message), "QC-CLT-UNATH-V1-006", ErrMessageUnauthorizedToken, w)
				return
			}

			userData := CheckUserV2Response{}
			err = json.Unmarshal([]byte(body), &userData)
			if err != nil {
				internalServerResponse(err, "QC-SVR-UNATH-V2-002", UnmarshallUserServiceError, w)
			}

			historyUserData := gormHistory.User{
				ID:       userData.Data.User.UUID,
				Type:     "USER",
				FullName: userData.Data.User.FullName,
				Email:    userData.Data.User.Email,
			}

			historySourceData := gormHistory.Source{
				RequestID: r.Header.Get("X-Request-ID"),
			}

			ctx := context.WithValue(r.Context(), ContextUserKey, userData)
			ctx = context.WithValue(ctx, ContextHistoryUserKey, historyUserData)
			ctx = context.WithValue(ctx, ContextHistorySourceKey, historySourceData)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
