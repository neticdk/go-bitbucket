package bitbucket

type Event struct {
	EventKey EventKey `json:"eventKey"`
	Date     ISOTime  `json:"date"`
	Actor    User     `json:"actor"`
}

type RepositoryPushEvent struct {
	Event

	Repository Repository                  `json:"repository"`
	Changes    []RepositoryPushEventChange `json:"changes"`
	Commits    []Commit                    `json:"commits,omitempty"`
	ToCommit   *Commit                     `json:"toCommit,omitempty"`
}

type RepositoryPushEventChange struct {
	Ref      RepositoryPushEventRef        `json:"ref"`
	RefId    string                        `json:"refId"`
	FromHash string                        `json:"fromHash"`
	ToHash   string                        `json:"toHash"`
	Type     RepositoryPushEventChangeType `json:"type"`
}

type RepositoryPushEventChangeType string

const (
	RepositoryPushEventChangeTypeAdd    RepositoryPushEventChangeType = "ADD"
	RepositoryPushEventChangeTypeUpdate RepositoryPushEventChangeType = "UPDATE"
	RepositoryPushEventChangeTypeDelete RepositoryPushEventChangeType = "DELETE"
)

type RepositoryPushEventRef struct {
	ID        string                     `json:"id"`
	DisplayID string                     `json:"displayId"`
	Type      RepositoryPushEventRefType `json:"type"`
}

type RepositoryPushEventRefType string

const (
	RepositoryPushEventRefTypeBranch RepositoryPushEventRefType = "BRANCH"
	RepositoryPushEventRefTypeTag    RepositoryPushEventRefType = "TAG"
)

type PullRequestEvent struct {
	Event

	PullRequest PullRequest `json:"pullRequest"`
}

type EventKey string

const (
	EventKeyRepoRefsChanged           EventKey = "repo:refs_changed"      // Repo push
	EventKeyRepoModified              EventKey = "repo:modified"          // Repo changed (name)
	EventKeyRepoFork                  EventKey = "repo:fork"              // Repo forked
	EventKeyCommentAdded              EventKey = "repo:comment:added"     // Repo comment on commit added
	EventKeyCommentEdited             EventKey = "repo:comment:edited"    // Repo comment on commit edited
	EventKeyCommentDeleted            EventKey = "repo:comment:deleted"   // Repo comment on commit deleted
	EventKeyPullRequestOpened         EventKey = "pr:opened"              // Pull request opened
	EventKeyPullRequestFrom           EventKey = "pr:from_ref_updated"    // Pull request source ref updated
	EventKeyPullRequestTo             EventKey = "pr:to_ref_updated"      // Pull request target ref updated
	EventkeyPullRequestModified       EventKey = "pr:modified"            // Pull request modified (title, description, target)
	EventKeyPullRequestReviewer       EventKey = "pr:reviewer:updated"    // Pull request reviewers updated
	EventKeyPullRequestApproved       EventKey = "pr:reviewer:approved"   // Pull request approved by reviewer
	EventKeyPullRequestUnapproved     EventKey = "pr:reviewer:unapproved" // Pull request approval withdrawn by reviewer
	EventKeyPullRequestNeedsWork      EventKey = "pr:reviewer:needs_work" // Pull request reviewer marked "needs work"
	EventKeyPullRequestMerged         EventKey = "pr:merged"
	EventKeyPullRequestDeclined       EventKey = "pr:declined"
	EventKeyPullRequestDeleted        EventKey = "pr:deleted"
	EventKeyPullRequestCommentAdded   EventKey = "pr:comment:added"
	EventKeyPullRequestCommentEdited  EventKey = "pr:comment:edited"
	EventKeyPullRequestCommentDeleted EventKey = "pr:comment:deleted"
)
