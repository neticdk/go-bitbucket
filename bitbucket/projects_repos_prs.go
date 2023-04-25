package bitbucket

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
	PullRequestStateOpen     PullRequestState = "MERGED"
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
