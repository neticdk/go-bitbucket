package bitbucket

type AccessTokensService service

const accessTokenApiName = "access-tokens"

type accessTokenList struct {
	ListResponse
	Tokens []AccessToken `json:"values"`
}

type AccessToken struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
	Created     *DateTime    `json:"createdDate,omitempty"`
	Expire      *DateTime    `json:"expiryDate,omitempty"`
	ExpireDays  int          `json:"expiryDays,omitempty"`
	Token       string       `json:"token,omitempty"`
}
