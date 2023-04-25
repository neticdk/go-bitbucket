package bitbucket

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListWebhooks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo/webhooks", req.URL.Path)
		rw.Write([]byte(listWebhooksResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	hooks, _, err := client.Projects.ListWebhooks(ctx, "PRJ", "repo", &ListOptions{})
	assert.NoError(t, err)
	assert.NotNil(t, hooks)
	assert.Len(t, hooks, 1)
	assert.Len(t, hooks[0].Events, 7)
}

func TestGetWebhook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo/webhooks/10", req.URL.Path)
		rw.Write([]byte(getWebhookResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	hook, _, err := client.Projects.GetWebhook(ctx, "PRJ", "repo", 10)
	assert.NoError(t, err)
	assert.NotNil(t, hook)
	assert.Len(t, hook.Events, 7)
}

func TestCreateWebhook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method)
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, "{\"name\":\"go-bitbucket-demo\",\"events\":[\"repo:refs_changed\"],\"url\":\"https://domain.com/webhook\",\"active\":true}\n", string(b))
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo/webhooks", req.URL.Path)
		rw.Write([]byte(getWebhookResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	in := &Webhook{
		Name:   "go-bitbucket-demo",
		Events: []EventKey{EventKeyRepoRefsChanged},
		URL:    "https://domain.com/webhook",
		Active: true,
	}
	repo, _, err := client.Projects.CreateWebhook(ctx, "PRJ", "repo", in)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	// Response is same as "get"
}

const listWebhooksResponse = `{
	"size": 1,
	"limit": 25,
	"isLastPage": true,
	"values": [
	  {
		"id": 10,
		"name": "drone",
		"createdDate": 1682407778168,
		"updatedDate": 1682407778168,
		"events": [
		  "pr:merged",
		  "pr:modified",
		  "pr:opened",
		  "repo:refs_changed",
		  "pr:declined",
		  "pr:deleted",
		  "pr:from_ref_updated"
		],
		"configuration": {
		  "secret": "1234567890abcdefghjikl"
		},
		"url": "https://ci.domain.com/hook",
		"active": true
	  }
	],
	"start": 0
  }`

const getWebhookResponse = `{
	"id": 10,
	"name": "drone",
	"createdDate": 1682408715521,
	"updatedDate": 1682408715521,
	"events": [
	  "pr:merged",
	  "pr:modified",
	  "pr:opened",
	  "repo:refs_changed",
	  "pr:declined",
	  "pr:deleted",
	  "pr:from_ref_updated"
	],
	"configuration": {
	  "secret": "1234567890abcdefghjikl"
	},
	"url": "https://ci.domain.com/hook",
	"active": true
  }`
