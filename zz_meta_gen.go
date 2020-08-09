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
MetaGet performs requests for "meta/get"

Get GitHub meta information.

  GET /meta

https://developer.github.com/v3/meta/#get-github-meta-information
*/
func MetaGet(ctx context.Context, req *MetaGetReq, opt ...options.Option) (*MetaGetResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(MetaGetReq)
	}
	resp := &MetaGetResponse{request: req}
	r, err := internal.DoRequest(ctx, req.requestBuilder(), opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = components.ApiOverview{}
	err = internal.DecodeResponseBody(r, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
MetaGet performs requests for "meta/get"

Get GitHub meta information.

  GET /meta

https://developer.github.com/v3/meta/#get-github-meta-information
*/
func (c Client) MetaGet(ctx context.Context, req *MetaGetReq, opt ...options.Option) (*MetaGetResponse, error) {
	return MetaGet(ctx, req, append(c, opt...)...)
}

/*
MetaGetReq is request data for Client.MetaGet

https://developer.github.com/v3/meta/#get-github-meta-information
*/
type MetaGetReq struct {
	_url string
}

// HTTPRequest builds an *http.Request
func (r *MetaGetReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *MetaGetReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": String("application/json")},
		Method:           "GET",
		OperationID:      "meta/get",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/meta"),
		URLQuery:         query,
		ValidStatuses:    []int{200, 304},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *MetaGetReq) Rel(link RelName, resp *MetaGetResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
MetaGetResponse is a response for MetaGet

https://developer.github.com/v3/meta/#get-github-meta-information
*/
type MetaGetResponse struct {
	internal.Response
	request *MetaGetReq
	Data    components.ApiOverview
}

/*
MetaGetOctocat performs requests for "meta/get-octocat"

Get Octocat.

  GET /octocat

*/
func MetaGetOctocat(ctx context.Context, req *MetaGetOctocatReq, opt ...options.Option) (*MetaGetOctocatResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(MetaGetOctocatReq)
	}
	resp := &MetaGetOctocatResponse{request: req}
	r, err := internal.DoRequest(ctx, req.requestBuilder(), opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	err = internal.DecodeResponseBody(r, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
MetaGetOctocat performs requests for "meta/get-octocat"

Get Octocat.

  GET /octocat

*/
func (c Client) MetaGetOctocat(ctx context.Context, req *MetaGetOctocatReq, opt ...options.Option) (*MetaGetOctocatResponse, error) {
	return MetaGetOctocat(ctx, req, append(c, opt...)...)
}

/*
MetaGetOctocatReq is request data for Client.MetaGetOctocat

*/
type MetaGetOctocatReq struct {
	_url string

	// The words to show in Octocat's speech bubble
	S *string
}

// HTTPRequest builds an *http.Request
func (r *MetaGetOctocatReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *MetaGetOctocatReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}
	if r.S != nil {
		query.Set("s", *r.S)
	}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{},
		Method:           "GET",
		OperationID:      "meta/get-octocat",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/octocat"),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *MetaGetOctocatReq) Rel(link RelName, resp *MetaGetOctocatResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
MetaGetOctocatResponse is a response for MetaGetOctocat

*/
type MetaGetOctocatResponse struct {
	internal.Response
	request *MetaGetOctocatReq
}

/*
MetaGetZen performs requests for "meta/get-zen"

Get the Zen of GitHub.

  GET /zen

*/
func MetaGetZen(ctx context.Context, req *MetaGetZenReq, opt ...options.Option) (*MetaGetZenResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(MetaGetZenReq)
	}
	resp := &MetaGetZenResponse{request: req}
	r, err := internal.DoRequest(ctx, req.requestBuilder(), opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	err = internal.DecodeResponseBody(r, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
MetaGetZen performs requests for "meta/get-zen"

Get the Zen of GitHub.

  GET /zen

*/
func (c Client) MetaGetZen(ctx context.Context, req *MetaGetZenReq, opt ...options.Option) (*MetaGetZenResponse, error) {
	return MetaGetZen(ctx, req, append(c, opt...)...)
}

/*
MetaGetZenReq is request data for Client.MetaGetZen

*/
type MetaGetZenReq struct {
	_url string
}

// HTTPRequest builds an *http.Request
func (r *MetaGetZenReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *MetaGetZenReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{},
		Method:           "GET",
		OperationID:      "meta/get-zen",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/zen"),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *MetaGetZenReq) Rel(link RelName, resp *MetaGetZenResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
MetaGetZenResponse is a response for MetaGetZen

*/
type MetaGetZenResponse struct {
	internal.Response
	request *MetaGetZenReq
}

/*
MetaRoot performs requests for "meta/root"

GitHub API Root.

  GET /

*/
func MetaRoot(ctx context.Context, req *MetaRootReq, opt ...options.Option) (*MetaRootResponse, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	if req == nil {
		req = new(MetaRootReq)
	}
	resp := &MetaRootResponse{request: req}
	r, err := internal.DoRequest(ctx, req.requestBuilder(), opts)

	if r != nil {
		resp.Response = *r
	}
	if err != nil {
		return resp, err
	}

	resp.Data = MetaRootResponseBody{}
	err = internal.DecodeResponseBody(r, &resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
MetaRoot performs requests for "meta/root"

GitHub API Root.

  GET /

*/
func (c Client) MetaRoot(ctx context.Context, req *MetaRootReq, opt ...options.Option) (*MetaRootResponse, error) {
	return MetaRoot(ctx, req, append(c, opt...)...)
}

/*
MetaRootReq is request data for Client.MetaRoot

*/
type MetaRootReq struct {
	_url string
}

// HTTPRequest builds an *http.Request
func (r *MetaRootReq) HTTPRequest(ctx context.Context, opt ...options.Option) (*http.Request, error) {
	opts, err := options.BuildOptions(opt...)
	if err != nil {
		return nil, err
	}
	return r.requestBuilder().HTTPRequest(ctx, opts)
}

func (r *MetaRootReq) requestBuilder() *internal.RequestBuilder {
	query := url.Values{}

	builder := &internal.RequestBuilder{
		AllPreviews:      []string{},
		Body:             nil,
		DataStatuses:     []int{200},
		ExplicitURL:      r._url,
		HeaderVals:       map[string]*string{"accept": String("application/json")},
		Method:           "GET",
		OperationID:      "meta/root",
		Previews:         map[string]bool{},
		RequiredPreviews: []string{},
		URLPath:          fmt.Sprintf("/"),
		URLQuery:         query,
		ValidStatuses:    []int{200},
	}
	return builder
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *MetaRootReq) Rel(link RelName, resp *MetaRootResponse) bool {
	u := resp.RelLink(string(link))
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
MetaRootResponseBody is a response body for MetaRoot

*/
type MetaRootResponseBody struct {
	AuthorizationsUrl                string `json:"authorizations_url"`
	CodeSearchUrl                    string `json:"code_search_url"`
	CommitSearchUrl                  string `json:"commit_search_url"`
	CurrentUserAuthorizationsHtmlUrl string `json:"current_user_authorizations_html_url"`
	CurrentUserRepositoriesUrl       string `json:"current_user_repositories_url"`
	CurrentUserUrl                   string `json:"current_user_url"`
	EmailsUrl                        string `json:"emails_url"`
	EmojisUrl                        string `json:"emojis_url"`
	EventsUrl                        string `json:"events_url"`
	FeedsUrl                         string `json:"feeds_url"`
	FollowersUrl                     string `json:"followers_url"`
	FollowingUrl                     string `json:"following_url"`
	GistsUrl                         string `json:"gists_url"`
	HubUrl                           string `json:"hub_url"`
	IssueSearchUrl                   string `json:"issue_search_url"`
	IssuesUrl                        string `json:"issues_url"`
	KeysUrl                          string `json:"keys_url"`
	LabelSearchUrl                   string `json:"label_search_url"`
	NotificationsUrl                 string `json:"notifications_url"`
	OrganizationRepositoriesUrl      string `json:"organization_repositories_url"`
	OrganizationTeamsUrl             string `json:"organization_teams_url"`
	OrganizationUrl                  string `json:"organization_url"`
	PublicGistsUrl                   string `json:"public_gists_url"`
	RateLimitUrl                     string `json:"rate_limit_url"`
	RepositorySearchUrl              string `json:"repository_search_url"`
	RepositoryUrl                    string `json:"repository_url"`
	StarredGistsUrl                  string `json:"starred_gists_url"`
	StarredUrl                       string `json:"starred_url"`
	TopicSearchUrl                   string `json:"topic_search_url,omitempty"`
	UserOrganizationsUrl             string `json:"user_organizations_url"`
	UserRepositoriesUrl              string `json:"user_repositories_url"`
	UserSearchUrl                    string `json:"user_search_url"`
	UserUrl                          string `json:"user_url"`
}

/*
MetaRootResponse is a response for MetaRoot

*/
type MetaRootResponse struct {
	internal.Response
	request *MetaRootReq
	Data    MetaRootResponseBody
}
