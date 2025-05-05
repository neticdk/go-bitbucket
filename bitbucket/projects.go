package bitbucket

import (
	"context"
	"fmt"
)

type ProjectsService service

const projectsApiName = "api"

type Project struct {
	ID    uint64            `json:"id,omitempty"`
	Key   string            `json:"key,omitempty"`
	Name  string            `json:"name"`
	Links map[string][]Link `json:"links,omitempty"`
}

type ProjectList struct {
	ListResponse
	Projects []*Project `json:"values"`
}

type ProjectPermissionSearchOptions struct {
	ListOptions

	Permission Permission            `url:"permission,omitempty"`
	Type       ProjectPermissionType `url:"type,omitempty"`
	Filter     string                `url:"filterText,omitempty"`
}

type ProjectPermission struct {
	Permission Permission `json:"permission"`
	User       User       `json:"user"`
	Group      string     `json:"group"`
}

type ProjectPermissionType string

const (
	ProjectPermissionTypeUser  = "USER"
	ProjectPermissionTypeGroup = "GROUP"
)

type ProjectPermissionList struct {
	ListResponse
	Permissions []*ProjectPermission `json:"values"`
}

func (s *ProjectsService) ListProjects(ctx context.Context, opts *ListOptions) ([]*Project, *Response, error) {
	var l ProjectList
	resp, err := s.client.GetPaged(ctx, projectsApiName, "projects", &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Projects, resp, nil
}

func (s *ProjectsService) SearchProjectPermissions(ctx context.Context, projectKey string, opts *ProjectPermissionSearchOptions) ([]*ProjectPermission, *Response, error) {
	p := fmt.Sprintf("projects/%s/permissions/search", projectKey)
	var l ProjectPermissionList
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Permissions, resp, nil
}
