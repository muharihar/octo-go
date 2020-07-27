// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	"net/http"
	"net/url"
)

/*
BillingGetGithubActionsBillingGhe performs requests for "billing/get-github-actions-billing-ghe"

Get GitHub Actions billing for an enterprise.

  GET /enterprises/{enterprise_id}/settings/billing/actions

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-enterprise
*/
func BillingGetGithubActionsBillingGhe(ctx context.Context, req *BillingGetGithubActionsBillingGheReq, opt ...RequestOption) (*BillingGetGithubActionsBillingGheResponse, error) {
	if req == nil {
		req = new(BillingGetGithubActionsBillingGheReq)
	}
	resp := &BillingGetGithubActionsBillingGheResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.ActionsBillingUsage{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
BillingGetGithubActionsBillingGhe performs requests for "billing/get-github-actions-billing-ghe"

Get GitHub Actions billing for an enterprise.

  GET /enterprises/{enterprise_id}/settings/billing/actions

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-enterprise
*/
func (c Client) BillingGetGithubActionsBillingGhe(ctx context.Context, req *BillingGetGithubActionsBillingGheReq, opt ...RequestOption) (*BillingGetGithubActionsBillingGheResponse, error) {
	return BillingGetGithubActionsBillingGhe(ctx, req, append(c, opt...)...)
}

/*
BillingGetGithubActionsBillingGheReq is request data for Client.BillingGetGithubActionsBillingGhe

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-enterprise
*/
type BillingGetGithubActionsBillingGheReq struct {
	_url string

	// Unique identifier of the GitHub Enterprise Cloud instance.
	EnterpriseId string
}

func (r *BillingGetGithubActionsBillingGheReq) url() string {
	return r._url
}

func (r *BillingGetGithubActionsBillingGheReq) urlPath() string {
	return fmt.Sprintf("/enterprises/%v/settings/billing/actions", r.EnterpriseId)
}

func (r *BillingGetGithubActionsBillingGheReq) method() string {
	return "GET"
}

func (r *BillingGetGithubActionsBillingGheReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *BillingGetGithubActionsBillingGheReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *BillingGetGithubActionsBillingGheReq) body() interface{} {
	return nil
}

func (r *BillingGetGithubActionsBillingGheReq) dataStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubActionsBillingGheReq) validStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubActionsBillingGheReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *BillingGetGithubActionsBillingGheReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *BillingGetGithubActionsBillingGheReq) Rel(link RelName, resp *BillingGetGithubActionsBillingGheResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
BillingGetGithubActionsBillingGheResponse is a response for BillingGetGithubActionsBillingGhe

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-enterprise
*/
type BillingGetGithubActionsBillingGheResponse struct {
	response
	request *BillingGetGithubActionsBillingGheReq
	Data    components.ActionsBillingUsage
}

/*
BillingGetGithubActionsBillingOrg performs requests for "billing/get-github-actions-billing-org"

Get GitHub Actions billing for an organization.

  GET /orgs/{org}/settings/billing/actions

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-organization
*/
func BillingGetGithubActionsBillingOrg(ctx context.Context, req *BillingGetGithubActionsBillingOrgReq, opt ...RequestOption) (*BillingGetGithubActionsBillingOrgResponse, error) {
	if req == nil {
		req = new(BillingGetGithubActionsBillingOrgReq)
	}
	resp := &BillingGetGithubActionsBillingOrgResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.ActionsBillingUsage{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
BillingGetGithubActionsBillingOrg performs requests for "billing/get-github-actions-billing-org"

Get GitHub Actions billing for an organization.

  GET /orgs/{org}/settings/billing/actions

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-organization
*/
func (c Client) BillingGetGithubActionsBillingOrg(ctx context.Context, req *BillingGetGithubActionsBillingOrgReq, opt ...RequestOption) (*BillingGetGithubActionsBillingOrgResponse, error) {
	return BillingGetGithubActionsBillingOrg(ctx, req, append(c, opt...)...)
}

/*
BillingGetGithubActionsBillingOrgReq is request data for Client.BillingGetGithubActionsBillingOrg

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-organization
*/
type BillingGetGithubActionsBillingOrgReq struct {
	_url string
	Org  string
}

func (r *BillingGetGithubActionsBillingOrgReq) url() string {
	return r._url
}

func (r *BillingGetGithubActionsBillingOrgReq) urlPath() string {
	return fmt.Sprintf("/orgs/%v/settings/billing/actions", r.Org)
}

func (r *BillingGetGithubActionsBillingOrgReq) method() string {
	return "GET"
}

func (r *BillingGetGithubActionsBillingOrgReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *BillingGetGithubActionsBillingOrgReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *BillingGetGithubActionsBillingOrgReq) body() interface{} {
	return nil
}

func (r *BillingGetGithubActionsBillingOrgReq) dataStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubActionsBillingOrgReq) validStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubActionsBillingOrgReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *BillingGetGithubActionsBillingOrgReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *BillingGetGithubActionsBillingOrgReq) Rel(link RelName, resp *BillingGetGithubActionsBillingOrgResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
BillingGetGithubActionsBillingOrgResponse is a response for BillingGetGithubActionsBillingOrg

https://developer.github.com/v3/billing/#get-github-actions-billing-for-an-organization
*/
type BillingGetGithubActionsBillingOrgResponse struct {
	response
	request *BillingGetGithubActionsBillingOrgReq
	Data    components.ActionsBillingUsage
}

/*
BillingGetGithubActionsBillingUser performs requests for "billing/get-github-actions-billing-user"

Get GitHub Actions billing for a user.

  GET /users/{username}/settings/billing/actions

https://developer.github.com/v3/billing/#get-github-actions-billing-for-a-user
*/
func BillingGetGithubActionsBillingUser(ctx context.Context, req *BillingGetGithubActionsBillingUserReq, opt ...RequestOption) (*BillingGetGithubActionsBillingUserResponse, error) {
	if req == nil {
		req = new(BillingGetGithubActionsBillingUserReq)
	}
	resp := &BillingGetGithubActionsBillingUserResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.ActionsBillingUsage{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
BillingGetGithubActionsBillingUser performs requests for "billing/get-github-actions-billing-user"

Get GitHub Actions billing for a user.

  GET /users/{username}/settings/billing/actions

https://developer.github.com/v3/billing/#get-github-actions-billing-for-a-user
*/
func (c Client) BillingGetGithubActionsBillingUser(ctx context.Context, req *BillingGetGithubActionsBillingUserReq, opt ...RequestOption) (*BillingGetGithubActionsBillingUserResponse, error) {
	return BillingGetGithubActionsBillingUser(ctx, req, append(c, opt...)...)
}

/*
BillingGetGithubActionsBillingUserReq is request data for Client.BillingGetGithubActionsBillingUser

https://developer.github.com/v3/billing/#get-github-actions-billing-for-a-user
*/
type BillingGetGithubActionsBillingUserReq struct {
	_url     string
	Username string
}

func (r *BillingGetGithubActionsBillingUserReq) url() string {
	return r._url
}

func (r *BillingGetGithubActionsBillingUserReq) urlPath() string {
	return fmt.Sprintf("/users/%v/settings/billing/actions", r.Username)
}

func (r *BillingGetGithubActionsBillingUserReq) method() string {
	return "GET"
}

func (r *BillingGetGithubActionsBillingUserReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *BillingGetGithubActionsBillingUserReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *BillingGetGithubActionsBillingUserReq) body() interface{} {
	return nil
}

func (r *BillingGetGithubActionsBillingUserReq) dataStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubActionsBillingUserReq) validStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubActionsBillingUserReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *BillingGetGithubActionsBillingUserReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *BillingGetGithubActionsBillingUserReq) Rel(link RelName, resp *BillingGetGithubActionsBillingUserResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
BillingGetGithubActionsBillingUserResponse is a response for BillingGetGithubActionsBillingUser

https://developer.github.com/v3/billing/#get-github-actions-billing-for-a-user
*/
type BillingGetGithubActionsBillingUserResponse struct {
	response
	request *BillingGetGithubActionsBillingUserReq
	Data    components.ActionsBillingUsage
}

/*
BillingGetGithubPackagesBillingGhe performs requests for "billing/get-github-packages-billing-ghe"

Get GitHub Packages billing for an enterprise.

  GET /enterprises/{enterprise_id}/settings/billing/packages

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-enterprise
*/
func BillingGetGithubPackagesBillingGhe(ctx context.Context, req *BillingGetGithubPackagesBillingGheReq, opt ...RequestOption) (*BillingGetGithubPackagesBillingGheResponse, error) {
	if req == nil {
		req = new(BillingGetGithubPackagesBillingGheReq)
	}
	resp := &BillingGetGithubPackagesBillingGheResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.PackagesBillingUsage{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
BillingGetGithubPackagesBillingGhe performs requests for "billing/get-github-packages-billing-ghe"

Get GitHub Packages billing for an enterprise.

  GET /enterprises/{enterprise_id}/settings/billing/packages

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-enterprise
*/
func (c Client) BillingGetGithubPackagesBillingGhe(ctx context.Context, req *BillingGetGithubPackagesBillingGheReq, opt ...RequestOption) (*BillingGetGithubPackagesBillingGheResponse, error) {
	return BillingGetGithubPackagesBillingGhe(ctx, req, append(c, opt...)...)
}

/*
BillingGetGithubPackagesBillingGheReq is request data for Client.BillingGetGithubPackagesBillingGhe

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-enterprise
*/
type BillingGetGithubPackagesBillingGheReq struct {
	_url string

	// Unique identifier of the GitHub Enterprise Cloud instance.
	EnterpriseId string
}

func (r *BillingGetGithubPackagesBillingGheReq) url() string {
	return r._url
}

func (r *BillingGetGithubPackagesBillingGheReq) urlPath() string {
	return fmt.Sprintf("/enterprises/%v/settings/billing/packages", r.EnterpriseId)
}

func (r *BillingGetGithubPackagesBillingGheReq) method() string {
	return "GET"
}

func (r *BillingGetGithubPackagesBillingGheReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *BillingGetGithubPackagesBillingGheReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *BillingGetGithubPackagesBillingGheReq) body() interface{} {
	return nil
}

func (r *BillingGetGithubPackagesBillingGheReq) dataStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubPackagesBillingGheReq) validStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubPackagesBillingGheReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *BillingGetGithubPackagesBillingGheReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *BillingGetGithubPackagesBillingGheReq) Rel(link RelName, resp *BillingGetGithubPackagesBillingGheResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
BillingGetGithubPackagesBillingGheResponse is a response for BillingGetGithubPackagesBillingGhe

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-enterprise
*/
type BillingGetGithubPackagesBillingGheResponse struct {
	response
	request *BillingGetGithubPackagesBillingGheReq
	Data    components.PackagesBillingUsage
}

/*
BillingGetGithubPackagesBillingOrg performs requests for "billing/get-github-packages-billing-org"

Get GitHub Packages billing for an organization.

  GET /orgs/{org}/settings/billing/packages

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-organization
*/
func BillingGetGithubPackagesBillingOrg(ctx context.Context, req *BillingGetGithubPackagesBillingOrgReq, opt ...RequestOption) (*BillingGetGithubPackagesBillingOrgResponse, error) {
	if req == nil {
		req = new(BillingGetGithubPackagesBillingOrgReq)
	}
	resp := &BillingGetGithubPackagesBillingOrgResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.PackagesBillingUsage{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
BillingGetGithubPackagesBillingOrg performs requests for "billing/get-github-packages-billing-org"

Get GitHub Packages billing for an organization.

  GET /orgs/{org}/settings/billing/packages

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-organization
*/
func (c Client) BillingGetGithubPackagesBillingOrg(ctx context.Context, req *BillingGetGithubPackagesBillingOrgReq, opt ...RequestOption) (*BillingGetGithubPackagesBillingOrgResponse, error) {
	return BillingGetGithubPackagesBillingOrg(ctx, req, append(c, opt...)...)
}

/*
BillingGetGithubPackagesBillingOrgReq is request data for Client.BillingGetGithubPackagesBillingOrg

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-organization
*/
type BillingGetGithubPackagesBillingOrgReq struct {
	_url string
	Org  string
}

func (r *BillingGetGithubPackagesBillingOrgReq) url() string {
	return r._url
}

func (r *BillingGetGithubPackagesBillingOrgReq) urlPath() string {
	return fmt.Sprintf("/orgs/%v/settings/billing/packages", r.Org)
}

func (r *BillingGetGithubPackagesBillingOrgReq) method() string {
	return "GET"
}

func (r *BillingGetGithubPackagesBillingOrgReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *BillingGetGithubPackagesBillingOrgReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *BillingGetGithubPackagesBillingOrgReq) body() interface{} {
	return nil
}

func (r *BillingGetGithubPackagesBillingOrgReq) dataStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubPackagesBillingOrgReq) validStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubPackagesBillingOrgReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *BillingGetGithubPackagesBillingOrgReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *BillingGetGithubPackagesBillingOrgReq) Rel(link RelName, resp *BillingGetGithubPackagesBillingOrgResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
BillingGetGithubPackagesBillingOrgResponse is a response for BillingGetGithubPackagesBillingOrg

https://developer.github.com/v3/billing/#get-github-packages-billing-for-an-organization
*/
type BillingGetGithubPackagesBillingOrgResponse struct {
	response
	request *BillingGetGithubPackagesBillingOrgReq
	Data    components.PackagesBillingUsage
}

/*
BillingGetGithubPackagesBillingUser performs requests for "billing/get-github-packages-billing-user"

Get GitHub Packages billing for a user.

  GET /users/{username}/settings/billing/packages

https://developer.github.com/v3/billing/#get-github-packages-billing-for-a-user
*/
func BillingGetGithubPackagesBillingUser(ctx context.Context, req *BillingGetGithubPackagesBillingUserReq, opt ...RequestOption) (*BillingGetGithubPackagesBillingUserResponse, error) {
	if req == nil {
		req = new(BillingGetGithubPackagesBillingUserReq)
	}
	resp := &BillingGetGithubPackagesBillingUserResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.PackagesBillingUsage{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
BillingGetGithubPackagesBillingUser performs requests for "billing/get-github-packages-billing-user"

Get GitHub Packages billing for a user.

  GET /users/{username}/settings/billing/packages

https://developer.github.com/v3/billing/#get-github-packages-billing-for-a-user
*/
func (c Client) BillingGetGithubPackagesBillingUser(ctx context.Context, req *BillingGetGithubPackagesBillingUserReq, opt ...RequestOption) (*BillingGetGithubPackagesBillingUserResponse, error) {
	return BillingGetGithubPackagesBillingUser(ctx, req, append(c, opt...)...)
}

/*
BillingGetGithubPackagesBillingUserReq is request data for Client.BillingGetGithubPackagesBillingUser

https://developer.github.com/v3/billing/#get-github-packages-billing-for-a-user
*/
type BillingGetGithubPackagesBillingUserReq struct {
	_url     string
	Username string
}

func (r *BillingGetGithubPackagesBillingUserReq) url() string {
	return r._url
}

func (r *BillingGetGithubPackagesBillingUserReq) urlPath() string {
	return fmt.Sprintf("/users/%v/settings/billing/packages", r.Username)
}

func (r *BillingGetGithubPackagesBillingUserReq) method() string {
	return "GET"
}

func (r *BillingGetGithubPackagesBillingUserReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *BillingGetGithubPackagesBillingUserReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *BillingGetGithubPackagesBillingUserReq) body() interface{} {
	return nil
}

func (r *BillingGetGithubPackagesBillingUserReq) dataStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubPackagesBillingUserReq) validStatuses() []int {
	return []int{200}
}

func (r *BillingGetGithubPackagesBillingUserReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *BillingGetGithubPackagesBillingUserReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *BillingGetGithubPackagesBillingUserReq) Rel(link RelName, resp *BillingGetGithubPackagesBillingUserResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
BillingGetGithubPackagesBillingUserResponse is a response for BillingGetGithubPackagesBillingUser

https://developer.github.com/v3/billing/#get-github-packages-billing-for-a-user
*/
type BillingGetGithubPackagesBillingUserResponse struct {
	response
	request *BillingGetGithubPackagesBillingUserReq
	Data    components.PackagesBillingUsage
}

/*
BillingGetSharedStorageBillingGhe performs requests for "billing/get-shared-storage-billing-ghe"

Get shared storage billing for an enterprise.

  GET /enterprises/{enterprise_id}/settings/billing/shared-storage

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-enterprise
*/
func BillingGetSharedStorageBillingGhe(ctx context.Context, req *BillingGetSharedStorageBillingGheReq, opt ...RequestOption) (*BillingGetSharedStorageBillingGheResponse, error) {
	if req == nil {
		req = new(BillingGetSharedStorageBillingGheReq)
	}
	resp := &BillingGetSharedStorageBillingGheResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.CombinedBillingUsage{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
BillingGetSharedStorageBillingGhe performs requests for "billing/get-shared-storage-billing-ghe"

Get shared storage billing for an enterprise.

  GET /enterprises/{enterprise_id}/settings/billing/shared-storage

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-enterprise
*/
func (c Client) BillingGetSharedStorageBillingGhe(ctx context.Context, req *BillingGetSharedStorageBillingGheReq, opt ...RequestOption) (*BillingGetSharedStorageBillingGheResponse, error) {
	return BillingGetSharedStorageBillingGhe(ctx, req, append(c, opt...)...)
}

/*
BillingGetSharedStorageBillingGheReq is request data for Client.BillingGetSharedStorageBillingGhe

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-enterprise
*/
type BillingGetSharedStorageBillingGheReq struct {
	_url string

	// Unique identifier of the GitHub Enterprise Cloud instance.
	EnterpriseId string
}

func (r *BillingGetSharedStorageBillingGheReq) url() string {
	return r._url
}

func (r *BillingGetSharedStorageBillingGheReq) urlPath() string {
	return fmt.Sprintf("/enterprises/%v/settings/billing/shared-storage", r.EnterpriseId)
}

func (r *BillingGetSharedStorageBillingGheReq) method() string {
	return "GET"
}

func (r *BillingGetSharedStorageBillingGheReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *BillingGetSharedStorageBillingGheReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *BillingGetSharedStorageBillingGheReq) body() interface{} {
	return nil
}

func (r *BillingGetSharedStorageBillingGheReq) dataStatuses() []int {
	return []int{200}
}

func (r *BillingGetSharedStorageBillingGheReq) validStatuses() []int {
	return []int{200}
}

func (r *BillingGetSharedStorageBillingGheReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *BillingGetSharedStorageBillingGheReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *BillingGetSharedStorageBillingGheReq) Rel(link RelName, resp *BillingGetSharedStorageBillingGheResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
BillingGetSharedStorageBillingGheResponse is a response for BillingGetSharedStorageBillingGhe

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-enterprise
*/
type BillingGetSharedStorageBillingGheResponse struct {
	response
	request *BillingGetSharedStorageBillingGheReq
	Data    components.CombinedBillingUsage
}

/*
BillingGetSharedStorageBillingOrg performs requests for "billing/get-shared-storage-billing-org"

Get shared storage billing for an organization.

  GET /orgs/{org}/settings/billing/shared-storage

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-organization
*/
func BillingGetSharedStorageBillingOrg(ctx context.Context, req *BillingGetSharedStorageBillingOrgReq, opt ...RequestOption) (*BillingGetSharedStorageBillingOrgResponse, error) {
	if req == nil {
		req = new(BillingGetSharedStorageBillingOrgReq)
	}
	resp := &BillingGetSharedStorageBillingOrgResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.CombinedBillingUsage{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
BillingGetSharedStorageBillingOrg performs requests for "billing/get-shared-storage-billing-org"

Get shared storage billing for an organization.

  GET /orgs/{org}/settings/billing/shared-storage

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-organization
*/
func (c Client) BillingGetSharedStorageBillingOrg(ctx context.Context, req *BillingGetSharedStorageBillingOrgReq, opt ...RequestOption) (*BillingGetSharedStorageBillingOrgResponse, error) {
	return BillingGetSharedStorageBillingOrg(ctx, req, append(c, opt...)...)
}

/*
BillingGetSharedStorageBillingOrgReq is request data for Client.BillingGetSharedStorageBillingOrg

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-organization
*/
type BillingGetSharedStorageBillingOrgReq struct {
	_url string
	Org  string
}

func (r *BillingGetSharedStorageBillingOrgReq) url() string {
	return r._url
}

func (r *BillingGetSharedStorageBillingOrgReq) urlPath() string {
	return fmt.Sprintf("/orgs/%v/settings/billing/shared-storage", r.Org)
}

func (r *BillingGetSharedStorageBillingOrgReq) method() string {
	return "GET"
}

func (r *BillingGetSharedStorageBillingOrgReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *BillingGetSharedStorageBillingOrgReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *BillingGetSharedStorageBillingOrgReq) body() interface{} {
	return nil
}

func (r *BillingGetSharedStorageBillingOrgReq) dataStatuses() []int {
	return []int{200}
}

func (r *BillingGetSharedStorageBillingOrgReq) validStatuses() []int {
	return []int{200}
}

func (r *BillingGetSharedStorageBillingOrgReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *BillingGetSharedStorageBillingOrgReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *BillingGetSharedStorageBillingOrgReq) Rel(link RelName, resp *BillingGetSharedStorageBillingOrgResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
BillingGetSharedStorageBillingOrgResponse is a response for BillingGetSharedStorageBillingOrg

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-an-organization
*/
type BillingGetSharedStorageBillingOrgResponse struct {
	response
	request *BillingGetSharedStorageBillingOrgReq
	Data    components.CombinedBillingUsage
}

/*
BillingGetSharedStorageBillingUser performs requests for "billing/get-shared-storage-billing-user"

Get shared storage billing for a user.

  GET /users/{username}/settings/billing/shared-storage

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-a-user
*/
func BillingGetSharedStorageBillingUser(ctx context.Context, req *BillingGetSharedStorageBillingUserReq, opt ...RequestOption) (*BillingGetSharedStorageBillingUserResponse, error) {
	if req == nil {
		req = new(BillingGetSharedStorageBillingUserReq)
	}
	resp := &BillingGetSharedStorageBillingUserResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = components.CombinedBillingUsage{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
BillingGetSharedStorageBillingUser performs requests for "billing/get-shared-storage-billing-user"

Get shared storage billing for a user.

  GET /users/{username}/settings/billing/shared-storage

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-a-user
*/
func (c Client) BillingGetSharedStorageBillingUser(ctx context.Context, req *BillingGetSharedStorageBillingUserReq, opt ...RequestOption) (*BillingGetSharedStorageBillingUserResponse, error) {
	return BillingGetSharedStorageBillingUser(ctx, req, append(c, opt...)...)
}

/*
BillingGetSharedStorageBillingUserReq is request data for Client.BillingGetSharedStorageBillingUser

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-a-user
*/
type BillingGetSharedStorageBillingUserReq struct {
	_url     string
	Username string
}

func (r *BillingGetSharedStorageBillingUserReq) url() string {
	return r._url
}

func (r *BillingGetSharedStorageBillingUserReq) urlPath() string {
	return fmt.Sprintf("/users/%v/settings/billing/shared-storage", r.Username)
}

func (r *BillingGetSharedStorageBillingUserReq) method() string {
	return "GET"
}

func (r *BillingGetSharedStorageBillingUserReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *BillingGetSharedStorageBillingUserReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *BillingGetSharedStorageBillingUserReq) body() interface{} {
	return nil
}

func (r *BillingGetSharedStorageBillingUserReq) dataStatuses() []int {
	return []int{200}
}

func (r *BillingGetSharedStorageBillingUserReq) validStatuses() []int {
	return []int{200}
}

func (r *BillingGetSharedStorageBillingUserReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *BillingGetSharedStorageBillingUserReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *BillingGetSharedStorageBillingUserReq) Rel(link RelName, resp *BillingGetSharedStorageBillingUserResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
BillingGetSharedStorageBillingUserResponse is a response for BillingGetSharedStorageBillingUser

https://developer.github.com/v3/billing/#get-shared-storage-billing-for-a-user
*/
type BillingGetSharedStorageBillingUserResponse struct {
	response
	request *BillingGetSharedStorageBillingUserReq
	Data    components.CombinedBillingUsage
}
