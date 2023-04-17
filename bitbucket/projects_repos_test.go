package bitbucket

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRepositories(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/repos", req.URL.Path)
		assert.Equal(t, "AVAILABLE", req.URL.Query().Get("state"))

		rw.Write([]byte(listProjectsRepositoriesResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	repos, resp, err := client.Projects.SearchRepositories(ctx, &RepositorySearchOptions{State: RepositoryStateAvailable})
	assert.NoError(t, err)
	assert.Len(t, repos, 3)
	assert.False(t, resp.LastPage)
	assert.Equal(t, uint(25), resp.Page.NextPageStart)
}

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
	assert.Equal(t, uint(25), resp.Page.NextPageStart)
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

func TestListFiles(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/REPO/files/path", req.URL.Path)
		assert.Equal(t, "ref/heads", req.URL.Query().Get("at"))
		rw.Write([]byte(listFilesResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	repos, resp, err := client.Projects.ListFiles(ctx, "PRJ", "REPO", "path", &FilesListOptions{At: "ref/heads"})
	assert.NoError(t, err)
	assert.Len(t, repos, 25)
	assert.False(t, resp.LastPage)
	assert.Equal(t, uint(25), resp.Page.NextPageStart)
}

func TestGetTextFileContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/REPO/browse/path", req.URL.Path)
		assert.Equal(t, "ref/heads", req.URL.Query().Get("at"))
		rw.Write([]byte(getFileContentText))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	repos, _, err := client.Projects.GetTextFileContent(ctx, "PRJ", "REPO", "path", "ref/heads")
	assert.NoError(t, err)
	assert.Equal(t, []byte(`ci {
  include    = [ "clusters/internal/prod1/releases/cortex-rules/.*/.*.yaml" ]
  baseBranch = "main"
}`), repos)
}

func TestGetTextFileContentBinary(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/REPO/browse/path", req.URL.Path)
		assert.Equal(t, "ref/heads", req.URL.Query().Get("at"))
		rw.Write([]byte(getFileContentBinary))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	_, _, err := client.Projects.GetTextFileContent(ctx, "PRJ", "REPO", "path", "ref/heads")
	assert.Error(t, err)
}

func TestGetTextFileContentList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/REPO/browse/path", req.URL.Path)
		assert.Equal(t, "ref/heads", req.URL.Query().Get("at"))
		rw.Write([]byte(getFileContentsList))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	_, _, err := client.Projects.GetTextFileContent(ctx, "PRJ", "REPO", "path", "ref/heads")
	assert.Error(t, err)
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
		  "description": "https://wiki/display/PD",
		  "public": false,
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git/projects/PD"
			  }
			]
		  }
		},
		"public": false,
		"archived": false,
		"links": {
		  "clone": [
			{
			  "href": "ssh://git@git:7999/pd/ansible-netic-kubernetes.git",
			  "name": "ssh"
			},
			{
			  "href": "https://git/scm/pd/ansible-netic-kubernetes.git",
			  "name": "http"
			}
		  ],
		  "self": [
			{
			  "href": "https://git/projects/PD/repos/ansible-netic-kubernetes/browse"
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
		  "description": "https://wiki/display/PD",
		  "public": false,
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git/projects/PD"
			  }
			]
		  }
		},
		"public": false,
		"archived": false,
		"links": {
		  "clone": [
			{
			  "href": "https://git/scm/pd/azure-rke.git",
			  "name": "http"
			},
			{
			  "href": "ssh://git@git:7999/pd/azure-rke.git",
			  "name": "ssh"
			}
		  ],
		  "self": [
			{
			  "href": "https://git/projects/PD/repos/azure-rke/browse"
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
		  "description": "https://wiki/display/PD",
		  "public": false,
		  "type": "NORMAL",
		  "links": {
			"self": [
			  {
				"href": "https://git/projects/PD"
			  }
			]
		  }
		},
		"public": false,
		"archived": false,
		"links": {
		  "clone": [
			{
			  "href": "ssh://git@git:7999/pd/ingest-netic-vector-template.git",
			  "name": "ssh"
			},
			{
			  "href": "https://git/scm/pd/ingest-netic-vector-template.git",
			  "name": "http"
			}
		  ],
		  "self": [
			{
			  "href": "https://git/projects/PD/repos/ingest-netic-vector-template/browse"
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
	  "description": "https://wiki/display/PD",
	  "public": false,
	  "type": "NORMAL",
	  "links": {
		"self": [
		  {
			"href": "https://git/projects/PD"
		  }
		]
	  }
	},
	"public": false,
	"archived": false,
	"links": {
	  "clone": [
		{
		  "href": "https://git/scm/pd/gotk-bootstrap-k8s.git",
		  "name": "http"
		},
		{
		  "href": "ssh://git@git:7999/pd/gotk-bootstrap-k8s.git",
		  "name": "ssh"
		}
	  ],
	  "self": [
		{
		  "href": "https://git/projects/PD/repos/gotk-bootstrap-k8s/browse"
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
	  "description": "https://wiki/display/PD",
	  "public": false,
	  "type": "NORMAL",
	  "links": {
		"self": [
		  {
			"href": "https://git/projects/PD"
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
		  "href": "https://git/projects/PD/repos/go-bitbucket-demo/browse"
		}
	  ]
	}
  }`

const listFilesResponse = `{
	"values": [
	  ".pipe.yml",
	  ".gitignore",
	  ".pint.hcl",
	  ".woodpecker.yaml",
	  "Contribution.md",
	  "README.md",
	  "clusters/netic-ci/prod1/bootstrap/kustomization.yaml",
	  "clusters/netic-ci/prod1/bootstrap/sync.yaml",
	  "clusters/netic-ci/prod1/namespaces/default-config.yaml",
	  "clusters/netic-ci/prod1/namespaces/kustomization.yaml",
	  "clusters/netic-ci/prod1/releases/addons.yaml",
	  "clusters/netic-ci/prod1/releases/external-dns.yaml",
	  "clusters/netic-ci/prod1/releases/kustomization.yaml",
	  "clusters/internal/mgmt1/bootstrap/infra-bootstrap.yaml",
	  "clusters/internal/mgmt1/bootstrap/kustomization.yaml",
	  "clusters/internal/mgmt1/bootstrap/sync.yaml",
	  "clusters/internal/prod1/bootstrap/kustomization.yaml",
	  "clusters/internal/prod1/bootstrap/sync.yaml",
	  "clusters/internal/prod1/namespaces/kustomization.yaml",
	  "clusters/internal/prod1/releases/addons.yaml",
	  "clusters/internal/prod1/releases/component/certificate.yaml",
	  "clusters/internal/prod1/releases/component/configmap.yaml",
	  "clusters/internal/prod1/releases/component/daemonset.yaml",
	  "clusters/internal/prod1/releases/component/deployment.yaml",
	  "clusters/internal/prod1/releases/component/kustomization.yaml"
	],
	"size": 25,
	"isLastPage": false,
	"start": 0,
	"limit": 25,
	"nextPageStart": 25
  }`

const getFileContentsList = `{
	"path": {
	  "components": [
		"other"
	  ],
	  "parent": "",
	  "name": "other",
	  "toString": "other"
	},
	"revision": "refs/heads/master",
	"children": {
	  "size": 11,
	  "limit": 500,
	  "isLastPage": true,
	  "values": [
		{
		  "path": {
			"components": [
			  "architecture.png"
			],
			"parent": "",
			"name": "architecture.png",
			"extension": "png",
			"toString": "architecture.png"
		  },
		  "contentId": "760e34efbbe9bf3573a5576cdca3b5fd9ef6ac5a",
		  "type": "FILE",
		  "size": 94229
		},
		{
		  "path": {
			"components": [
			  "data_distribution.png"
			],
			"parent": "",
			"name": "data_distribution.png",
			"extension": "png",
			"toString": "data_distribution.png"
		  },
		  "contentId": "442c8c25bc3150e72b893af1ddba278d27216077",
		  "type": "FILE",
		  "size": 25666
		},
		{
		  "path": {
			"components": [
			  "screenshot1.png"
			],
			"parent": "",
			"name": "screenshot1.png",
			"extension": "png",
			"toString": "screenshot1.png"
		  },
		  "contentId": "e78458dc47beccd2f74150261f8b8cdee4a0c0c0",
		  "type": "FILE",
		  "size": 396234
		},
		{
		  "path": {
			"components": [
			  "sloop-test.png"
			],
			"parent": "",
			"name": "sloop-test.png",
			"extension": "png",
			"toString": "sloop-test.png"
		  },
		  "contentId": "e0152066e2d6fb5ac1b35633cf215d87569746d4",
		  "type": "FILE",
		  "size": 102921
		},
		{
		  "path": {
			"components": [
			  "sloop_logo_black.eps"
			],
			"parent": "",
			"name": "sloop_logo_black.eps",
			"extension": "eps",
			"toString": "sloop_logo_black.eps"
		  },
		  "contentId": "68f8ac5a3b6ccb56671cb5b84e70a0930f6f8d05",
		  "type": "FILE",
		  "size": 1935722
		},
		{
		  "path": {
			"components": [
			  "sloop_logo_black.png"
			],
			"parent": "",
			"name": "sloop_logo_black.png",
			"extension": "png",
			"toString": "sloop_logo_black.png"
		  },
		  "contentId": "1e815ea4f0a9fbe111f3368814ea0d778677e373",
		  "type": "FILE",
		  "size": 10248
		},
		{
		  "path": {
			"components": [
			  "sloop_logo_color.eps"
			],
			"parent": "",
			"name": "sloop_logo_color.eps",
			"extension": "eps",
			"toString": "sloop_logo_color.eps"
		  },
		  "contentId": "70a8c0cbc4a10722764c69d2cbb8e5e2ee199a60",
		  "type": "FILE",
		  "size": 1946114
		},
		{
		  "path": {
			"components": [
			  "sloop_logo_color.png"
			],
			"parent": "",
			"name": "sloop_logo_color.png",
			"extension": "png",
			"toString": "sloop_logo_color.png"
		  },
		  "contentId": "6d27619c6cd478ccd0c9254722ac0a105fcc73c8",
		  "type": "FILE",
		  "size": 17452
		},
		{
		  "path": {
			"components": [
			  "sloop_logo_color_small_notext.png"
			],
			"parent": "",
			"name": "sloop_logo_color_small_notext.png",
			"extension": "png",
			"toString": "sloop_logo_color_small_notext.png"
		  },
		  "contentId": "0350b1aaac9ccb7e9806a251f1d95f12532770dc",
		  "type": "FILE",
		  "size": 10927
		},
		{
		  "path": {
			"components": [
			  "sloop_logo_white.eps"
			],
			"parent": "",
			"name": "sloop_logo_white.eps",
			"extension": "eps",
			"toString": "sloop_logo_white.eps"
		  },
		  "contentId": "aa62f81e169cf298608b99c73f46bb4ac05f6f00",
		  "type": "FILE",
		  "size": 1915218
		},
		{
		  "path": {
			"components": [
			  "sloop_logo_white.png"
			],
			"parent": "",
			"name": "sloop_logo_white.png",
			"extension": "png",
			"toString": "sloop_logo_white.png"
		  },
		  "contentId": "b474d2116e735e2eee2ba23519b0dc7ed31d9388",
		  "type": "FILE",
		  "size": 10197
		}
	  ],
	  "start": 0
	}
  }`

const getFileContentText = `{
	"lines": [
	  {
		"text": "ci {"
	  },
	  {
		"text": "  include    = [ \"clusters/internal/prod1/releases/cortex-rules/.*/.*.yaml\" ]"
	  },
	  {
		"text": "  baseBranch = \"main\""
	  },
	  {
		"text": "}"
	  }
	],
	"start": 0,
	"size": 4,
	"isLastPage": true,
	"limit": 500,
	"nextPageStart": null
  }`

const getFileContentBinary = `{
	"binary": true,
	"path": {
	  "components": [
		"other",
		"sloop_logo_color_small_notext.png"
	  ],
	  "parent": "other",
	  "name": "sloop_logo_color_small_notext.png",
	  "extension": "png",
	  "toString": "other/sloop_logo_color_small_notext.png"
	}
  }`
