package bitbucket

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProjectRepositories(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos", req.URL.Path)
		rw.Write([]byte(listProjectsRepositoriesResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	repos, resp, err := client.Projects.ListRepositories(ctx, "PRJ", &ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, repos, 3)
	assert.False(t, resp.LastPage)
	assert.Equal(t, 25, resp.Page.NextPageStart)
}

func TestListProjectRepositoriesNextPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos", req.URL.Path)
		assert.Equal(t, "25", req.URL.Query().Get("start"))
		rw.Write([]byte(listProjectsRepositoriesResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	repos, _, err := client.Projects.ListRepositories(ctx, "PRJ", &ListOptions{Start: 25})
	assert.NoError(t, err)
	assert.Len(t, repos, 3)
}

func TestGetRepository(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo", req.URL.Path)
		rw.Write([]byte(getProjectsRepositoryResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	repo, _, err := client.Projects.GetRepository(ctx, "PRJ", "repo")
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, uint64(1405), repo.ID)
	assert.Equal(t, "repo", repo.Slug)
	assert.Equal(t, "repo", repo.Name)
}

func TestCreateRepository(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method)
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, "{\"scmId\":\"git\",\"name\":\"go-bitbucket-demo\"}\n", string(b))
		assert.Equal(t, "/api/latest/projects/PRJ/repos", req.URL.Path)
		rw.Write([]byte(createProjectRepositoryResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	in := &Repository{
		Name:  "go-bitbucket-demo",
		ScmID: "git",
	}
	repo, _, err := client.Projects.CreateRepository(ctx, "PRJ", in)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, "go-bitbucket-demo", repo.Slug)
	assert.ElementsMatch(t, []Link{{Href: "https://git/scm/pd/go-bitbucket-demo.git", Name: "http"}, {Href: "ssh://git@git:7999/pd/go-bitbucket-demo.git", Name: "ssh"}}, repo.Links["clone"])
}

func TestDeleteRepository(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "DELETE", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo", req.URL.Path)
		rw.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	_, err := client.Projects.DeleteRepository(ctx, "PRJ", "repo")
	assert.NoError(t, err)
}

const listProjectsRepositoriesResponse = `{
	"size": 25,
	"limit": 25,
	"isLastPage": false,
	"values": [
	  {
		"slug": "ansible-netic-kubernetes",
		"id": 1373,
		"name": "ansible-netic-kubernetes",
		"description": "Ansible Galaxy Collection for Netic Kubernetes provisioning",
		"hierarchyId": "5b74a4aa2e26460a2a99",
		"scmId": "git",
		"state": "AVAILABLE",
		"statusMessage": "Available",
		"forkable": true,
		"project": {
		  "key": "PD",
		  "id": 1084,
		  "name": "Netic Platform Development",
		  "description": "https://docs.netic.dk/display/PD",
		  "public": false,
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git.netic.dk/projects/PD"
			  }
			]
		  }
		},
		"public": false,
		"archived": false,
		"links": {
		  "clone": [
			{
			  "href": "ssh://git@git.netic.dk:7999/pd/ansible-netic-kubernetes.git",
			  "name": "ssh"
			},
			{
			  "href": "https://git.netic.dk/scm/pd/ansible-netic-kubernetes.git",
			  "name": "http"
			}
		  ],
		  "self": [
			{
			  "href": "https://git.netic.dk/projects/PD/repos/ansible-netic-kubernetes/browse"
			}
		  ]
		}
	  },
	  {
		"slug": "azure-rke",
		"id": 1473,
		"name": "azure-rke",
		"description": "Showcase deploying RKE2 cluster on Azure.",
		"hierarchyId": "3bd1375a3eaad48aeaf2",
		"scmId": "git",
		"state": "AVAILABLE",
		"statusMessage": "Available",
		"forkable": true,
		"project": {
		  "key": "PD",
		  "id": 1084,
		  "name": "Netic Platform Development",
		  "description": "https://docs.netic.dk/display/PD",
		  "public": false,
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git.netic.dk/projects/PD"
			  }
			]
		  }
		},
		"public": false,
		"archived": false,
		"links": {
		  "clone": [
			{
			  "href": "https://git.netic.dk/scm/pd/azure-rke.git",
			  "name": "http"
			},
			{
			  "href": "ssh://git@git.netic.dk:7999/pd/azure-rke.git",
			  "name": "ssh"
			}
		  ],
		  "self": [
			{
			  "href": "https://git.netic.dk/projects/PD/repos/azure-rke/browse"
			}
		  ]
		}
	  },
	  {
		"slug": "ingest-netic-vector-template",
		"id": 1158,
		"name": "ingest-netic-vector-template",
		"description": "A template for ingest based on vector. This is used for creation of configurations that a customer needs for ingest. the template is used to construct a set of files with customer specific values and they are then used in a customer specific repo.",
		"hierarchyId": "f5f57804346e61413309",
		"scmId": "git",
		"state": "AVAILABLE",
		"statusMessage": "Available",
		"forkable": true,
		"project": {
		  "key": "PD",
		  "id": 1084,
		  "name": "Netic Platform Development",
		  "description": "https://docs.netic.dk/display/PD",
		  "public": false,
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git.netic.dk/projects/PD"
			  }
			]
		  }
		},
		"public": false,
		"archived": false,
		"links": {
		  "clone": [
			{
			  "href": "ssh://git@git.netic.dk:7999/pd/ingest-netic-vector-template.git",
			  "name": "ssh"
			},
			{
			  "href": "https://git.netic.dk/scm/pd/ingest-netic-vector-template.git",
			  "name": "http"
			}
		  ],
		  "self": [
			{
			  "href": "https://git.netic.dk/projects/PD/repos/ingest-netic-vector-template/browse"
			}
		  ]
		}
	  }
	],
	"start": 0,
	"nextPageStart": 25
  }`

const getProjectsRepositoryResponse = `{
	"slug": "repo",
	"id": 1405,
	"name": "repo",
	"description": "Repository deploying basic component on Kubernetes cluster based on flux2/gotk.",
	"hierarchyId": "782aff6acef3df32ebed",
	"scmId": "git",
	"state": "AVAILABLE",
	"statusMessage": "Available",
	"forkable": true,
	"project": {
	  "key": "PD",
	  "id": 1084,
	  "name": "Netic Platform Development",
	  "description": "https://docs.netic.dk/display/PD",
	  "public": false,
	  "type": "NORMAL",
	  "links": {
		"self": [
		  {
			"href": "https://git.netic.dk/projects/PD"
		  }
		]
	  }
	},
	"public": false,
	"archived": false,
	"links": {
	  "clone": [
		{
		  "href": "https://git.netic.dk/scm/pd/gotk-bootstrap-k8s.git",
		  "name": "http"
		},
		{
		  "href": "ssh://git@git.netic.dk:7999/pd/gotk-bootstrap-k8s.git",
		  "name": "ssh"
		}
	  ],
	  "self": [
		{
		  "href": "https://git.netic.dk/projects/PD/repos/gotk-bootstrap-k8s/browse"
		}
	  ]
	}
  }`

const createProjectRepositoryResponse = `{
	"slug": "go-bitbucket-demo",
	"id": 2123,
	"name": "go-bitbucket-demo",
	"hierarchyId": "ff9d1bb49a803771364e",
	"scmId": "git",
	"state": "AVAILABLE",
	"statusMessage": "Available",
	"forkable": true,
	"project": {
	  "key": "PD",
	  "id": 1084,
	  "name": "Netic Platform Development",
	  "description": "https://docs.netic.dk/display/PD",
	  "public": false,
	  "type": "NORMAL",
	  "links": {
		"self": [
		  {
			"href": "https://git.netic.dk/projects/PD"
		  }
		]
	  }
	},
	"public": false,
	"archived": false,
	"links": {
	  "clone": [
		{
		  "href": "https://git/scm/pd/go-bitbucket-demo.git",
		  "name": "http"
		},
		{
		  "href": "ssh://git@git:7999/pd/go-bitbucket-demo.git",
		  "name": "ssh"
		}
	  ],
	  "self": [
		{
		  "href": "https://git.netic.dk/projects/PD/repos/go-bitbucket-demo/browse"
		}
	  ]
	}
  }`
