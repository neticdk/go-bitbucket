package bitbucket

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRepositoryTokens(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/access-tokens/latest/projects/PRJ/repos/repo", req.URL.Path)
		rw.Write([]byte(listTokenRepoResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	tokens, _, err := client.AccessTokens.ListRepositoryTokens(ctx, "PRJ", "repo", &ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, tokens, 2)
}

func TestGetRepositoryToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/access-tokens/latest/projects/PRJ/repos/repo/373646823580", req.URL.Path)
		rw.Write([]byte(getTokenRepoResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	token, _, err := client.AccessTokens.GetRepositoryToken(ctx, "PRJ", "repo", "373646823580")
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, "demo", token.Name)
	assert.Len(t, token.Permissions, 1)
}

func TestGetRepositoryTokenNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/access-tokens/latest/projects/PRJ/repos/repo/373646823580", req.URL.Path)
		rw.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	_, _, err := client.AccessTokens.GetRepositoryToken(ctx, "PRJ", "repo", "373646823580")
	assert.Error(t, err)
	er, _ := err.(*ErrorResponse)
	assert.Equal(t, http.StatusNotFound, er.Response.StatusCode)
}

func TestCreateRepositoryToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "PUT", req.Method)
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, "{\"name\":\"Demo 3\",\"permissions\":[\"REPO_READ\"],\"expiryDays\":128}\n", string(b))
		assert.Equal(t, "/access-tokens/latest/projects/PRJ/repos/repo", req.URL.Path)
		rw.Write([]byte(createTokenResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	in := &AccessToken{
		Name:        "Demo 3",
		ExpireDays:  128,
		Permissions: []Permission{PermissionRepoRead},
	}
	token, _, err := client.AccessTokens.CreateRepositoryToken(ctx, "PRJ", "repo", in)
	assert.NoError(t, err)
	assert.Equal(t, "BBDC-xxxx", token.Token)
	assert.Equal(t, "access-token-user/2/1405", token.User.Name)
}

func TestDeleteRepositoryToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "DELETE", req.Method)
		assert.Equal(t, "/access-tokens/latest/projects/PRJ/repos/repo/848894838086", req.URL.Path)
		rw.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	_, err := client.AccessTokens.DeleteRepositoryToken(ctx, "PRJ", "repo", "848894838086")
	assert.NoError(t, err)
}

const listTokenRepoResponse = `{
	"size": 2,
	"limit": 25,
	"isLastPage": true,
	"values": [
	  {
		"id": "848894838086",
		"createdDate": 1672997457717,
		"expiryDate": 1680769857717,
		"name": "demo expire",
		"permissions": [
		  "REPO_READ"
		],
		"user": {
		  "name": "access-token-user/2/1405",
		  "active": true,
		  "displayName": "Access Token User - xxxx",
		  "emailAddress": null,
		  "id": 16772,
		  "slug": "access-token-user_2_1405",
		  "type": "SERVICE",
		  "links": {
			"self": [
			  {
				"href": "https://git/bots/access-token-user_2_1405"
			  }
			]
		  }
		}
	  },
	  {
		"id": "373646823580",
		"createdDate": 1672996351030,
		"name": "demo",
		"permissions": [
		  "REPO_READ"
		],
		"user": {
		  "name": "access-token-user/2/1405",
		  "active": true,
		  "displayName": "Access Token User - xxx",
		  "emailAddress": null,
		  "id": 16772,
		  "slug": "access-token-user_2_1405",
		  "type": "SERVICE",
		  "links": {
			"self": [
			  {
				"href": "https://git/bots/access-token-user_2_1405"
			  }
			]
		  }
		}
	  }
	],
	"start": 0
  }`

const getTokenRepoResponse = `{
	"id": "373646823580",
	"createdDate": 1672996351030,
	"name": "demo",
	"permissions": [
	  "REPO_READ"
	],
	"user": {
	  "name": "access-token-user/2/1405",
	  "active": true,
	  "displayName": "Access Token User - xxx",
	  "emailAddress": null,
	  "id": 16772,
	  "slug": "access-token-user_2_1405",
	  "type": "SERVICE",
	  "links": {
		"self": [
		  {
			"href": "https://git/bots/access-token-user_2_1405"
		  }
		]
	  }
	}
  }`

const createTokenResponse = `{
	"id": "909906670557",
	"createdDate": 1673007051341,
	"expiryDate": 1684062651341,
	"name": "Demo 3",
	"permissions": [
	  "REPO_READ"
	],
	"user": {
	  "name": "access-token-user/2/1405",
	  "active": true,
	  "displayName": "Access Token User - xx",
	  "emailAddress": null,
	  "id": 16772,
	  "slug": "access-token-user_2_1405",
	  "type": "SERVICE",
	  "links": {
		"self": [
		  {
			"href": "https://git/bots/access-token-user_2_1405"
		  }
		]
	  }
	},
	"token": "BBDC-xxxx"
  }`
