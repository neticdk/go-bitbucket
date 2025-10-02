package bitbucket

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchPullRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/KEY/repos/repo/pull-requests", req.URL.Path)

		rw.Write([]byte(searchPullRequestsResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	prs, resp, err := client.Projects.SearchPullRequests(ctx, "KEY", "repo", &PullRequestSearchOptions{})
	assert.NoError(t, err)
	assert.Len(t, prs, 1)
	assert.True(t, resp.LastPage)
	assert.Equal(t, uint64(376), prs[0].ID)
	assert.Equal(t, PullRequestStateOpen, prs[0].State)
	assert.Equal(t, "5fd97804dda64ee31b4541340f9ef16043232518", prs[0].Target.Latest)
}

func TestGetPullRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo/pull-requests/376", req.URL.Path)
		rw.Write([]byte(getPullRequestResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	pr, _, err := client.Projects.GetPullRequest(ctx, "PRJ", "repo", 376)
	assert.NoError(t, err)
	assert.NotNil(t, pr)
	assert.Equal(t, uint64(376), pr.ID)
	assert.Equal(t, PullRequestStateOpen, pr.State)
	assert.Equal(t, "5fd97804dda64ee31b4541340f9ef16043232518", pr.Target.Latest)
}

func TestListPullRequestChanges(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo/pull-requests/376/changes", req.URL.Path)
		rw.Write([]byte(listPullRequestChangesResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	changes, resp, err := client.Projects.ListPullRequestChanges(ctx, "PRJ", "repo", 376, nil)

	assert.NoError(t, err)
	assert.NotNil(t, changes)
	assert.NotNil(t, resp)
	assert.Len(t, changes, 5)

	testCases := []struct {
		index      int
		path       string
		changeType ChangeType
		nodeType   ChangeNodeType
	}{
		{0, "src/file1.txt", ChangeType("MODIFY"), ChangeNodeType("FILE")},
		{1, "src/file2.txt", ChangeType("MODIFY"), ChangeNodeType("FILE")},
		{2, "src/utils/file3.txt", ChangeType("MODIFY"), ChangeNodeType("FILE")},
		{3, "src/models/file4.txt", ChangeType("MODIFY"), ChangeNodeType("FILE")},
		{4, "config/file5.txt", ChangeType("MODIFY"), ChangeNodeType("FILE")},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("change[%d]-%s", tc.index, tc.path), func(t *testing.T) {
			change := changes[tc.index]
			assert.Equal(t, tc.path, change.Path.Title)
			assert.Equal(t, tc.changeType, change.Type)
			assert.Equal(t, tc.nodeType, change.NodeType)
		})
	}
}

const searchPullRequestsResponse = `{
	"size": 1,
	"limit": 25,
	"isLastPage": true,
	"values": [
	  {
		"id": 376,
		"version": 0,
		"title": "Promote gitOps from: internal-1  to: innovators-1",
		"description": "Merge request to promote  internal-1  to: innovators-1",
		"state": "OPEN",
		"open": true,
		"closed": false,
		"createdDate": 1682575589130,
		"updatedDate": 1682575589130,
		"fromRef": {
		  "id": "refs/heads/internal-1",
		  "displayId": "internal-1",
		  "latestCommit": "975a45ad48262e2ff55a73f8531430c627d3e2cf",
		  "type": "BRANCH",
		  "repository": {
			"slug": "kubernetes-infrastructure-config",
			"id": 1919,
			"name": "kubernetes-infrastructure-config",
			"hierarchyId": "44ba623c3c521e9be2a1",
			"scmId": "git",
			"state": "AVAILABLE",
			"statusMessage": "Available",
			"forkable": true,
			"project": {
			  "key": "KUB",
			  "id": 1465,
			  "name": "KUBERNETES",
			  "public": false,
			  "type": "NORMAL",
			  "links": {
				"self": [
				  {
					"href": "https://git.domain.com/projects/KUB"
				  }
				]
			  }
			},
			"public": false,
			"archived": false,
			"links": {
			  "clone": [
				{
				  "href": "ssh://git@git.domain.com:7999/kub/kubernetes-infrastructure-config.git",
				  "name": "ssh"
				},
				{
				  "href": "https://git.domain.com/scm/kub/kubernetes-infrastructure-config.git",
				  "name": "http"
				}
			  ],
			  "self": [
				{
				  "href": "https://git.domain.com/projects/KUB/repos/kubernetes-infrastructure-config/browse"
				}
			  ]
			}
		  }
		},
		"toRef": {
		  "id": "refs/heads/innovators-1",
		  "displayId": "innovators-1",
		  "latestCommit": "5fd97804dda64ee31b4541340f9ef16043232518",
		  "type": "BRANCH",
		  "repository": {
			"slug": "kubernetes-infrastructure-config",
			"id": 1919,
			"name": "kubernetes-infrastructure-config",
			"hierarchyId": "44ba623c3c521e9be2a1",
			"scmId": "git",
			"state": "AVAILABLE",
			"statusMessage": "Available",
			"forkable": true,
			"project": {
			  "key": "KUB",
			  "id": 1465,
			  "name": "KUBERNETES",
			  "public": false,
			  "type": "NORMAL",
			  "links": {
				"self": [
				  {
					"href": "https://git.domain.com/projects/KUB"
				  }
				]
			  }
			},
			"public": false,
			"archived": false,
			"links": {
			  "clone": [
				{
				  "href": "ssh://git@git.domain.com:7999/kub/kubernetes-infrastructure-config.git",
				  "name": "ssh"
				},
				{
				  "href": "https://git.domain.com/scm/kub/kubernetes-infrastructure-config.git",
				  "name": "http"
				}
			  ],
			  "self": [
				{
				  "href": "https://git.domain.com/projects/KUB/repos/kubernetes-infrastructure-config/browse"
				}
			  ]
			}
		  }
		},
		"locked": false,
		"author": {
		  "user": {
			"name": "john.doe",
			"emailAddress": "john.doe@domain.com",
			"active": true,
			"displayName": "John Doe",
			"id": 12910,
			"slug": "john.doe",
			"type": "NORMAL",
			"links": {
			  "self": [
				{
				  "href": "https://git.domain.com/users/john.doe"
				}
			  ]
			}
		  },
		  "role": "AUTHOR",
		  "approved": false,
		  "status": "UNAPPROVED"
		},
		"reviewers": [],
		"participants": [],
		"properties": {
		  "resolvedTaskCount": 0,
		  "commentCount": 0,
		  "openTaskCount": 1
		},
		"links": {
		  "self": [
			{
			  "href": "https://git.domain.com/projects/KUB/repos/kubernetes-infrastructure-config/pull-requests/376"
			}
		  ]
		}
	  }
	],
	"start": 0
  }`

const getPullRequestResponse = `{
	"id": 376,
	"version": 0,
	"title": "Promote gitOps from: internal-1  to: innovators-1",
	"description": "Merge request to promote  internal-1  to: innovators-1",
	"state": "OPEN",
	"open": true,
	"closed": false,
	"createdDate": 1682575589130,
	"updatedDate": 1682575589130,
	"fromRef": {
	  "id": "refs/heads/internal-1",
	  "displayId": "internal-1",
	  "latestCommit": "975a45ad48262e2ff55a73f8531430c627d3e2cf",
	  "type": "BRANCH",
	  "repository": {
		"slug": "kubernetes-infrastructure-config",
		"id": 1919,
		"name": "kubernetes-infrastructure-config",
		"hierarchyId": "44ba623c3c521e9be2a1",
		"scmId": "git",
		"state": "AVAILABLE",
		"statusMessage": "Available",
		"forkable": true,
		"project": {
		  "key": "KUB",
		  "id": 1465,
		  "name": "KUBERNETES",
		  "public": false,
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git.domain.com/projects/KUB"
			  }
			]
		  }
		},
		"public": false,
		"archived": false,
		"links": {
		  "clone": [
			{
			  "href": "ssh://git@git.domain.com:7999/kub/kubernetes-infrastructure-config.git",
			  "name": "ssh"
			},
			{
			  "href": "https://git.domain.com/scm/kub/kubernetes-infrastructure-config.git",
			  "name": "http"
			}
		  ],
		  "self": [
			{
			  "href": "https://git.domain.com/projects/KUB/repos/kubernetes-infrastructure-config/browse"
			}
		  ]
		}
	  }
	},
	"toRef": {
	  "id": "refs/heads/innovators-1",
	  "displayId": "innovators-1",
	  "latestCommit": "5fd97804dda64ee31b4541340f9ef16043232518",
	  "type": "BRANCH",
	  "repository": {
		"slug": "kubernetes-infrastructure-config",
		"id": 1919,
		"name": "kubernetes-infrastructure-config",
		"hierarchyId": "44ba623c3c521e9be2a1",
		"scmId": "git",
		"state": "AVAILABLE",
		"statusMessage": "Available",
		"forkable": true,
		"project": {
		  "key": "KUB",
		  "id": 1465,
		  "name": "KUBERNETES",
		  "public": false,
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git.domain.com/projects/KUB"
			  }
			]
		  }
		},
		"public": false,
		"archived": false,
		"links": {
		  "clone": [
			{
			  "href": "ssh://git@git.domain.com:7999/kub/kubernetes-infrastructure-config.git",
			  "name": "ssh"
			},
			{
			  "href": "https://git.domain.com/scm/kub/kubernetes-infrastructure-config.git",
			  "name": "http"
			}
		  ],
		  "self": [
			{
			  "href": "https://git.domain.com/projects/KUB/repos/kubernetes-infrastructure-config/browse"
			}
		  ]
		}
	  }
	},
	"locked": false,
	"author": {
	  "user": {
		"name": "john.doe",
		"emailAddress": "john.doe@domain.com",
		"active": true,
		"displayName": "John Doe",
		"id": 12910,
		"slug": "john.doe",
		"type": "NORMAL",
		"links": {
		  "self": [
			{
			  "href": "https://git.domain.com/users/john.doe"
			}
		  ]
		}
	  },
	  "role": "AUTHOR",
	  "approved": false,
	  "status": "UNAPPROVED"
	},
	"reviewers": [],
	"participants": [],
	"properties": {
	  "resolvedTaskCount": 0,
	  "commentCount": 0,
	  "openTaskCount": 1
	},
	"links": {
	  "self": [
		{
		  "href": "https://git.domain.com/projects/KUB/repos/kubernetes-infrastructure-config/pull-requests/376"
		}
	  ]
	}
  }`

const listPullRequestChangesResponse = `
{
  "fromHash": "0ed160fdfaed81516ccb7c82052123d8ceda16a9",
  "toHash": "c6ab4787d2eb37aa23fdc26515cd22be9669d650",
  "properties": {
    "changeScope": "ALL"
  },
  "values": [
    {
      "contentId": "a7d89a3ff131473187e67f141a801f119b7938e7",
      "fromContentId": "e9f4f2be8fb748cbc0774ac829c7618c3d1f6d90",
      "path": {
        "components": [
          "src",
          "file1.txt"
        ],
        "parent": "src",
        "name": "file1.txt",
        "extension": "txt",
        "toString": "src/file1.txt"
      },
      "executable": false,
      "percentUnchanged": -1,
      "type": "MODIFY",
      "nodeType": "FILE",
      "srcExecutable": false,
      "links": {
        "self": [null]
      },
      "properties": {
        "gitChangeType": "MODIFY"
      }
    },
    {
      "contentId": "2020b30a29976b5ef2c4dfc236bab602ed2239b7",
      "fromContentId": "4da9dc7d9d2c9b3ef9c0767f4327c06beb25d1c1",
      "path": {
        "components": [
          "src",
          "file2.txt"
        ],
        "parent": "src",
        "name": "file2.txt",
        "extension": "txt",
        "toString": "src/file2.txt"
      },
      "executable": false,
      "percentUnchanged": -1,
      "type": "MODIFY",
      "nodeType": "FILE",
      "srcExecutable": false,
      "links": {
        "self": [null]
      },
      "properties": {
	  	"orphanedComments": 0,
        "gitChangeType": "MODIFY",
		"activeComments": 1
      }
    },
    {
      "contentId": "459ba25dfe9b3390b6f8e44a8e77a97a8d09b827",
      "fromContentId": "86fe0b7db922b165caa3124ebc2ffdff96e7c470",
      "path": {
        "components": [
          "src",
          "utils",
          "file3.txt"
        ],
        "parent": "src/utils",
        "name": "file3.txt",
        "extension": "txt",
        "toString": "src/utils/file3.txt"
      },
      "executable": false,
      "percentUnchanged": -1,
      "type": "MODIFY",
      "nodeType": "FILE",
      "srcExecutable": false,
      "links": {
        "self": [null]
      },
      "properties": {
        "gitChangeType": "MODIFY"
      }
    },
    {
      "contentId": "03a08aac85126cc6c74fb4b92ad19159fd933c35",
      "fromContentId": "db088f1d195717cf44ec08f34e2dac7c2d90bbfe",
      "path": {
        "components": [
          "src",
          "models",
          "file4.txt"
        ],
        "parent": "src/models",
        "name": "file4.txt",
        "extension": "txt",
        "toString": "src/models/file4.txt"
      },
      "executable": false,
      "percentUnchanged": -1,
      "type": "MODIFY",
      "nodeType": "FILE",
      "srcExecutable": false,
      "links": {
        "self": [null]
      },
      "properties": {
        "gitChangeType": "MODIFY"
      }
    },
    {
      "contentId": "2d296ff0c4e8bf30c8c1164bef5ef62f7ae4630b",
      "fromContentId": "628d96bba563507aa0ebae791886686fed43dcd8",
      "path": {
        "components": [
          "config",
          "file5.txt"
        ],
        "parent": "config",
        "name": "file5.txt",
        "extension": "txt",
        "toString": "config/file5.txt"
      },
      "executable": false,
      "percentUnchanged": -1,
      "type": "MODIFY",
      "nodeType": "FILE",
      "srcExecutable": false,
      "links": {
        "self": [null]
      },
      "properties": {
        "gitChangeType": "MODIFY"
      }
    }
  ],
  "size": 5,
  "isLastPage": true,
  "start": 0,
  "limit": 25,
  "nextPageStart": null
}
`
