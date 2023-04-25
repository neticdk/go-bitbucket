package bitbucket

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/api/latest/users/api-user", req.URL.Path)
		rw.Write([]byte(getUserResponse))
	}))
	defer server.Close()

	client, _ := NewClient(server.URL, nil)
	ctx := context.Background()
	user, _, err := client.Users.GetUser(ctx, "api-user")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, uint64(12910), user.ID)
	assert.Equal(t, "apiuser", user.Slug)
	assert.Equal(t, "api-user", user.Name)
	assert.Equal(t, UserTypeNormal, user.Type)
}

const getUserResponse = `{
	"name": "api-user",
	"emailAddress": "api-user@e.mail",
	"active": true,
	"displayName": "Api User",
	"id": 12910,
	"slug": "apiuser",
	"type": "NORMAL"
}`
