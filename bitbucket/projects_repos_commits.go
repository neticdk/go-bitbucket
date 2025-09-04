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

type BuildStatus struct {
	Key         string                 `json:"key"`
	State       BuildStatusState       `json:"state"`
	URL         string                 `json:"url"`
	BuildNumber string                 `json:"buildNumber,omitempty"`
	DateAdded   DateTime               `json:"dateAdded,omitempty"`
	Description string                 `json:"description,omitempty"`
	Duration    uint64                 `json:"duration,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Parent      string                 `json:"parent,omitempty"`
	Ref         string                 `json:"ref,omitempty"`
	TestResult  *BuildStatusTestResult `json:"testResults,omitempty"`
}

type BuildStatusTestResult struct {
	Failed     uint32 `json:"failed"`
	Skipped    uint32 `json:"skipped"`
	Successful uint32 `json:"successful"`
}

type BuildStatusState string

const (
	BuildStatusStateCancelled  BuildStatusState = "CANCELLED"
	BuildStatusStateFailed     BuildStatusState = "FAILED"
	BuildStatusStateInProgress BuildStatusState = "INPROGRESS"
	BuildStatusStateSuccessful BuildStatusState = "SUCCESSFUL"
	BuildStatusStateUnknown    BuildStatusState = "UNKNOWN"
)

type Change struct {
	ContentId  string            `json:"contentId"`
	Path       ChangePath        `json:"path"`
	Executable bool              `json:"executable"`
	Unchanged  int               `json:"percentUnchanged"`
	Type       ChangeType        `json:"type"`
	NodeType   ChangeNodeType    `json:"nodeType"`
	Properties map[string]string `json:"properties"`
}

type ChangePath struct {
	Components []string `json:"components"`
	Parent     string   `json:"parent"`
	Name       string   `json:"name"`
	Extension  string   `json:"extension"`
	Title      string   `json:"toString"`
}

type ChangeType string

type ChangeNodeType string

type ChangeList struct {
	ListResponse

	Changes []*Change `json:"values"`
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

func (s *ProjectsService) CreateBuildStatus(ctx context.Context, projectKey, repositorySlug, commitId string, status *BuildStatus) (*Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/commits/%s/builds", projectKey, repositorySlug, commitId)
	req, err := s.client.NewRequest("POST", projectsApiName, p, status)
	if err != nil {
		return nil, err
	}

	var r Repository
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *ProjectsService) ListChanges(ctx context.Context, projectKey, repositorySlug, commitId string, opts *ListOptions) ([]*Change, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/commits/%s/changes", projectKey, repositorySlug, commitId)
	var l ChangeList
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Changes, resp, nil
}
