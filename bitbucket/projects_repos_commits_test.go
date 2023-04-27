package bitbucket

import (
	"context"
	"io"
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

func TestCreateBuildStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method)
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, "{\"key\":\"BUILD-ID\",\"state\":\"INPROGRESS\",\"url\":\"https://ci.domain.com/builds/BUILD-ID\",\"buildNumber\":\"number\",\"duration\":10000,\"ref\":\"refs/head\"}\n", string(b))
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo/commits/commit/builds", req.URL.Path)
		rw.Write([]byte(getWebhookResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	in := &BuildStatus{
		Key:         "BUILD-ID",
		State:       BuildStatusStateInProgress,
		URL:         "https://ci.domain.com/builds/BUILD-ID",
		BuildNumber: "number",
		Duration:    10000,
		Ref:         "refs/head",
	}
	_, err := client.Projects.CreateBuildStatus(ctx, "PRJ", "repo", "commit", in)
	assert.NoError(t, err)
}

func TestListChanges(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/KEY/repos/repo/commits/00a2f8656ec89118e65588ff3ac18328db35f6bc/changes", req.URL.Path)

		rw.Write([]byte(listChangesResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	commits, resp, err := client.Projects.ListChanges(ctx, "KEY", "repo", "00a2f8656ec89118e65588ff3ac18328db35f6bc", &ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, commits, 1)
	assert.True(t, resp.LastPage)
	assert.Equal(t, "clusters/netic-internal/prod1/releases/k8s-inventory-collector/secrets.yaml", commits[0].Path.Title)
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

const listChangesResponse = `{
	"fromHash": null,
	"toHash": "00a2f8656ec89118e65588ff3ac18328db35f6bc",
	"properties": {},
	"values": [
	  {
		"contentId": "1c602c0799d6d852a07affd0ac06b4ce8548d6f0",
		"fromContentId": "e00a3bef11aec0b71858ebbd07d3158c3ef8e142",
		"path": {
		  "components": [
			"clusters",
			"netic-internal",
			"prod1",
			"releases",
			"k8s-inventory-collector",
			"secrets.yaml"
		  ],
		  "parent": "clusters/netic-internal/prod1/releases/k8s-inventory-collector",
		  "name": "secrets.yaml",
		  "extension": "yaml",
		  "toString": "clusters/netic-internal/prod1/releases/k8s-inventory-collector/secrets.yaml"
		},
		"executable": false,
		"percentUnchanged": -1,
		"type": "MODIFY",
		"nodeType": "FILE",
		"srcExecutable": false,
		"links": {
		  "self": [
			{
			  "href": "https://git.domain.com/projects/KUB/repos/kubernetes-config/commits/00a2f8656ec89118e65588ff3ac18328db35f6bc#clusters/netic-internal/prod1/releases/k8s-inventory-collector/secrets.yaml"
			}
		  ]
		},
		"properties": {
		  "gitChangeType": "MODIFY"
		}
	  }
	],
	"size": 1,
	"isLastPage": true,
	"start": 0,
	"limit": 25,
	"nextPageStart": null
  }`
