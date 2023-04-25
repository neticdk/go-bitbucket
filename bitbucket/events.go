package bitbucket

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

type RepositoryPushEvent struct {
	EventKey   EventKey                    `json:"eventKey"`
	Date       ISOTime                     `json:"date"`
	Actor      User                        `json:"actor"`
	Repository Repository                  `json:"repository"`
	Changes    []RepositoryPushEventChange `json:"changes"`
	Commits    []Commit                    `json:"commits"`
	ToCommit   Commit                      `json:"toCommit"`
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
	EventKey    EventKey    `json:"eventKey"`
	Date        ISOTime     `json:"date"`
	Actor       User        `json:"actor"`
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
