package bitbucket

import (
	"context"
	"fmt"
)

type PullRequest struct {
	ID           uint64                   `json:"id,omitempty"`
	Version      uint64                   `json:"version,omitempty"`
	Title        string                   `json:"title"`
	State        PullRequestState         `json:"state"`
	Open         bool                     `json:"open"`
	Closed       bool                     `json:"closed"`
	Created      *DateTime                `json:"createdDate"`
	Updated      *DateTime                `json:"updatedDate"`
	Source       PullRequestRef           `json:"fromRef"`
	Target       PullRequestRef           `json:"toRef"`
	Locked       bool                     `json:"locked"`
	Author       PullRequestParticipant   `json:"author"`
	Reviewers    []PullRequestParticipant `json:"reviewers"`
	Participants []PullRequestParticipant `json:"participants"`
}

type PullRequestRef struct {
	ID         string     `json:"id"`
	DisplayID  string     `json:"displayId"`
	Latest     string     `json:"latestCommit"`
	Repository Repository `json:"repository"`
}

type PullRequestParticipant struct {
	Author   User                    `json:"user"`
	Role     PullRequestAuthorRole   `json:"role"`
	Approved bool                    `json:"approved"`
	Status   PullRequestAuthorStatus `json:"status"`
	Commit   string                  `json:"lastReviewedCommit,omitempty"`
}

type PullRequestState string

const (
	PullRequestStateDeclined PullRequestState = "DECLINED"
	PullRequestStateMerged   PullRequestState = "MERGED"
	PullRequestStateOpen     PullRequestState = "OPEN"
)

type PullRequestAuthorRole string

const (
	PullRequestAuthorRoleAuthor      PullRequestAuthorRole = "AUTHOR"
	PullRequestAuthorRoleReviewer    PullRequestAuthorRole = "REVIEWER"
	PullRequestAuthorRoleParticipant PullRequestAuthorRole = "PARTICIPANT"
)

type PullRequestAuthorStatus string

const (
	PullRequestAuthorStatusApproved   PullRequestAuthorStatus = "APPROVED"
	PullRequestAuthorStatusUnapproved PullRequestAuthorStatus = "UNAPPROVED"
	PullRequestAuthorStatusNeedsWork  PullRequestAuthorStatus = "NEEDS_WORK"
)

type PullRequestList struct {
	ListResponse

	PullRequests []*PullRequest `json:"values"`
}

type PullRequestSearchOptions struct {
	ListOptions

	At     string            `url:"at,omitempty"`
	Filter string            `url:"filterText,omitempty"`
	State  *PullRequestState `url:"state,omitempty"`
}

func (s *ProjectsService) SearchPullRequests(ctx context.Context, projectKey, repositorySlug string, opts *PullRequestSearchOptions) ([]*PullRequest, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/pull-requests", projectKey, repositorySlug)
	var l PullRequestList
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.PullRequests, resp, nil
}

func (s *ProjectsService) GetPullRequest(ctx context.Context, projectKey, repositorySlug string, pullRequestId uint64) (*PullRequest, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/pull-requests/%d", projectKey, repositorySlug, pullRequestId)
	var pr PullRequest
	resp, err := s.client.Get(ctx, projectsApiName, p, &pr)
	if err != nil {
		return nil, resp, err
	}
	return &pr, resp, nil
}
