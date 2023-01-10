package bitbucket

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Version = "v1.0.0"

	defaultApiVersion = "latest"
	defaultUserAgent  = "go-bitbucket" + "/" + Version
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	ApiVersion string
	UserAgent  string

	common service

	AccessTokens *AccessTokensService
	Keys         *KeysService
	Projects     *ProjectsService
}

type service struct {
	client *Client
}

type Page struct {
	// The following properties support the paged APIs.
	Size     uint
	Limit    uint
	LastPage bool
	Start    uint
	// The next page start should be used with the ListOptions struct.
	NextPageStart uint
}

// Paged defines interface to be supported by responses from Paged APIs
type Paged interface {
	Current() *Page
}

// ListResponse defines the common properties of a list response
type ListResponse struct {
	Size          uint `json:"size"`
	Limit         uint `json:"limit"`
	LastPage      bool `json:"isLastPage"`
	Start         uint `json:"start"`
	NextPageStart uint `json:"nextPageStart"`
}

func (r *ListResponse) Current() *Page {
	return &Page{
		r.Size,
		r.Limit,
		r.LastPage,
		r.Start,
		r.NextPageStart,
	}
}

type ListOptions struct {
	Limit uint
	Start uint
}

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(bytes []byte) error {
	var raw int64
	err := json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}
	raw = raw / 1000
	*t = DateTime(time.Unix(raw, 0))
	return nil
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix() * 1000)
}

type Permission string

const (
	PermissionRepoRead      Permission = "REPO_READ"
	PermissionRepoWrite     Permission = "REPO_WRITE"
	PermissionRepoAdmin     Permission = "REPO_ADMIN"
	PermissionRepoCreate    Permission = "REPO_CREATE"
	PermissionProjectView   Permission = "PROJECT_VIEW"
	PermissionProjectRead   Permission = "PROJECT_READ"
	PermissionProjectWrite  Permission = "PROJECT_WRITE"
	PermissionProjectAdmin  Permission = "PROJECT_ADMIN"
	PermissionProjectCreate Permission = "PROJECT_CREATE"
	PermissionUserAdmin     Permission = "USER_ADMIN"
	PermissionLicensedUser  Permission = "LICENSED_USER"
	PermissionAdmin         Permission = "ADMIN"
	PermissionSysAdmin      Permission = "SYS_ADMIN"
)

type Response struct {
	*http.Response
	*Page
}

type ErrorResponse struct {
	*http.Response
	Errors []ErrorMessage
}

type ErrorMessage struct {
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		e.Response.Request.Method, e.Response.Request.URL,
		e.Response.StatusCode, e.Errors)
}

// NewClient returns new Bitbucket client for accessing Bitbucket APIs
func NewClient(baseURL string, httpClient *http.Client) (*Client, error) {
	baseEndpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{client: httpClient, ApiVersion: defaultApiVersion, UserAgent: defaultUserAgent}
	c.BaseURL = baseEndpoint
	c.common.client = c
	c.AccessTokens = (*AccessTokensService)(&c.common)
	c.Keys = (*KeysService)(&c.common)
	c.Projects = (*ProjectsService)(&c.common)
	return c, nil
}

// NewRequest created a new http request to call the Bitbucket API
func (c *Client) NewRequest(method, apiName, path string, body interface{}) (*http.Request, error) {
	p := fmt.Sprintf("%s/%s/%s", apiName, c.ApiVersion, path)
	u, err := c.BaseURL.Parse(p)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	r, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	resp := &Response{Response: r}
	defer resp.Body.Close()

	err = CheckResponse(r)
	if err != nil {
		return resp, err
	}

	if v == nil {
		return resp, nil
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		if err != io.EOF {
			return resp, err
		}
	}

	if p, ok := v.(Paged); ok {
		resp.Page = p.Current()
	}

	return resp, nil
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	return &ErrorResponse{Response: r}
}

func (c *Client) Get(ctx context.Context, api, path string, v interface{}) (*Response, error) {
	return c.GetPaged(ctx, api, path, v, nil)
}

func (c *Client) GetPaged(ctx context.Context, api, path string, v interface{}, opts *ListOptions) (*Response, error) {
	req, err := c.NewRequest("GET", api, path, nil)
	if err != nil {
		return nil, err
	}
	if opts != nil {
		query := req.URL.Query()
		if opts.Limit != 0 {
			query.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
		if opts.Start != 0 {
			query.Set("start", fmt.Sprintf("%d", opts.Start))
		}
		req.URL.RawQuery = query.Encode()
	}
	return c.Do(ctx, req, v)
}

// BasicAuthTransport supports creating a http client passing username/password as basic authentication header.
type BasicAuthTransport struct {
	Username string
	Password string
}

// RoundTrip implements the RoundTripper interface.
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := setCredentialsAsHeaders(req, t.Username, t.Password)
	return http.DefaultTransport.RoundTrip(req2)
}

func setCredentialsAsHeaders(req *http.Request, id, secret string) *http.Request {
	// To set extra headers, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of http.RoundTripper.
	//
	// Since we are going to modify only req.Header here, we only need a deep copy
	// of req.Header.
	convertedRequest := new(http.Request)
	*convertedRequest = *req
	convertedRequest.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		convertedRequest.Header[k] = append([]string(nil), s...)
	}
	convertedRequest.SetBasicAuth(id, secret)
	return convertedRequest
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.
func (t *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}
