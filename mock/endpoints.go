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
	ListRepositories = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos", Method: "GET"}
	GetRepository    = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug", Method: "GET"}
	CreateRepository = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos", Method: "POST"}
	DeleteRepository = EndpointPattern{Pattern: "/api/latest/projects/:projectKey/repos/:repositorySlug", Method: "DELETE"}
)
