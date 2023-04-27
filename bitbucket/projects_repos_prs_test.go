package bitbucket

import (
	"context"
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
