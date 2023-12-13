package github

import (
	"fmt"
	"github.com/realfabecker/bogo/internal/core/entities"
	"github.com/realfabecker/bogo/internal/core/ports"
	"strings"
)

// GistRepoConfigDowloader config downloader
type GistRepoConfigDowloader struct {
	api Api
	jsx ports.JsonValidator
}

// NewGistRepoConfigDownloader struct api constructor
func NewGistRepoConfigDownloader(api Api, jsx ports.JsonValidator) ports.RepoConfigDownloader {
	return &GistRepoConfigDowloader{api: api, jsx: jsx}
}

// Download obtain repositories configuration from gist config
func (g GistRepoConfigDowloader) Download(url string) ([]byte, error) {
	p := strings.Split(url, "/")

	gist, err := g.api.GetGist(p[len(p)-1])
	if err != nil {
		return nil, err
	}

	data, err := g.api.GetFile(gist, "repositories.json")
	if err != nil {
		return nil, fmt.Errorf("get-file: %w", err)
	}

	if _, err := g.jsx.Validate(data, entities.ProjectListSchema); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	return data, nil
}
