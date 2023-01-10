package bitbucket

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUserTokens(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/access-tokens/latest/users/user-id", req.URL.Path)
		rw.Write([]byte(listTokensUsersResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	tokens, _, err := client.AccessTokens.ListUserTokens(ctx, "user-id", &ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, tokens, 5)
}

func TestGetUserToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/access-tokens/latest/users/user-id/313791499779", req.URL.Path)
		rw.Write([]byte(getTokenUserResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	token, _, err := client.AccessTokens.GetUserToken(ctx, "user-id", "313791499779")
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, token.Name, "write-apikey")
	assert.ElementsMatch(t, []Permission{PermissionProjectWrite, PermissionRepoWrite}, token.Permissions)
}

func TestCreateUserToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "PUT", req.Method)
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, "{\"name\":\"Demo 3\",\"permissions\":[\"PROJECT_READ\",\"REPO_READ\"],\"expiryDays\":128}\n", string(b))
		assert.Equal(t, "/access-tokens/latest/users/user-id", req.URL.Path)
		rw.Write([]byte(createTokenUserResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	in := &AccessToken{
		Name:        "Demo 3",
		ExpireDays:  128,
		Permissions: []Permission{PermissionProjectRead, PermissionRepoRead},
	}
	token, _, err := client.AccessTokens.CreateUserToken(ctx, "user-id", in)
	assert.NoError(t, err)
	assert.Equal(t, "BBDC-..", token.Token)
}

func TestDeleteUserToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "DELETE", req.Method)
		assert.Equal(t, "/access-tokens/latest/users/user-id/token-id", req.URL.Path)
		rw.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	_, err := client.AccessTokens.DeleteUserToken(ctx, "user-id", "token-id")
	assert.NoError(t, err)
}

const listTokensUsersResponse = `{
	"size": 5,
	"limit": 25,
	"isLastPage": true,
	"values": [
	  {
		"id": "2512744587",
		"createdDate": 1654079617056,
		"lastAuthenticated": 1673222715653,
		"name": "Drone 01",
		"permissions": [
		  "PROJECT_WRITE",
		  "REPO_WRITE"
		],
		"user": {
		  "name": "api-user",
		  "emailAddress": "api-user@e.mail",
		  "active": true,
		  "displayName": "Api User",
		  "id": 12910,
		  "slug": "api-user",
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git/users/api-user"
			  }
			]
		  }
		}
	  },
	  {
		"id": "946698762753",
		"createdDate": 1649398269249,
		"lastAuthenticated": 1673271946464,
		"name": "kubernetes-01",
		"permissions": [
		  "PROJECT_READ",
		  "REPO_READ"
		],
		"user": {
		  "name": "api-user",
		  "emailAddress": "api-user@e.mail",
		  "active": true,
		  "displayName": "Api User",
		  "id": 12910,
		  "slug": "api-user",
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git/users/api-user"
			  }
			]
		  }
		}
	  },
	  {
		"id": "913771124628",
		"createdDate": 1618915825168,
		"lastAuthenticated": 1673271960116,
		"name": "shared",
		"permissions": [
		  "PROJECT_READ",
		  "REPO_READ"
		],
		"user": {
		  "name": "api-user",
		  "emailAddress": "api-user@e.mail",
		  "active": true,
		  "displayName": "Api User",
		  "id": 12910,
		  "slug": "api-user",
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git/users/api-user"
			  }
			]
		  }
		}
	  },
	  {
		"id": "783783127394",
		"createdDate": 1614338175360,
		"lastAuthenticated": 1673271611941,
		"name": "drone",
		"permissions": [
		  "PROJECT_READ",
		  "REPO_READ"
		],
		"user": {
		  "name": "api-user",
		  "emailAddress": "api-user@e.mail",
		  "active": true,
		  "displayName": "Api User",
		  "id": 12910,
		  "slug": "api-user",
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git/users/api-user"
			  }
			]
		  }
		}
	  },
	  {
		"id": "313791499779",
		"createdDate": 1612337130284,
		"lastAuthenticated": 1672809528662,
		"name": "write-apikey",
		"permissions": [
		  "PROJECT_WRITE",
		  "REPO_WRITE"
		],
		"user": {
		  "name": "api-user",
		  "emailAddress": "api-user@e.mail",
		  "active": true,
		  "displayName": "Api User",
		  "id": 12910,
		  "slug": "api-user",
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git/users/api-user"
			  }
			]
		  }
		}
	  }
	],
	"start": 0
  }`

const getTokenUserResponse = `{
	"id": "313791499779",
	"createdDate": 1612337130284,
	"lastAuthenticated": 1672809528662,
	"name": "write-apikey",
	"permissions": [
	  "PROJECT_WRITE",
	  "REPO_WRITE"
	],
	"user": {
	  "name": "api-user",
	  "emailAddress": "api-user@e.mail",
	  "active": true,
	  "displayName": "Api User",
	  "id": 12910,
	  "slug": "api-user",
	  "type": "NORMAL",
	  "links": {
		"self": [
		  {
			"href": "https://git/users/api-user"
		  }
		]
	  }
	}
  }`

const createTokenUserResponse = `{
	"id": "786587856882",
	"createdDate": 1673273224446,
	"expiryDate": 1684328824446,
	"name": "Demo 3",
	"permissions": [
	  "PROJECT_READ",
	  "REPO_READ"
	],
	"user": {
	  "name": "api-user",
	  "emailAddress": "api-user@e.mail",
	  "active": true,
	  "displayName": "Api User",
	  "id": 12910,
	  "slug": "api-user",
	  "type": "NORMAL",
	  "links": {
		"self": [
		  {
			"href": "https://git/users/api-user"
		  }
		]
	  }
	},
	"token": "BBDC-.."
  }`
