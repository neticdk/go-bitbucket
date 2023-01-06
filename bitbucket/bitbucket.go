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
}

type service struct {
	client *Client
}

type ListResponse struct {
	Size          int  `json:"size"`
	Limit         int  `json:"limit"`
	LastPage      bool `json:"isLastPage"`
	Start         int  `json:"start"`
	NextPageStart int  `json:"nextPageStart"`
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

type Response struct {
	*http.Response
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

	// If status code 4xx - parse errors and add to response?

	if v == nil {
		return resp, nil
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		if err != io.EOF {
			return resp, err
		}
	}

	return resp, nil
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	return &ErrorResponse{Response: r}
}
