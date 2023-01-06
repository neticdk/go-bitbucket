package bitbucket

import (
	"context"
	"fmt"
)

type AccessTokenList struct {
	ListResponse
	Tokens []AccessToken `json:"values"`
}

type AccessToken struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Permissions []string  `json:"permissions"`
	Created     *DateTime `json:"createdDate,omitempty"`
	Expire      *DateTime `json:"expiryDate,omitempty"`
	ExpireDays  int       `json:"expiryDays,omitempty"`
	Token       string    `json:"token,omitempty"`
}

const (
	TokenPermissionRepoRead  = "REPO_READ"
	TokenPermissionRepoWrite = "REPO_WRITE"
)

func (s *AccessTokensService) ListRepositoryTokens(ctx context.Context, projectKey, repositorySlug string) ([]AccessToken, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s", projectKey, repositorySlug)
	req, err := s.client.NewRequest("GET", accessTokenApiName, p, nil)
	if err != nil {
		return nil, nil, err
	}
	var list AccessTokenList
	resp, err := s.client.Do(ctx, req, &list)
	if err != nil {
		return nil, resp, err
	}
	return list.Tokens, resp, nil
}

func (s *AccessTokensService) GetRepositoryToken(ctx context.Context, projectKey, repositorySlug, tokenId string) (*AccessToken, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/%s", projectKey, repositorySlug, tokenId)
	req, err := s.client.NewRequest("GET", accessTokenApiName, p, nil)
	if err != nil {
		return nil, nil, err
	}
	var token AccessToken
	resp, err := s.client.Do(ctx, req, &token)
	if err != nil {
		return nil, resp, err
	}
	return &token, resp, nil
}

func (s *AccessTokensService) CreateRepositoryToken(ctx context.Context, projectKey, repositorySlug string, token *AccessToken) (*AccessToken, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s", projectKey, repositorySlug)
	req, err := s.client.NewRequest("PUT", accessTokenApiName, p, token)
	if err != nil {
		return nil, nil, err
	}
	var t AccessToken
	resp, err := s.client.Do(ctx, req, &t)
	if err != nil {
		return nil, resp, err
	}
	return &t, resp, nil
}

func (s *AccessTokensService) DeleteRepositoryToken(ctx context.Context, projectKey, repositorySlug, tokenId string) (*Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/%s", projectKey, repositorySlug, tokenId)
	req, err := s.client.NewRequest("DELETE", accessTokenApiName, p, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
