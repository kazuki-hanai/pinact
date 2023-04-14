package run

import (
	"context"
	"fmt"

	"github.com/suzuki-shunsuke/pinact/pkg/github"
)

type GitService interface {
	GetRef(ctx context.Context, owner string, repo string, ref string) (*github.Reference, *github.Response, error)
}

type GetRefResponse struct {
	Reference *github.Reference
	Response  *github.Response
	err       error
}

type GitServiceImpl struct {
	GitService GitService
	m          map[string]*GetRefResponse
}

func (gitService *GitServiceImpl) GetRef(ctx context.Context, owner string, repo string, ref string) (*github.Reference, *github.Response, error) {
	key := fmt.Sprintf("%s/%s/%s", owner, repo, ref)
	a, ok := gitService.m[key]
	if ok {
		return a.Reference, a.Response, a.err
	}
	r, resp, err := gitService.GitService.GetRef(ctx, owner, repo, ref)
	gitService.m[key] = &GetRefResponse{
		Reference: r,
		Response:  resp,
		err:       err,
	}
	return r, resp, err //nolint:wrapcheck
}

type RepositoriesService interface {
	ListTags(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.RepositoryTag, *github.Response, error)
}

type ListTagsResult struct {
	Tags     []*github.RepositoryTag
	Response *github.Response
	err      error
}

type RepositoriesServiceImpl struct {
	RepositoriesService RepositoriesService
	m                   map[string]*ListTagsResult
}

func (repositoriesService *RepositoriesServiceImpl) ListTags(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.RepositoryTag, *github.Response, error) {
	key := fmt.Sprintf("%s/%s/%v", owner, repo, opts.Page)
	a, ok := repositoriesService.m[key]
	if ok {
		return a.Tags, a.Response, a.err
	}
	tags, resp, err := repositoriesService.RepositoriesService.ListTags(ctx, owner, repo, opts)
	repositoriesService.m[key] = &ListTagsResult{
		Tags:     tags,
		Response: resp,
		err:      err,
	}
	return tags, resp, err //nolint:wrapcheck
}
