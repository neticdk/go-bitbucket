package bitbucket

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
