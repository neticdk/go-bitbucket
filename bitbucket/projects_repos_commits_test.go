package bitbucket

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
		assert.Equal(t, "{\"key\":\"BUILD-ID\",\"state\":\"INPROGRESS\",\"url\":\"https://ci.domain.com/builds/BUILD-ID\",\"buildNumber\":\"number\",\"dateAdded\":1680350400000,\"duration\":10000,\"name\":\"my-build\",\"parent\":\"parentKey\",\"ref\":\"refs/head\"}\n", string(b))
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
		DateAdded:   DateTime(time.Date(2023, 4, 1, 12, 0, 0, 0, time.UTC)),
		Duration:    10000,
		Ref:         "refs/head",
		Name:        "my-build",
		Parent:      "parentKey",
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

func TestCompareChanges(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/KUB/repos/kubernetes-config/compare/changes", req.URL.Path)
		assert.Equal(t, "from=a&start=0&to=b", req.URL.Query().Encode())
		rw.Write([]byte(compareChangesResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	changes, resp, err := client.Projects.CompareChanges(ctx, "KUB", "kubernetes-config", "a", "b", &CompareChangesOptions{})
	assert.NoError(t, err)
	assert.Len(t, changes, 3)
	assert.True(t, resp.LastPage)
	assert.Equal(t, "lib.c", changes[1].Path.Name)

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

const compareChangesResponse = `{
    "fromHash": "2d9dd5ddcecaa4255e43a20517a2392f3271feb5",
    "isLastPage": true,
    "limit": 25,
    "nextPageStart": null,
    "properties": {},
    "size": 3,
    "start": 0,
    "toHash": "d3560245aa15e5f6d0d9315fc1d85a68af82e705",
    "values": [
        {
            "contentId": "18abf8a5cafd28c37954e8324f1d938cb2fea916",
            "executable": false,
            "fromContentId": "3ad61a0da64ebd1d739b1d204456ff3fbb4fc64d",
            "links": {
                "self": [
                    null
                ]
            },
            "nodeType": "FILE",
            "path": {
                "components": [
                    "README.md"
                ],
                "extension": "md",
                "name": "README.md",
                "parent": "",
                "toString": "README.md"
            },
            "percentUnchanged": -1,
            "properties": {
                "gitChangeType": "MODIFY"
            },
            "srcExecutable": false,
            "type": "MODIFY"
        },
        {
            "contentId": "cf8c875b48d99dd9e721b15088da2c59a65d3e28",
            "executable": false,
            "fromContentId": "0000000000000000000000000000000000000000",
            "links": {
                "self": [
                    null
                ]
            },
            "nodeType": "FILE",
            "path": {
                "components": [
                    "lib.c"
                ],
                "extension": "c",
                "name": "lib.c",
                "parent": "",
                "toString": "lib.c"
            },
            "percentUnchanged": -1,
            "properties": {
                "gitChangeType": "ADD"
            },
            "type": "ADD"
        },
        {
            "contentId": "08494a77a4e97a7139ee7a4d24b57bd139e474bc",
            "executable": false,
            "fromContentId": "d3afd5565cf301872bd447b164ef8c58047f1f8f",
            "links": {
                "self": [
                    null
                ]
            },
            "nodeType": "FILE",
            "path": {
                "components": [
                    "main.c"
                ],
                "extension": "c",
                "name": "main.c",
                "parent": "",
                "toString": "main.c"
            },
            "percentUnchanged": -1,
            "properties": {
                "gitChangeType": "MODIFY"
            },
            "srcExecutable": false,
            "type": "MODIFY"
        }
    ]
}`
