package bitbucket

import (
	"context"
	"fmt"
)

type WebhookList struct {
	ListResponse
	Webhooks []*Webhook `json:"values"`
}

type Webhook struct {
	ID      uint64                `json:"id,omitempty"`
	Name    string                `json:"name"`
	Created *DateTime             `json:"createdDate,omitempty"`
	Updated *DateTime             `json:"updatedDate,omitempty"`
	Events  []EventKey            `json:"events"`
	Config  *WebhookConfiguration `json:"configuration,omitempty"`
	URL     string                `json:"url"`
	Active  bool                  `json:"active"`
}

type WebhookConfiguration struct {
	Secret string `json:"secret,omitempty"`
}

func (s *ProjectsService) ListWebhooks(ctx context.Context, projectKey, repositorySlug string, opts *ListOptions) ([]*Webhook, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/webhooks", projectKey, repositorySlug)
	var l WebhookList
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Webhooks, resp, nil
}

func (s *ProjectsService) GetWebhook(ctx context.Context, projectKey, repositorySlug string, id uint64) (*Webhook, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/webhooks/%d", projectKey, repositorySlug, id)
	var w Webhook
	resp, err := s.client.Get(ctx, projectsApiName, p, &w)
	if err != nil {
		return nil, resp, err
	}
	return &w, resp, nil
}

func (s *ProjectsService) CreateWebhook(ctx context.Context, projectKey, repositorySlug string, webhook *Webhook) (*Webhook, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/webhooks", projectKey, repositorySlug)
	req, err := s.client.NewRequest("POST", projectsApiName, p, webhook)
	if err != nil {
		return nil, nil, err
	}

	var w Webhook
	resp, err := s.client.Do(ctx, req, &w)
	if err != nil {
		return nil, resp, err
	}
	return &w, resp, nil
}

func (s *ProjectsService) DeleteWebhook(ctx context.Context, projectKey, repositorySlug string, id uint64) (*Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/webhooks/%d", projectKey, repositorySlug, id)
	req, err := s.client.NewRequest("DELETE", projectsApiName, p, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
