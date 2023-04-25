package bitbucket

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBranches(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/KEY/repos/repo/branches", req.URL.Path)
		assert.Equal(t, "update", req.URL.Query().Get("filterText"))

		rw.Write([]byte(searchBranchesResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	branches, resp, err := client.Projects.SearchBranches(ctx, "KEY", "repo", &BranchSearchOptions{Filter: "update"})
	assert.NoError(t, err)
	assert.Len(t, branches, 25)
	assert.False(t, resp.LastPage)
	assert.Equal(t, uint(25), resp.Page.NextPageStart)
	assert.Equal(t, "refs/heads/upgrade/cert-manager-upgrade-v1.10.2", branches[0].ID)
	assert.Equal(t, "upgrade/cert-manager-upgrade-v1.10.2", branches[0].DisplayID)
}

func TestGetDefaultBranch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/projects/PRJ/repos/repo/branches/default", req.URL.Path)
		rw.Write([]byte(getDefaultBranchResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	branch, _, err := client.Projects.GetDefaultBranch(ctx, "PRJ", "repo")
	assert.NoError(t, err)
	assert.NotNil(t, branch)
	assert.Equal(t, "refs/heads/main", branch.ID)
	assert.Equal(t, "main", branch.DisplayID)
	assert.Equal(t, "b2f97dd9d05e82b1e5308bcffbb5c013065dfeb2", branch.LatestCommit)
	assert.True(t, branch.Default)
}

const searchBranchesResponse = `{
	"size": 25,
	"limit": 25,
	"isLastPage": false,
	"values": [
	  {
		"id": "refs/heads/upgrade/cert-manager-upgrade-v1.10.2",
		"displayId": "upgrade/cert-manager-upgrade-v1.10.2",
		"type": "BRANCH",
		"latestCommit": "ae25676b718886fcc0d14fedc5ef72d1b0762c63",
		"latestChangeset": "ae25676b718886fcc0d14fedc5ef72d1b0762c63",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/cert-manager-upgrade-v1.11.1",
		"displayId": "upgrade/cert-manager-upgrade-v1.11.1",
		"type": "BRANCH",
		"latestCommit": "2f593f67a3115b2bdbdbc2cdbfe15b7d2796b0d4",
		"latestChangeset": "2f593f67a3115b2bdbdbc2cdbfe15b7d2796b0d4",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-10.1.4",
		"displayId": "upgrade/contour-upgrade-10.1.4",
		"type": "BRANCH",
		"latestCommit": "aec64e0ddf63be3a62a1b001a96f87fe2f093a2f",
		"latestChangeset": "aec64e0ddf63be3a62a1b001a96f87fe2f093a2f",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-10.1.5",
		"displayId": "upgrade/contour-upgrade-10.1.5",
		"type": "BRANCH",
		"latestCommit": "6fa13ceae3f1e6b436eb4a94fb5b47949d6813ec",
		"latestChangeset": "6fa13ceae3f1e6b436eb4a94fb5b47949d6813ec",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-10.2.0",
		"displayId": "upgrade/contour-upgrade-10.2.0",
		"type": "BRANCH",
		"latestCommit": "1b14764f352b09b3267221b364726887b6b18354",
		"latestChangeset": "1b14764f352b09b3267221b364726887b6b18354",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-10.2.2",
		"displayId": "upgrade/contour-upgrade-10.2.2",
		"type": "BRANCH",
		"latestCommit": "aa4d4dde2865b27e26e2622594d20a11da05f82a",
		"latestChangeset": "aa4d4dde2865b27e26e2622594d20a11da05f82a",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.0.0",
		"displayId": "upgrade/contour-upgrade-11.0.0",
		"type": "BRANCH",
		"latestCommit": "d3ade4cd65a0c42304c3482e1ef98a6c95c278b1",
		"latestChangeset": "d3ade4cd65a0c42304c3482e1ef98a6c95c278b1",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.0.1",
		"displayId": "upgrade/contour-upgrade-11.0.1",
		"type": "BRANCH",
		"latestCommit": "7b2f2bbf9111b11180dc7ddae53aa97fc8818845",
		"latestChangeset": "7b2f2bbf9111b11180dc7ddae53aa97fc8818845",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.0.2",
		"displayId": "upgrade/contour-upgrade-11.0.2",
		"type": "BRANCH",
		"latestCommit": "cfe52f22f6d37632f875771da9dbbd78da8a4c56",
		"latestChangeset": "cfe52f22f6d37632f875771da9dbbd78da8a4c56",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.0.3",
		"displayId": "upgrade/contour-upgrade-11.0.3",
		"type": "BRANCH",
		"latestCommit": "a39c8dbb723cede1854de6b05f88ac357c9537b7",
		"latestChangeset": "a39c8dbb723cede1854de6b05f88ac357c9537b7",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.0.4",
		"displayId": "upgrade/contour-upgrade-11.0.4",
		"type": "BRANCH",
		"latestCommit": "eafc8266c59513bf578dae3456ae78965202c42d",
		"latestChangeset": "eafc8266c59513bf578dae3456ae78965202c42d",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.1.0",
		"displayId": "upgrade/contour-upgrade-11.1.0",
		"type": "BRANCH",
		"latestCommit": "aad635d787f83a3079c1c516cc9fa3e5d81035d9",
		"latestChangeset": "aad635d787f83a3079c1c516cc9fa3e5d81035d9",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.1.1",
		"displayId": "upgrade/contour-upgrade-11.1.1",
		"type": "BRANCH",
		"latestCommit": "d9ae8522bc71070f71d9ea2f9cf0eb5bb30c1e57",
		"latestChangeset": "d9ae8522bc71070f71d9ea2f9cf0eb5bb30c1e57",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.1.2",
		"displayId": "upgrade/contour-upgrade-11.1.2",
		"type": "BRANCH",
		"latestCommit": "9f42af0cd0dd782b776891c32c6bda58a1489792",
		"latestChangeset": "9f42af0cd0dd782b776891c32c6bda58a1489792",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.1.3",
		"displayId": "upgrade/contour-upgrade-11.1.3",
		"type": "BRANCH",
		"latestCommit": "981dd5bd4dad1ce8709358cd890cd79f271db90c",
		"latestChangeset": "981dd5bd4dad1ce8709358cd890cd79f271db90c",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/contour-upgrade-11.3.0",
		"displayId": "upgrade/contour-upgrade-11.3.0",
		"type": "BRANCH",
		"latestCommit": "ffb23a874d4a9a49efbf4526769b271bea3b7cae",
		"latestChangeset": "ffb23a874d4a9a49efbf4526769b271bea3b7cae",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/external-secrets-upgrade-0.5.7",
		"displayId": "upgrade/external-secrets-upgrade-0.5.7",
		"type": "BRANCH",
		"latestCommit": "dd08d6e495b6bb7c17fe461038c7e6a7fc1ecc83",
		"latestChangeset": "dd08d6e495b6bb7c17fe461038c7e6a7fc1ecc83",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/external-secrets-upgrade-0.5.8",
		"displayId": "upgrade/external-secrets-upgrade-0.5.8",
		"type": "BRANCH",
		"latestCommit": "62cdc55c0999e2f1510b7ab546833df7dc93fba7",
		"latestChangeset": "62cdc55c0999e2f1510b7ab546833df7dc93fba7",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/external-secrets-upgrade-0.5.9",
		"displayId": "upgrade/external-secrets-upgrade-0.5.9",
		"type": "BRANCH",
		"latestCommit": "7a257ecca24d22f6faeb7bfb957906da72619d29",
		"latestChangeset": "7a257ecca24d22f6faeb7bfb957906da72619d29",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/external-secrets-upgrade-0.6.0",
		"displayId": "upgrade/external-secrets-upgrade-0.6.0",
		"type": "BRANCH",
		"latestCommit": "5e4971e93aedd530b8291b5c909e57ccedf06127",
		"latestChangeset": "5e4971e93aedd530b8291b5c909e57ccedf06127",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/external-secrets-upgrade-0.6.1",
		"displayId": "upgrade/external-secrets-upgrade-0.6.1",
		"type": "BRANCH",
		"latestCommit": "c3962f1ef504ba19d7c9e5b84bbe09b504a1638b",
		"latestChangeset": "c3962f1ef504ba19d7c9e5b84bbe09b504a1638b",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/external-secrets-upgrade-0.7.0",
		"displayId": "upgrade/external-secrets-upgrade-0.7.0",
		"type": "BRANCH",
		"latestCommit": "05c2cd9b57959ed25d330f8f5419aa1ce91b4f77",
		"latestChangeset": "05c2cd9b57959ed25d330f8f5419aa1ce91b4f77",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/external-secrets-upgrade-0.7.1",
		"displayId": "upgrade/external-secrets-upgrade-0.7.1",
		"type": "BRANCH",
		"latestCommit": "821fd66273cd80b7abaa9e248baf12d5cee8e7b6",
		"latestChangeset": "821fd66273cd80b7abaa9e248baf12d5cee8e7b6",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/external-secrets-upgrade-0.8.0",
		"displayId": "upgrade/external-secrets-upgrade-0.8.0",
		"type": "BRANCH",
		"latestCommit": "7d295860e6f68213efda935c6647866af97561d2",
		"latestChangeset": "7d295860e6f68213efda935c6647866af97561d2",
		"isDefault": false
	  },
	  {
		"id": "refs/heads/upgrade/external-secrets-upgrade-0.8.1",
		"displayId": "upgrade/external-secrets-upgrade-0.8.1",
		"type": "BRANCH",
		"latestCommit": "536e0a6499b97a4d60dfade62dd178f366438c47",
		"latestChangeset": "536e0a6499b97a4d60dfade62dd178f366438c47",
		"isDefault": false
	  }
	],
	"start": 0,
	"nextPageStart": 25
  }`

const getDefaultBranchResponse = `{
	"id": "refs/heads/main",
	"displayId": "main",
	"type": "BRANCH",
	"latestCommit": "b2f97dd9d05e82b1e5308bcffbb5c013065dfeb2",
	"latestChangeset": "b2f97dd9d05e82b1e5308bcffbb5c013065dfeb2",
	"isDefault": true
  }`
