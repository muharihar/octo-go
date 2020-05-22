// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

/*
RateLimitGetReq builds requests for "rate-limit/get"

Get your current rate limit status.

  GET /rate_limit

https://developer.github.com/v3/rate_limit/#get-your-current-rate-limit-status
*/
type RateLimitGetReq struct{}

func (r RateLimitGetReq) urlPath() string {
	return fmt.Sprintf("/rate_limit")
}

func (r RateLimitGetReq) method() string {
	return "GET"
}

func (r RateLimitGetReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r RateLimitGetReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r RateLimitGetReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
RateLimitGetResponseBody200 is a response body for rate-limit/get

API documentation: https://developer.github.com/v3/rate_limit/#get-your-current-rate-limit-status
*/
type RateLimitGetResponseBody200 struct {
	Rate struct {
		Limit     int64 `json:"limit,omitempty"`
		Remaining int64 `json:"remaining,omitempty"`
		Reset     int64 `json:"reset,omitempty"`
	} `json:"rate,omitempty"`
	Resources struct {
		Core struct {
			Limit     int64 `json:"limit,omitempty"`
			Remaining int64 `json:"remaining,omitempty"`
			Reset     int64 `json:"reset,omitempty"`
		} `json:"core,omitempty"`
		Graphql struct {
			Limit     int64 `json:"limit,omitempty"`
			Remaining int64 `json:"remaining,omitempty"`
			Reset     int64 `json:"reset,omitempty"`
		} `json:"graphql,omitempty"`
		IntegrationManifest struct {
			Limit     int64 `json:"limit,omitempty"`
			Remaining int64 `json:"remaining,omitempty"`
			Reset     int64 `json:"reset,omitempty"`
		} `json:"integration_manifest,omitempty"`
		Search struct {
			Limit     int64 `json:"limit,omitempty"`
			Remaining int64 `json:"remaining,omitempty"`
			Reset     int64 `json:"reset,omitempty"`
		} `json:"search,omitempty"`
	} `json:"resources,omitempty"`
}