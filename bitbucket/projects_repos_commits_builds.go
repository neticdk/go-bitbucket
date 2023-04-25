package bitbucket

import (
	"context"
	"fmt"
)

type BuildStatus struct {
	Key         string                 `json:"key"`
	State       BuildStatusState       `json:"state"`
	URL         string                 `json:"url"`
	BuildNumber string                 `json:"buildNumber,omitempty"`
	Description string                 `json:"description,omitempty"`
	Duration    uint64                 `json:"duration,omitempty"`
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
