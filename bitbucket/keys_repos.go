package bitbucket

import (
	"context"
	"fmt"
)

type sshKeyList struct {
	ListResponse
	Values []internalSshKey `json:"values"`
}

type internalSshKey struct {
	Key struct {
		ID        uint64 `json:"id,omitempty"`
		Text      string `json:"text"`
		Label     string `json:"label"`
		Algorithm string `json:"algorithmType"`
		Length    uint   `json:"bitLength"`
	} `json:"key"`
	Permission Permission `json:"permission"`
}

// SshKey defines Bitbucket representation of ssh-key
type SshKey struct {
	ID         uint64
	Text       string
	Label      string
	Algorithm  string
	Length     uint
	Permission Permission
}

func (s *KeysService) ListRepositoryKeys(ctx context.Context, projectKey, repositorySlug string, opts *ListOptions) ([]*SshKey, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/ssh", projectKey, repositorySlug)
	var list sshKeyList
	resp, err := s.client.GetPaged(ctx, keysApiName, p, &list, opts)
	if err != nil {
		return nil, resp, err
	}
	keys := make([]*SshKey, 0)
	for _, k := range list.Values {
		keys = append(keys, newSshKey(k))
	}
	return keys, resp, nil
}

func (s *KeysService) GetRepositoryKey(ctx context.Context, projectKey, repositorySlug string, keyId uint64) (*SshKey, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/ssh/%d", projectKey, repositorySlug, keyId)
	var k internalSshKey
	resp, err := s.client.Get(ctx, keysApiName, p, &k)
	if err != nil {
		return nil, resp, err
	}
	return newSshKey(k), resp, nil
}

func (s *KeysService) CreateRepositoryKey(ctx context.Context, projectKey, repositorySlug string, key *SshKey) (*SshKey, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/ssh", projectKey, repositorySlug)
	k := &internalSshKey{
		Permission: key.Permission,
	}
	k.Key.Text = key.Text
	k.Key.Algorithm = key.Algorithm
	k.Key.Length = key.Length
	req, err := s.client.NewRequest("POST", keysApiName, p, k)
	if err != nil {
		return nil, nil, err
	}

	k = &internalSshKey{}
	resp, err := s.client.Do(ctx, req, k)
	if err != nil {
		return nil, resp, err
	}
	result := newSshKey(*k)
	return result, resp, nil
}

func (s *KeysService) DeleteRepositoryKey(ctx context.Context, projectKey, repositorySlug string, keyId uint64) (*Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/ssh/%d", projectKey, repositorySlug, keyId)
	req, err := s.client.NewRequest("DELETE", keysApiName, p, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func newSshKey(k internalSshKey) *SshKey {
	return &SshKey{
		ID:         k.Key.ID,
		Text:       k.Key.Text,
		Label:      k.Key.Label,
		Algorithm:  k.Key.Algorithm,
		Length:     k.Key.Length,
		Permission: k.Permission,
	}
}
