package mock

var (
	ListAccessTokensRepository  = EndpointPattern{Pattern: "/access-tokens/latest/projects/:projectKey/repos/:repositorySlug", Method: "GET"}
	GetAccessTokenRepository    = EndpointPattern{Pattern: "/access-tokens/latest/projects/:projectKey/repos/:repositorySlug/:tokenId", Method: "GET"}
	CreateAccessTokenRepository = EndpointPattern{Pattern: "/access-tokens/latest/projects/:projectKey/repos/:repositorySlug", Method: "PUT"}
	DeleteAccessTokenRepository = EndpointPattern{Pattern: "/access-tokens/latest/projects/:projectKey/repos/:repositorySlug/:tokenId", Method: "DELETE"}
)

var (
	ListAccessTokensUser  = EndpointPattern{Pattern: "/access-tokens/latest/users/:userSlug", Method: "GET"}
	GetAccessTokenUser    = EndpointPattern{Pattern: "/access-tokens/latest/users/:userSlug/:tokenId", Method: "GET"}
	CreateAccessTokenUser = EndpointPattern{Pattern: "/access-tokens/latest/users/:userSlug", Method: "PUT"}
	DeleteAccessTokenUser = EndpointPattern{Pattern: "/access-tokens/latest/users/:userSlug/:tokenId", Method: "DELETE"}
)

var (
	ListKeysRepository  = EndpointPattern{Pattern: "/keys/latest/projects/:projectKey/repos/:repositorySlug/ssh", Method: "GET"}
	GetKeyRepository    = EndpointPattern{Pattern: "/keys/latest/projects/:projectKey/repos/:repositorySlug/ssh/:keyId", Method: "GET"}
	CreateKeyRepository = EndpointPattern{Pattern: "/keys/latest/projects/:projectKey/repos/:repositorySlug/ssh", Method: "POST"}
	DeleteKeyRepository = EndpointPattern{Pattern: "/keys/latest/projects/:projectKey/repos/:repositorySlug/ssh/:keyId", Method: "DELETE"}
)

var (
	ListProjects             = EndpointPattern{Pattern: "/api/latest/projects", Method: "GET"}
	SearchProjectPermissions = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/permissions/search", Method: "GET"}
	SearchRepositories       = EndpointPattern{Pattern: "/api/latest/repos", Method: "GET"}
	ListRepositories         = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos", Method: "GET"}
	GetRepository            = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug", Method: "GET"}
	CreateRepository         = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos", Method: "POST"}
	DeleteRepository         = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug", Method: "DELETE"}
	SearchBranches           = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/branches", Method: "GET"}
	GetDefaultBranch         = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/branches/default", Method: "GET"}
	SearchCommits            = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/commits", Method: "GET"}
	GetCommit                = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/commits/:commitId", Method: "GET"}
	SearchPullRequests       = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/pull-requests", Method: "GET"}
	GetPullRequest           = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/pull-requests/:pullRequestId", Method: "GET"}
	ListWebhooks             = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/webhooks", Method: "GET"}
	GetWebhook               = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/webhooks/:id", Method: "GET"}
	CreateWebhook            = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/webhooks", Method: "POST"}
	DeleteWebhook            = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug/webhooks/:id", Method: "DELETE"}
)

var (
	GetUser = EndpointPattern{Pattern: "/api/latest/users/:userSlug", Method: "GET"}
)
