package bitbucket

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCommits(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/KEY/repos/repo/commits", req.URL.Path)

		rw.Write([]byte(searchCommitsResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	commits, resp, err := client.Projects.SearchCommits(ctx, "KEY", "repo", &CommitSearchOptions{})
	assert.NoError(t, err)
	assert.Len(t, commits, 1)
	assert.True(t, resp.LastPage)
	assert.Equal(t, "967e7c24bb92c983221ff176eb202f847dd64e6b", commits[0].ID)
}

func TestGetCommit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo/commits/61228bababcf80662a4ea69fc50eaad440b56ff4", req.URL.Path)
		rw.Write([]byte(getCommitResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	commit, _, err := client.Projects.GetCommit(ctx, "PRJ", "repo", "61228bababcf80662a4ea69fc50eaad440b56ff4")
	assert.NoError(t, err)
	assert.NotNil(t, commit)
	assert.Equal(t, "61228bababcf80662a4ea69fc50eaad440b56ff4", commit.ID)
	assert.Equal(t, "john.doe@mail.com", commit.Author.Name)
	assert.Equal(t, "john.doe@mail.com", commit.Author.Email)
	assert.Equal(t, "Added my awesome feature", commit.Message)

}

const searchCommitsResponse = `{
	"values": [
	  {
		"id": "967e7c24bb92c983221ff176eb202f847dd64e6b",
		"displayId": "967e7c24bb9",
		"author": {
		  "name": "john.doe@domain.com",
		  "emailAddress": "john.doe@domain.com",
		  "active": true,
		  "displayName": "John Doe [Netic]",
		  "id": 57,
		  "slug": "john.doe_domain.com",
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git.domain.com/users/john.doe_domain.com"
			  }
			]
		  }
		},
		"authorTimestamp": 1682577046000,
		"committer": {
		  "name": "john.doe@domain.com",
		  "emailAddress": "john.doe@domain.com",
		  "active": true,
		  "displayName": "John Doe [Netic]",
		  "id": 57,
		  "slug": "john.doe_domain.com",
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git.domain.com/users/john.doe_domain.com"
			  }
			]
		  }
		},
		"committerTimestamp": 1682577046000,
		"message": "Pull request #883: Update default custom",
		"parents": [
		  {
			"id": "31c6e8f2eb4e66689ba6d3a51a87b07987153236",
			"displayId": "31c6e8f2eb4"
		  },
		  {
			"id": "97e0a9cf1cb5086a1c5820cdbc0ccd71ece1bd3c",
			"displayId": "97e0a9cf1cb"
		  }
		]
	  }
	],
	"size": 1,
	"isLastPage": true,
	"start": 0,
	"limit": 25,
	"nextPageStart": 0
  }`

const getCommitResponse = `{
	"id": "61228bababcf80662a4ea69fc50eaad440b56ff4",
	"displayId": "61228bababc",
	"author": {
	  "name": "john.doe@mail.com",
	  "emailAddress": "john.doe@mail.com",
	  "active": true,
	  "displayName": "John Doe",
	  "id": 115,
	  "slug": "john.doe_mail.com",
	  "type": "NORMAL",
	  "links": {
		"self": [
		  {
			"href": "https://git.domain.com/users/john.doe_mail.com"
		  }
		]
	  }
	},
	"authorTimestamp": 1682516769000,
	"committer": {
	  "name": "john.doe@mail.com",
	  "emailAddress": "john.doe@mail.com",
	  "active": true,
	  "displayName": "John Doe",
	  "id": 115,
	  "slug": "john.doe_mail.com",
	  "type": "NORMAL",
	  "links": {
		"self": [
		  {
			"href": "https://git.domain.com/users/john.doe_mail.com"
		  }
		]
	  }
	},
	"committerTimestamp": 1682516769000,
	"message": "Added my awesome feature",
	"parents": [
	  {
		"id": "5ee1bf90b03e88e8fa033a1bacba1fa03627f92d",
		"displayId": "5ee1bf90b03",
		"author": {
		  "name": "John Doe",
		  "emailAddress": "john.doe@mail.com"
		},
		"authorTimestamp": 1682516555000,
		"committer": {
		  "name": "John Doe",
		  "emailAddress": "john.doe@mail.com"
		},
		"committerTimestamp": 1682516555000,
		"message": "Added hcp exporter",
		"parents": [
		  {
			"id": "9da150629fe63cee29ebd589ebaff79138058abe",
			"displayId": "9da150629fe"
		  }
		]
	  }
	]
  }`
