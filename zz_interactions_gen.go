// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

/*
InteractionsAddOrUpdateRestrictionsForOrgReq builds requests for "interactions/add-or-update-restrictions-for-org"

Add or update interaction restrictions for an organization.

  PUT /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#add-or-update-interaction-restrictions-for-an-organization
*/
type InteractionsAddOrUpdateRestrictionsForOrgReq struct {
	Org         string
	RequestBody InteractionsAddOrUpdateRestrictionsForOrgReqBody

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

func (r InteractionsAddOrUpdateRestrictionsForOrgReq) urlPath() string {
	return fmt.Sprintf("/orgs/%v/interaction-limits", r.Org)
}

func (r InteractionsAddOrUpdateRestrictionsForOrgReq) method() string {
	return "PUT"
}

func (r InteractionsAddOrUpdateRestrictionsForOrgReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r InteractionsAddOrUpdateRestrictionsForOrgReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{"sombra": r.SombraPreview}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r InteractionsAddOrUpdateRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), r.RequestBody, opt)
}

/*
InteractionsAddOrUpdateRestrictionsForOrgReqBody is a request body for interactions/add-or-update-restrictions-for-org

API documentation: https://developer.github.com/v3/interactions/orgs/#add-or-update-interaction-restrictions-for-an-organization
*/
type InteractionsAddOrUpdateRestrictionsForOrgReqBody struct {

	/*
	   Specifies the group of GitHub users who can comment, open issues, or create pull
	   requests in public repositories for the given organization. Must be one of:
	   `existing_users`, `contributors_only`, or `collaborators_only`.
	*/
	Limit *string `json:"limit"`
}

/*
InteractionsAddOrUpdateRestrictionsForOrgResponseBody200 is a response body for interactions/add-or-update-restrictions-for-org

API documentation: https://developer.github.com/v3/interactions/orgs/#add-or-update-interaction-restrictions-for-an-organization
*/
type InteractionsAddOrUpdateRestrictionsForOrgResponseBody200 struct {
	ExpiresAt string `json:"expires_at,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Origin    string `json:"origin,omitempty"`
}

/*
InteractionsAddOrUpdateRestrictionsForRepoReq builds requests for "interactions/add-or-update-restrictions-for-repo"

Add or update interaction restrictions for a repository.

  PUT /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#add-or-update-interaction-restrictions-for-a-repository
*/
type InteractionsAddOrUpdateRestrictionsForRepoReq struct {
	Owner       string
	Repo        string
	RequestBody InteractionsAddOrUpdateRestrictionsForRepoReqBody

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

func (r InteractionsAddOrUpdateRestrictionsForRepoReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo)
}

func (r InteractionsAddOrUpdateRestrictionsForRepoReq) method() string {
	return "PUT"
}

func (r InteractionsAddOrUpdateRestrictionsForRepoReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r InteractionsAddOrUpdateRestrictionsForRepoReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{"sombra": r.SombraPreview}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r InteractionsAddOrUpdateRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), r.RequestBody, opt)
}

/*
InteractionsAddOrUpdateRestrictionsForRepoReqBody is a request body for interactions/add-or-update-restrictions-for-repo

API documentation: https://developer.github.com/v3/interactions/repos/#add-or-update-interaction-restrictions-for-a-repository
*/
type InteractionsAddOrUpdateRestrictionsForRepoReqBody struct {

	/*
	   Specifies the group of GitHub users who can comment, open issues, or create pull
	   requests for the given repository. Must be one of: `existing_users`,
	   `contributors_only`, or `collaborators_only`.
	*/
	Limit *string `json:"limit"`
}

/*
InteractionsAddOrUpdateRestrictionsForRepoResponseBody200 is a response body for interactions/add-or-update-restrictions-for-repo

API documentation: https://developer.github.com/v3/interactions/repos/#add-or-update-interaction-restrictions-for-a-repository
*/
type InteractionsAddOrUpdateRestrictionsForRepoResponseBody200 struct {
	ExpiresAt string `json:"expires_at,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Origin    string `json:"origin,omitempty"`
}

/*
InteractionsGetRestrictionsForOrgReq builds requests for "interactions/get-restrictions-for-org"

Get interaction restrictions for an organization.

  GET /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
type InteractionsGetRestrictionsForOrgReq struct {
	Org string

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

func (r InteractionsGetRestrictionsForOrgReq) urlPath() string {
	return fmt.Sprintf("/orgs/%v/interaction-limits", r.Org)
}

func (r InteractionsGetRestrictionsForOrgReq) method() string {
	return "GET"
}

func (r InteractionsGetRestrictionsForOrgReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r InteractionsGetRestrictionsForOrgReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{"sombra": r.SombraPreview}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r InteractionsGetRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
InteractionsGetRestrictionsForOrgResponseBody200 is a response body for interactions/get-restrictions-for-org

API documentation: https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization
*/
type InteractionsGetRestrictionsForOrgResponseBody200 struct {
	ExpiresAt string `json:"expires_at,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Origin    string `json:"origin,omitempty"`
}

/*
InteractionsGetRestrictionsForRepoReq builds requests for "interactions/get-restrictions-for-repo"

Get interaction restrictions for a repository.

  GET /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
type InteractionsGetRestrictionsForRepoReq struct {
	Owner string
	Repo  string

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

func (r InteractionsGetRestrictionsForRepoReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo)
}

func (r InteractionsGetRestrictionsForRepoReq) method() string {
	return "GET"
}

func (r InteractionsGetRestrictionsForRepoReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r InteractionsGetRestrictionsForRepoReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{"sombra": r.SombraPreview}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r InteractionsGetRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
InteractionsGetRestrictionsForRepoResponseBody200 is a response body for interactions/get-restrictions-for-repo

API documentation: https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository
*/
type InteractionsGetRestrictionsForRepoResponseBody200 struct {
	ExpiresAt string `json:"expires_at,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Origin    string `json:"origin,omitempty"`
}

/*
InteractionsRemoveRestrictionsForOrgReq builds requests for "interactions/remove-restrictions-for-org"

Remove interaction restrictions for an organization.

  DELETE /orgs/{org}/interaction-limits

https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization
*/
type InteractionsRemoveRestrictionsForOrgReq struct {
	Org string

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

func (r InteractionsRemoveRestrictionsForOrgReq) urlPath() string {
	return fmt.Sprintf("/orgs/%v/interaction-limits", r.Org)
}

func (r InteractionsRemoveRestrictionsForOrgReq) method() string {
	return "DELETE"
}

func (r InteractionsRemoveRestrictionsForOrgReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r InteractionsRemoveRestrictionsForOrgReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{"sombra": r.SombraPreview}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r InteractionsRemoveRestrictionsForOrgReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}

/*
InteractionsRemoveRestrictionsForRepoReq builds requests for "interactions/remove-restrictions-for-repo"

Remove interaction restrictions for a repository.

  DELETE /repos/{owner}/{repo}/interaction-limits

https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository
*/
type InteractionsRemoveRestrictionsForRepoReq struct {
	Owner string
	Repo  string

	/*
	The Interactions API is currently in public preview. See the [blog
	post](https://developer.github.com/changes/2018-12-18-interactions-preview)
	preview for more details. To access the API during the preview period, you must
	set this to true.
	*/
	SombraPreview bool
}

func (r InteractionsRemoveRestrictionsForRepoReq) urlPath() string {
	return fmt.Sprintf("/repos/%v/%v/interaction-limits", r.Owner, r.Repo)
}

func (r InteractionsRemoveRestrictionsForRepoReq) method() string {
	return "DELETE"
}

func (r InteractionsRemoveRestrictionsForRepoReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r InteractionsRemoveRestrictionsForRepoReq) header() http.Header {
	headerVals := map[string]*string{}
	previewVals := map[string]bool{"sombra": r.SombraPreview}
	return requestHeaders(headerVals, previewVals)
}

// HTTPRequest creates an http request
func (r InteractionsRemoveRestrictionsForRepoReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return httpRequest(ctx, r.urlPath(), r.method(), r.urlQuery(), r.header(), nil, opt)
}
