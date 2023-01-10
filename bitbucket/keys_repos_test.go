package bitbucket

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRepositoryKeys(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/keys/latest/projects/PRJ/repos/repo/ssh", req.URL.Path)
		rw.Write([]byte(listKeysRepoResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	keys, _, err := client.Keys.ListRepositoryKeys(ctx, "PRJ", "repo", &ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, keys, 1)
	assert.Equal(t, uint64(601), keys[0].ID)
	assert.Equal(t, "RSA", keys[0].Algorithm)
	assert.Equal(t, PermissionRepoRead, keys[0].Permission)
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADA.. label01", keys[0].Text)
}

func TestGetRepositoryKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/keys/latest/projects/PRJ/repos/repo/ssh/601", req.URL.Path)
		rw.Write([]byte(getKeyResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	key, _, err := client.Keys.GetRepositoryKey(ctx, "PRJ", "repo", 601)
	assert.NoError(t, err)
	assert.NotNil(t, key)
	assert.Equal(t, uint64(601), key.ID)
	assert.Equal(t, "RSA", key.Algorithm)
	assert.Equal(t, PermissionRepoRead, key.Permission)
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADA.. label01", key.Text)
}

func TestCreateRepositoryKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method)
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, "{\"key\":{\"text\":\"ssh-rsa AAAAB3NzaC1yc2EAAAADAQAB.. label\",\"label\":\"\",\"algorithmType\":\"RSA\",\"bitLength\":4096},\"permission\":\"REPO_READ\"}\n", string(b))
		assert.Equal(t, "/keys/latest/projects/PRJ/repos/repo/ssh", req.URL.Path)
		rw.Write([]byte(createKeyResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	in := &SshKey{
		Text:       "ssh-rsa AAAAB3NzaC1yc2EAAAADAQAB.. label",
		Algorithm:  "RSA",
		Length:     4096,
		Permission: PermissionRepoRead,
	}
	key, _, err := client.Keys.CreateRepositoryKey(ctx, "PRJ", "repo", in)
	assert.NoError(t, err)
	assert.NotNil(t, key)
	assert.Equal(t, "mylabel", key.Label)
}

func TestDeleteRepositoryKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "DELETE", req.Method)
		assert.Equal(t, "/keys/latest/projects/PRJ/repos/repo/ssh/602", req.URL.Path)
		rw.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	_, err := client.Keys.DeleteRepositoryKey(ctx, "PRJ", "repo", 602)
	assert.NoError(t, err)
}

const listKeysRepoResponse = `{
	"size": 1,
	"limit": 25,
	"isLastPage": true,
	"values": [
	  {
		"key": {
		  "id": 601,
		  "text": "ssh-rsa AAAAB3NzaC1yc2EAAAADA.. label01",
		  "label": "label01",
		  "algorithmType": "RSA",
		  "bitLength": 4096
		},
		"repository": {
		  "slug": "gotk-bootstrap-k8s",
		  "id": 1405,
		  "name": "gotk-bootstrap-k8s",
		  "description": "Repository deploying basic components",
		  "hierarchyId": "782aff6f3df32ebed",
		  "scmId": "git",
		  "state": "AVAILABLE",
		  "statusMessage": "Available",
		  "forkable": true,
		  "project": {
			"key": "PD",
			"id": 1084,
			"name": "Platform Development",
			"description": "https://jira/display/PD",
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
		},
		"permission": "REPO_READ"
	  }
	],
	"start": 0
  }`

const getKeyResponse = `{
	"key": {
	  "id": 601,
	  "text": "ssh-rsa AAAAB3NzaC1yc2EAAAADA.. label01",
	  "label": "label01",
	  "algorithmType": "RSA",
	  "bitLength": 4096
	},
	"repository": {
	  "slug": "gotk-bootstrap-k8s",
	  "id": 1405,
	  "name": "gotk-bootstrap-k8s",
	  "description": "Repository deploying basic component on Kubernetes cluster based on flux2/gotk.",
	  "hierarchyId": "782aff6acef3df32ebed",
	  "scmId": "git",
	  "state": "AVAILABLE",
	  "statusMessage": "Available",
	  "forkable": true,
	  "project": {
		"key": "PD",
		"id": 1084,
		"name": "Platform Development",
		"description": "https://jira/display/PD",
		"public": false,
		"type": "NORMAL",
		"links": {
		  "self": [
			{
			  "href": "https://jira/projects/PD"
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
	},
	"permission": "REPO_READ"
  }`
const createKeyResponse = `{
	"key": {
	  "id": 603,
	  "text": "ssh-rsa AAAAB3NzaC1yc2EAAAADA.. mylabel",
	  "label": "mylabel",
	  "algorithmType": "RSA",
	  "bitLength": 4096
	},
	"repository": {
	  "slug": "gotk-bootstrap-k8s",
	  "id": 1405,
	  "name": "gotk-bootstrap-k8s",
	  "description": "Repository deploying basic component on Kubernetes cluster based on flux2/gotk.",
	  "hierarchyId": "782aff6acef3df32ebed",
	  "scmId": "git",
	  "state": "AVAILABLE",
	  "statusMessage": "Available",
	  "forkable": true,
	  "project": {
		"key": "PD",
		"id": 1084,
		"name": "Platform Development",
		"description": "https://jira/display/PD",
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
	},
	"permission": "REPO_READ"
  }`
