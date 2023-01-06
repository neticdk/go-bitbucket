package bitbucket

import "net/http"

type Client struct {
	client *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		client: httpClient,
	}
}
