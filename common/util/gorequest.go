package util

import (
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

// GorequestHTTPClient construct http call with go request
// given timeout value 0 to disable timeout
func GorequestHTTPClient(method string, url string, headers map[string]string, payload interface{}, isMultipart bool, timeout time.Duration) (*http.Response, string, error) {

	request := gorequest.New().CustomMethod(method, url)

	if len(headers) > 0 {
		for k, v := range headers {
			request.Set(k, v)
		}
	}

	if timeout > 0 {
		request.Timeout(timeout)
	}

	if isMultipart {
		request.Type("multipart")
	}

	if method == http.MethodPost || method == http.MethodPut {
		request.Send(payload)
	}

	res, body, requestErrors := request.End()

	if requestErrors != nil {
		return nil, "", requestErrors[0]
	}
	defer res.Body.Close()

	return res, body, nil
}

// GorequestSuperAgent construct http call with go request return superAgent
func GorequestSuperAgent(method string, url string, headers map[string]string, payload interface{}, isMultipart bool) *gorequest.SuperAgent {

	request := gorequest.New().CustomMethod(method, url)

	if len(headers) > 0 {
		for k, v := range headers {
			request.Set(k, v)
		}
	}

	if isMultipart {
		request.Type("multipart")
	}

	if method == http.MethodPost || method == http.MethodPut {
		request.Send(payload)
	}

	return request
}
