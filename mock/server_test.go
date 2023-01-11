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
		WithRequestMatch(GetRepository, bitbucket.Repository{
			ID:    123,
			Name:  "repo",
			Slug:  "repo",
			ScmID: "git",
			Links: map[string][]bitbucket.Link{
				"clone": {
					{
						Href: "https://git/scm/prj/repo.git",
						Name: "http",
					},
					{
						Href: "ssh://git@git:7999/prj/repo.git",
						Name: "ssh",
					},
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

	_, _, err = c.AccessTokens.ListRepositoryTokens(ctx, "prj", "repo", &bitbucket.ListOptions{})
	assert.NoError(t, err)
	_, _, err = c.AccessTokens.ListRepositoryTokens(ctx, "prj", "repo", &bitbucket.ListOptions{})
	assert.NoError(t, err)

	repo, _, err := c.Projects.GetRepository(ctx, "prj", "repo")
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	assert.Len(t, repo.Links["clone"], 2)
}
