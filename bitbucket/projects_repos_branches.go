package bitbucket

import (
	"context"
	"fmt"
)

type BranchList struct {
	ListResponse

	Branches []*Branch `json:"values"`
}

type Branch struct {
	ID              string     `json:"id"`
	DisplayID       string     `json:"displayId"`
	Type            BranchType `json:"type"`
	LatestCommit    string     `json:"latestCommit"`
	LatestChangeset string     `json:"latestChangeset"`
	Default         bool       `json:"isDefault"`
}

type BranchType string

type BranchSearchOptions struct {
	ListOptions

	Filter string            `url:"filterText,omitempty"`
	Order  BranchSearchOrder `url:"orderBy,omitempty"`
}

type BranchSearchOrder string

const (
	BranchSearchOrderAlpha    BranchSearchOrder = "ALPHABETICAL"
	BranchSearchOrderModified BranchSearchOrder = "MODIFICATION"
)

func (s *ProjectsService) SearchBranches(ctx context.Context, projectKey, repositorySlug string, opts *BranchSearchOptions) ([]*Branch, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/branches", projectKey, repositorySlug)
	var l BranchList
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Branches, resp, nil
}

func (s *ProjectsService) GetDefaultBranch(ctx context.Context, projectKey, repositorySlug string) (*Branch, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/branches/default", projectKey, repositorySlug)
	var b Branch
	resp, err := s.client.Get(ctx, projectsApiName, p, &b)
	if err != nil {
		return nil, resp, err
	}
	return &b, resp, nil
}
