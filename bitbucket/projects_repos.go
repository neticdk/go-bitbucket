package bitbucket

import (
	"context"
	"fmt"
)

type repositoryList struct {
	ListResponse
	Repositories []*Repository `json:"values"`
}

type Repository struct {
	ID          uint64            `json:"id,omitempty"`
	Slug        string            `json:"slug,omitempty"`
	ScmID       string            `json:"scmId"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Public      bool              `json:"public,omitempty"`
	Archived    bool              `json:"archived,omitempty"`
	State       RepositoryState   `json:"state,omitempty"`
	Links       map[string][]Link `json:"links,omitempty"`
}

type Link struct {
	Href string `json:"href"`
	Name string `json:"name,omitempty"`
}

type RepositoryState string

const (
	RepositoryStateAvailable  RepositoryState = "AVAILABLE"
	RepositoryStateInitFailed RepositoryState = "INITIALISATION_FAILED"
	RepositoryStateInit       RepositoryState = "INITIALISING"
	RepositoryStateOffline    RepositoryState = "OFFLINE"
)

func (s *ProjectsService) ListRepositories(ctx context.Context, projectKey string, opts *ListOptions) ([]*Repository, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos", projectKey)
	var l repositoryList
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Repositories, resp, nil
}

func (s *ProjectsService) GetRepository(ctx context.Context, projectKey, repositorySlug string) (*Repository, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s", projectKey, repositorySlug)
	var r Repository
	resp, err := s.client.Get(ctx, projectsApiName, p, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}

func (s *ProjectsService) CreateRepository(ctx context.Context, projectKey string, repo *Repository) (*Repository, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos", projectKey)
	req, err := s.client.NewRequest("POST", projectsApiName, p, repo)
	if err != nil {
		return nil, nil, err
	}

	var r Repository
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}

func (s *ProjectsService) DeleteRepository(ctx context.Context, projectKey, repositorySlug string) (*Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s", projectKey, repositorySlug)
	req, err := s.client.NewRequest("DELETE", projectsApiName, p, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
