package bitbucket

import (
	"context"
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

func (s *ProjectsService) ListProjects(ctx context.Context, opts *ListOptions) ([]*Project, *Response, error) {
	var l ProjectList
	resp, err := s.client.GetPaged(ctx, projectsApiName, "projects", &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Projects, resp, nil
}
