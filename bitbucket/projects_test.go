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
