package bitbucket

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	secretKey := []byte("0123456789abcdef")

	buf := bytes.NewBufferString(repoPushEvent01)
	req, err := http.NewRequest("POST", "http://server.io/webhook", buf)
	req.Header.Add(EventSignatureHeader, "sha256=d82c0422a140fc24335536d9450538aeaa978dbc741262a161ee12b99a6bf05d")
	req.Header.Add(EventKeyHeader, "repo:refs_changed")
	assert.NoError(t, err)

	ev, err := ParsePayload(req, secretKey)
	assert.NoError(t, err)
	assert.NotNil(t, ev)

	repoEv, ok := ev.(*RepositoryPushEvent)
	assert.True(t, ok)
	assert.Equal(t, "rep_1", repoEv.Repository.Slug)
}
