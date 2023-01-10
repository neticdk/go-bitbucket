package bitbucket

import (
	"context"
	"fmt"
)

func (s *AccessTokensService) ListUserTokens(ctx context.Context, userSlug string, opts *ListOptions) ([]*AccessToken, *Response, error) {
	p := fmt.Sprintf("users/%s", userSlug)
	var list accessTokenList
	resp, err := s.client.GetPaged(ctx, accessTokenApiName, p, &list, opts)
	if err != nil {
		return nil, resp, err
	}
	return list.Tokens, resp, nil
}

func (s *AccessTokensService) GetUserToken(ctx context.Context, userSlug, tokenId string) (*AccessToken, *Response, error) {
	p := fmt.Sprintf("users/%s/%s", userSlug, tokenId)
	var token AccessToken
	resp, err := s.client.Get(ctx, accessTokenApiName, p, &token)
	if err != nil {
		return nil, resp, err
	}
	return &token, resp, nil
}

func (s *AccessTokensService) CreateUserToken(ctx context.Context, userSlug string, token *AccessToken) (*AccessToken, *Response, error) {
	p := fmt.Sprintf("users/%s", userSlug)
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

func (s *AccessTokensService) DeleteUserToken(ctx context.Context, userSlug, tokenId string) (*Response, error) {
	p := fmt.Sprintf("users/%s/%s", userSlug, tokenId)
	req, err := s.client.NewRequest("DELETE", accessTokenApiName, p, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
