package bitbucket

type AccessTokensService service

const accessTokenApiName = "access-tokens"

type AccessTokenList struct {
	ListResponse
	Tokens []*AccessToken `json:"values"`
}

type AccessToken struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
	Created     *DateTime    `json:"createdDate,omitempty"`
	Expire      *DateTime    `json:"expiryDate,omitempty"`
	ExpireDays  int          `json:"expiryDays,omitempty"`
	Token       string       `json:"token,omitempty"`
	User        *User        `json:"user,omitempty"`
}

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
