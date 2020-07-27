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
GitignoreGetAllTemplates performs requests for "gitignore/get-all-templates"

Get all gitignore templates.

  GET /gitignore/templates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
func GitignoreGetAllTemplates(ctx context.Context, req *GitignoreGetAllTemplatesReq, opt ...RequestOption) (*GitignoreGetAllTemplatesResponse, error) {
	if req == nil {
		req = new(GitignoreGetAllTemplatesReq)
	}
	resp := &GitignoreGetAllTemplatesResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = GitignoreGetAllTemplatesResponseBody{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
GitignoreGetAllTemplates performs requests for "gitignore/get-all-templates"

Get all gitignore templates.

  GET /gitignore/templates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
func (c Client) GitignoreGetAllTemplates(ctx context.Context, req *GitignoreGetAllTemplatesReq, opt ...RequestOption) (*GitignoreGetAllTemplatesResponse, error) {
	return GitignoreGetAllTemplates(ctx, req, append(c, opt...)...)
}

/*
GitignoreGetAllTemplatesReq is request data for Client.GitignoreGetAllTemplates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
type GitignoreGetAllTemplatesReq struct {
	_url string
}

func (r *GitignoreGetAllTemplatesReq) url() string {
	return r._url
}

func (r *GitignoreGetAllTemplatesReq) urlPath() string {
	return fmt.Sprintf("/gitignore/templates")
}

func (r *GitignoreGetAllTemplatesReq) method() string {
	return "GET"
}

func (r *GitignoreGetAllTemplatesReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *GitignoreGetAllTemplatesReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *GitignoreGetAllTemplatesReq) body() interface{} {
	return nil
}

func (r *GitignoreGetAllTemplatesReq) dataStatuses() []int {
	return []int{200}
}

func (r *GitignoreGetAllTemplatesReq) validStatuses() []int {
	return []int{200, 304}
}

func (r *GitignoreGetAllTemplatesReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *GitignoreGetAllTemplatesReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GitignoreGetAllTemplatesReq) Rel(link RelName, resp *GitignoreGetAllTemplatesResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GitignoreGetAllTemplatesResponseBody is a response body for GitignoreGetAllTemplates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
type GitignoreGetAllTemplatesResponseBody []string

/*
GitignoreGetAllTemplatesResponse is a response for GitignoreGetAllTemplates

https://developer.github.com/v3/gitignore/#get-all-gitignore-templates
*/
type GitignoreGetAllTemplatesResponse struct {
	response
	request *GitignoreGetAllTemplatesReq
	Data    GitignoreGetAllTemplatesResponseBody
}

/*
GitignoreGetTemplate performs requests for "gitignore/get-template"

Get a gitignore template.

  GET /gitignore/templates/{name}

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
func GitignoreGetTemplate(ctx context.Context, req *GitignoreGetTemplateReq, opt ...RequestOption) (*GitignoreGetTemplateResponse, error) {
	if req == nil {
		req = new(GitignoreGetTemplateReq)
	}
	resp := &GitignoreGetTemplateResponse{request: req}
	r, err := doRequest(ctx, req, opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = GitignoreGetTemplateResponseBody{}
	err = r.decodeBody(&resp.Data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
GitignoreGetTemplate performs requests for "gitignore/get-template"

Get a gitignore template.

  GET /gitignore/templates/{name}

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
func (c Client) GitignoreGetTemplate(ctx context.Context, req *GitignoreGetTemplateReq, opt ...RequestOption) (*GitignoreGetTemplateResponse, error) {
	return GitignoreGetTemplate(ctx, req, append(c, opt...)...)
}

/*
GitignoreGetTemplateReq is request data for Client.GitignoreGetTemplate

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
type GitignoreGetTemplateReq struct {
	_url string

	// name parameter
	Name string
}

func (r *GitignoreGetTemplateReq) url() string {
	return r._url
}

func (r *GitignoreGetTemplateReq) urlPath() string {
	return fmt.Sprintf("/gitignore/templates/%v", r.Name)
}

func (r *GitignoreGetTemplateReq) method() string {
	return "GET"
}

func (r *GitignoreGetTemplateReq) urlQuery() url.Values {
	query := url.Values{}
	return query
}

func (r *GitignoreGetTemplateReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *GitignoreGetTemplateReq) body() interface{} {
	return nil
}

func (r *GitignoreGetTemplateReq) dataStatuses() []int {
	return []int{200}
}

func (r *GitignoreGetTemplateReq) validStatuses() []int {
	return []int{200, 304}
}

func (r *GitignoreGetTemplateReq) endpointAttributes() []endpointAttribute {
	return []endpointAttribute{}
}

// HTTPRequest builds an *http.Request
func (r *GitignoreGetTemplateReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *GitignoreGetTemplateReq) Rel(link RelName, resp *GitignoreGetTemplateResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
GitignoreGetTemplateResponseBody is a response body for GitignoreGetTemplate

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
type GitignoreGetTemplateResponseBody components.GitignoreTemplate

/*
GitignoreGetTemplateResponse is a response for GitignoreGetTemplate

https://developer.github.com/v3/gitignore/#get-a-gitignore-template
*/
type GitignoreGetTemplateResponse struct {
	response
	request *GitignoreGetTemplateReq
	Data    GitignoreGetTemplateResponseBody
}
