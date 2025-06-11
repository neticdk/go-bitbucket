package bitbucket

import (
	"bytes"
	"context"
	"fmt"
)

type RepositoryList struct {
	ListResponse
	Repositories []*Repository `json:"values"`
}

type Repository struct {
	ID          uint64            `json:"id,omitempty"`
	Slug        string            `json:"slug,omitempty"`
	ScmID       string            `json:"scmId"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Public      bool              `json:"public,omitempty"`
	Archived    bool              `json:"archived,omitempty"`
	State       *RepositoryState  `json:"state,omitempty"`
	Project     *Project          `json:"project,omitempty"`
	Links       map[string][]Link `json:"links,omitempty"`
}

type Link struct {
	Href string `json:"href"`
	Name string `json:"name,omitempty"`
}

type RepositoryState string

const (
	RepositoryStateAvailable  RepositoryState = "AVAILABLE"
	RepositoryStateInitFailed RepositoryState = "INITIALISATION_FAILED"
	RepositoryStateInit       RepositoryState = "INITIALISING"
	RepositoryStateOffline    RepositoryState = "OFFLINE"
)

type FileList struct {
	ListResponse

	Files []string `json:"values"`
}

type FileContent struct {
	ListResponse

	Binary   bool       `json:"binary,omitempty"`
	Lines    []FileLine `json:"lines,omitempty"`
	Path     *FilePath  `json:"path,omitempty"`
	Revision string     `json:"revision,omitempty"`
}

type FileLine struct {
	Text string `json:"text"`
}

type FilePath struct {
	Components []string `json:"components"`
	Parent     string   `json:"parent"`
	Name       string   `json:"name"`
	Extension  string   `json:"extension"`
}

type RepositorySearchOptions struct {
	ListOptions

	Archived    RepositoryArchived   `url:"archived,omitempty"`
	ProjectName string               `url:"projectname,omitempty"`
	ProjectKey  string               `url:"projectkey,omitempty"`
	Visibility  RepositoryVisibility `url:"visibility,omitempty"`
	Name        string               `url:"name,omitempty"`
	Permission  Permission           `url:"permission,omitempty"`
	State       RepositoryState      `url:"state,omitempty"`
}

type RepositoryArchived string

const (
	RepositoryArchivedActive   RepositoryArchived = "ACTIVE"
	RepositoryArchivedArchived RepositoryArchived = "ARCHIVED"
	RepositoryArchivedAll      RepositoryArchived = "ALL"
)

type RepositoryVisibility string

const (
	RepositoryVisbibilityPrivate RepositoryVisibility = "private"
	RepositoryVisbibilityPublic  RepositoryVisibility = "public"
)

type FilesListOptions struct {
	ListOptions

	At string `url:"at,omitempty"`
}

type FileContentOptions struct {
	ListOptions

	At string `url:"at,omitempty"`
}

func (s *ProjectsService) SearchRepositories(ctx context.Context, opts *RepositorySearchOptions) ([]*Repository, *Response, error) {
	var l RepositoryList
	resp, err := s.client.GetPaged(ctx, projectsApiName, "repos", &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Repositories, resp, nil
}

func (s *ProjectsService) ListRepositories(ctx context.Context, projectKey string, opts *ListOptions) ([]*Repository, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos", projectKey)
	var l RepositoryList
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Repositories, resp, nil
}

func (s *ProjectsService) GetRepository(ctx context.Context, projectKey, repositorySlug string) (*Repository, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s", projectKey, repositorySlug)
	var r Repository
	resp, err := s.client.Get(ctx, projectsApiName, p, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}

func (s *ProjectsService) CreateRepository(ctx context.Context, projectKey string, repo *Repository) (*Repository, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos", projectKey)
	req, err := s.client.NewRequest("POST", projectsApiName, p, repo)
	if err != nil {
		return nil, nil, err
	}

	var r Repository
	resp, err := s.client.Do(ctx, req, &r)
	if err != nil {
		return nil, resp, err
	}
	return &r, resp, nil
}

func (s *ProjectsService) DeleteRepository(ctx context.Context, projectKey, repositorySlug string) (*Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s", projectKey, repositorySlug)
	req, err := s.client.NewRequest("DELETE", projectsApiName, p, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *ProjectsService) ListFiles(ctx context.Context, projectKey, repositorySlug, path string, opts *FilesListOptions) ([]string, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/files/%s", projectKey, repositorySlug, path)

	var l FileList
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.Files, resp, nil
}

type ErrOnlyTextFilesSupported struct{}

func (e *ErrOnlyTextFilesSupported) Error() string {
	return "only files with text content supported"
}

func (s *ProjectsService) GetTextFileContent(ctx context.Context, projectKey, repositorySlug, path, at string) ([]byte, *Response, error) {
	p := fmt.Sprintf("projects/%s/repos/%s/browse/%s", projectKey, repositorySlug, path)

	var f FileContent
	resp, err := s.client.GetPaged(ctx, projectsApiName, p, &f, &FileContentOptions{At: at})
	if err != nil {
		return nil, resp, err
	}

	if f.Binary {
		return nil, nil, &ErrOnlyTextFilesSupported{}
	} else if f.Path != nil {
		return nil, nil, &ErrOnlyTextFilesSupported{}
	}

	var b bytes.Buffer
	opts := &FileContentOptions{At: at, ListOptions: ListOptions{Limit: resp.Limit}}
	for {
		for _, l := range f.Lines {
			if b.Len() > 0 {
				b.WriteString("\n")
			}
			b.WriteString(l.Text)
		}
		if resp.LastPage {
			break
		}
		opts.Start = resp.NextPageStart
		resp, err = s.client.GetPaged(ctx, projectsApiName, p, &f, opts)
		if err != nil {
			return nil, nil, err
		}
	}
	return b.Bytes(), resp, nil
}
