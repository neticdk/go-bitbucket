package bitbucket

import (
	"context"
	"fmt"
)

type UsersService service

const usersApiName = "api"

type User struct {
	ID          uint64   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Active      bool     `json:"active"`
	DisplayName string   `json:"displayName"`
	Email       string   `json:"emailAddress,omitempty"`
	Type        UserType `json:"type,omitempty"`
}

type UserType string

const (
	UserTypeNormal  UserType = "NORMAL"
	UserTypeService UserType = "SERVICE"
)

func (s *UsersService) GetUser(ctx context.Context, userSlug string) (*User, *Response, error) {
	p := fmt.Sprintf("users/%s", userSlug)
	var u User
	resp, err := s.client.Get(ctx, usersApiName, p, &u)
	if err != nil {
		return nil, resp, err
	}
	return &u, resp, nil
}
