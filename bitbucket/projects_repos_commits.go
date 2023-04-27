package bitbucket

import (
	"context"
	"fmt"
)

type GitUser struct {
	Name  string `json:"name"`
	Email string `json:"emailAddress"`
}

type CommitData struct {
	ID        string   `json:"id"`
	DisplayID string   `json:"displayId"`
	Message   string   `json:"message"`
	Author    GitUser  `json:"author"`
	Authored  DateTime `json:"authorTimestamp"`
	Comitter  GitUser  `json:"committer"`
	Comitted  DateTime `json:"committerTimestamp"`
}

type Commit struct {
	CommitData
	Parents []CommitData `json:"parents"`
}

type CommitSearchOptions struct {
	ListOptions

	// An optional path to filter commits by
	Path string `url:"path,omitempty"`

	// If true, the commit history of the specified file will be followed past renames. Only valid for a path to a single file.
	FollowRenames bool `url:"followRenames,omitempty"`

	// The commit ID (SHA1) or ref (inclusively) to retrieve commits before
	Until string `url:"until,omitempty"`

	// The commit ID or ref (exclusively) to retrieve commits after
	Since string `url:"since,omitempty"`

	Merges CommitSearchMerges `json:"merges,omitempty"`

	// true to ignore missing commits, false otherwise
	IgnoreMissing bool `json:"ignoreMissing,omitempty"`
}

type CommitSearchMerges string

const (
	CommitSearchMergesExclude CommitSearchMerges = "exclude"
	CommitSearchMergesInclude CommitSearchMerges = "include"
	CommitSearchMergesOnly    CommitSearchMerges = "only"
)

type CommitList struct {
	ListResponse

	Commits []*Commit `json:"values"`
}

func (s *ProjectsService) SearchCommits(ctx context.Context, projectKey, repositorySlug string, opts *CommitSearchOptions) ([]*Commit, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/commits", projectKey, repositorySlug)
	var l CommitList
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Commits, resp, nil
}

func (s *ProjectsService) GetCommit(ctx context.Context, projectKey, repositorySlug, commitId string) (*Commit, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/commits/%s", projectKey, repositorySlug, commitId)
	var c Commit
	resp, err := s.client.Get(ctx, projectsApiName, p, &c)
	if err != nil {
		return nil, resp, err
	}
	return &c, resp, nil
}
