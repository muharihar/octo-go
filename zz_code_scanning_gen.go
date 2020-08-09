// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	internal "github.com/willabides/octo-go/internal"
	options "github.com/willabides/octo-go/options"
	"net/http"
	"net/url"
)

/*
CodeScanningGetAlert performs requests for "code-scanning/get-alert"

Get a code scanning alert.

  GET /repos/{owner}/{repo}/code-scanning/alerts/{alert_id}

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
func CodeScanningGetAlert(ctx context.Context, req *CodeScanningGetAlertReq, opt ...options.Option) (*CodeScanningGetAlertResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(CodeScanningGetAlertReq)
	}
	resp := &CodeScanningGetAlertResponse{request: req}
	r, err := internal.DoRequest(ctx, req.requestBuilder(), opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = components.CodeScanningAlert{}
	err = internal.DecodeResponseBody(r, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodeScanningGetAlert performs requests for "code-scanning/get-alert"

Get a code scanning alert.

  GET /repos/{owner}/{repo}/code-scanning/alerts/{alert_id}

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
func (c Client) CodeScanningGetAlert(ctx context.Context, req *CodeScanningGetAlertReq, opt ...options.Option) (*CodeScanningGetAlertResponse, error) {
	return CodeScanningGetAlert(ctx, req, append(c, opt...)...)
}

/*
CodeScanningGetAlertReq is request data for Client.CodeScanningGetAlert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
type CodeScanningGetAlertReq struct {
	_url  string
	Owner string
	Repo  string

	// alert_id parameter
	AlertId int64
}

// HTTPRequest builds an *http.Request
func (r *CodeScanningGetAlertReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *CodeScanningGetAlertReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": String("application/json")},
		Method:           "GET",
		OperationID:      "code-scanning/get-alert",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/repos/%v/%v/code-scanning/alerts/%v", r.Owner, r.Repo, r.AlertId),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodeScanningGetAlertReq) Rel(link RelName, resp *CodeScanningGetAlertResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
CodeScanningGetAlertResponse is a response for CodeScanningGetAlert

https://developer.github.com/v3/code-scanning/#get-a-code-scanning-alert
*/
type CodeScanningGetAlertResponse struct {
	internal.Response
	request *CodeScanningGetAlertReq
	Data    components.CodeScanningAlert
}

/*
CodeScanningListAlertsForRepo performs requests for "code-scanning/list-alerts-for-repo"

List code scanning alerts for a repository.

  GET /repos/{owner}/{repo}/code-scanning/alerts

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
func CodeScanningListAlertsForRepo(ctx context.Context, req *CodeScanningListAlertsForRepoReq, opt ...options.Option) (*CodeScanningListAlertsForRepoResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(CodeScanningListAlertsForRepoReq)
	}
	resp := &CodeScanningListAlertsForRepoResponse{request: req}
	r, err := internal.DoRequest(ctx, req.requestBuilder(), opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = []components.CodeScanningAlert{}
	err = internal.DecodeResponseBody(r, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
CodeScanningListAlertsForRepo performs requests for "code-scanning/list-alerts-for-repo"

List code scanning alerts for a repository.

  GET /repos/{owner}/{repo}/code-scanning/alerts

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
func (c Client) CodeScanningListAlertsForRepo(ctx context.Context, req *CodeScanningListAlertsForRepoReq, opt ...options.Option) (*CodeScanningListAlertsForRepoResponse, error) {
	return CodeScanningListAlertsForRepo(ctx, req, append(c, opt...)...)
}

/*
CodeScanningListAlertsForRepoReq is request data for Client.CodeScanningListAlertsForRepo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
type CodeScanningListAlertsForRepoReq struct {
	_url  string
	Owner string
	Repo  string

	// Set to `closed` to list only closed code scanning alerts.
	State *string

	/*
	Returns a list of code scanning alerts for a specific brach reference. The `ref`
	must be formatted as `heads/<branch name>`.
	*/
	Ref *string
}

// HTTPRequest builds an *http.Request
func (r *CodeScanningListAlertsForRepoReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *CodeScanningListAlertsForRepoReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}
	if r.State != nil {
		query.Set("state", *r.State)
	}
	if r.Ref != nil {
		query.Set("ref", *r.Ref)
	}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": String("application/json")},
		Method:           "GET",
		OperationID:      "code-scanning/list-alerts-for-repo",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/repos/%v/%v/code-scanning/alerts", r.Owner, r.Repo),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *CodeScanningListAlertsForRepoReq) Rel(link RelName, resp *CodeScanningListAlertsForRepoResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
CodeScanningListAlertsForRepoResponse is a response for CodeScanningListAlertsForRepo

https://developer.github.com/v3/code-scanning/#list-code-scanning-alerts-for-a-repository
*/
type CodeScanningListAlertsForRepoResponse struct {
	internal.Response
	request *CodeScanningListAlertsForRepoReq
	Data    []components.CodeScanningAlert
}
