package bitbucket

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProjects(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects", req.URL.Path)
		rw.Write([]byte(listProjectsResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	repos, resp, err := client.Projects.ListProjects(ctx, &ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, repos, 2)
	assert.Equal(t, uint64(363), repos[0].ID)
	assert.True(t, resp.LastPage)
	assert.Equal(t, uint(0), resp.Page.NextPageStart)
}

func TestSearchProjectPermissions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/BB/permissions/search", req.URL.Path)
		rw.Write([]byte(SearchProjectPermissionsResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	perms, resp, err := client.Projects.SearchProjectPermissions(ctx, "BB", &ProjectPermissionSearchOptions{})
	assert.NoError(t, err)
	assert.Len(t, perms, 9)
	assert.True(t, resp.LastPage)
	assert.Equal(t, "bitbucket-administrators", perms[0].Group)
}

const listProjectsResponse = `{
	"size": 2,
	"limit": 25,
	"isLastPage": true,
  	"start": 0,
  	"nextPageStart": 0,
  	"values": [
    {
      "key": "PRJ1",
      "id": 363,
      "name": "Project 1",
      "description": "My project 1",
      "public": false,
      "type": "NORMAL",
      "links": {
        "self": [
          {
            "href": "https://git/projects/PRJ1"
          }
        ]
      }
    },
    {
      "key": "PRJ2",
      "id": 909,
      "name": "Project 2",
      "description": "My Project 2",
      "public": false,
      "type": "NORMAL",
      "links": {
        "self": [
          {
            "href": "https://git/projects/PRJ2"
          }
        ]
      }
    }
    ]
}
`

const SearchProjectPermissionsResponse = `{
  "size": 9,
  "limit": 25,
  "isLastPage": true,
  "values": [
    {
      "permission": "ADMIN",
      "group": "bitbucket-administrators"
    },
    {
      "permission": "SYS_ADMIN",
      "group": "bitbucket-system-administrators"
    },
    {
      "permission": "REPO_CREATE",
      "group": "project-bb-createrepo"
    },
    {
      "permission": "PROJECT_READ",
      "group": "project-bb-read"
    },
    {
      "permission": "PROJECT_WRITE",
      "group": "project-bb-write"
    },
    {
      "permission": "SYS_ADMIN",
      "user": {
        "name": "superadmin",
        "emailAddress": "stash@mymail.dk",
        "active": true,
        "displayName": "Admin",
        "id": 1,
        "slug": "superadmin",
        "type": "NORMAL",
        "links": {
          "self": [
            {
              "href": "https://git.netic.dk/users/superadmin"
            }
          ]
        }
      }
    },
    {
      "permission": "REPO_CREATE",
      "group": "netic-stash-users"
    },
    {
      "permission": "PROJECT_ADMIN",
      "user": {
        "name": "myself@mymail.dk",
        "emailAddress": "myself@mymail.dk",
        "active": true,
        "displayName": "Super Admin",
        "id": 11895,
        "slug": "tal_netic.dk",
        "type": "NORMAL",
        "links": {
          "self": [
            {
              "href": "https://git.netic.dk/users/tal_netic.dk"
            }
          ]
        }
      }
    },
    {
      "permission": "PROJECT_READ",
      "group": "trifork-operations-stash-users-readonly"
    }
  ],
  "start": 0
}`
