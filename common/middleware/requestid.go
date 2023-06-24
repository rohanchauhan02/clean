package middleware

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	bodyDumpResponseWriter struct {
		io.Writer
		http.ResponseWriter
	}
)

func MiddlewareRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get(echo.HeaderXRequestID)
			if requestID == "" {
				requestID = generateRequestID()
			}
			c.Request().Header.Set(echo.HeaderXRequestID, requestID)
			c.Response().Header().Set(echo.HeaderXRequestID, requestID)
			return next(c)
		}
	}
}

func generateRequestID() string {
	return uuid.New().String()
}

func MiddlewareDumpRequestResponse(skipper middleware.Skipper) echo.MiddlewareFunc {
	return BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: skipper, //we can skip for spesific endpoint or other rules in the future
	})
}

// BodyDumpWithConfig returns a BodyDump middleware with config.
// See: `BodyDump()`.
func BodyDumpWithConfig(config middleware.BodyDumpConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultBodyDumpConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}
			idrequest := time.Now().UnixNano()
			// Request
			reqBody := []byte{}
			if c.Request().Body != nil { // Read
				reqBody, _ = ioutil.ReadAll(c.Request().Body)
				//print out the request
				cleanBody := strings.ReplaceAll(string(reqBody), "\n", "")
				cleanBody = strings.ReplaceAll(cleanBody, "\t", "")
				cleanBody = strings.ReplaceAll(cleanBody, " ", "")
				logger.Infof("url: %s identifier:%d -- request: %s ", c.Request().RequestURI, idrequest, cleanBody)
			}
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset

			// Response
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(c.Response().Writer, resBody)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
			c.Response().Writer = writer

			if err = next(c); err != nil {
				c.Error(err)
			}

			//print out the response
			logger.Infof("url: %s identifier:%d -- response: %s ", c.Request().RequestURI, idrequest, resBody.String())
			return
		}
	}
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func (w *bodyDumpResponseWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}
