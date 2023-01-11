package mock

import (
	"context"
	"testing"

	"github.com/neticdk/go-bitbucket/bitbucket"
	"github.com/stretchr/testify/assert"
)

func TestNewMockServer(t *testing.T) {
	mockServer := NewMockServer(
		WithRequestMatch(ListAccessTokensRepository, bitbucket.AccessTokenList{
			Tokens: []*bitbucket.AccessToken{
				{
					ID:   "id",
					Name: "name",
				},
			},
		}),
	)

	ctx := context.Background()
	c, _ := bitbucket.NewClient(mockServer.URL, nil)
	tokens, _, err := c.AccessTokens.ListRepositoryTokens(ctx, "prj", "repo", &bitbucket.ListOptions{})
	assert.NoError(t, err)
	assert.Len(t, tokens, 1)
	assert.Equal(t, "id", tokens[0].ID)
	assert.Equal(t, "name", tokens[0].Name)
}
