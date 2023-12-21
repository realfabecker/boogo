package github

import (
	"encoding/json"
	"fmt"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/entities"
	"github.com/realfabecker/bogo/internal/core/ports"
	"gopkg.in/yaml.v3"
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

	data, err := g.api.GetFile(gist, "repositories.yaml")
	if err != nil {
		return nil, fmt.Errorf("get-file: %w", err)
	}

	var x domain.RepoConfig
	if err := yaml.Unmarshal(data, &x); err != nil {
		return nil, fmt.Errorf("yaml: %w", err)
	}

	d, err := json.MarshalIndent(x, "", " ")
	if err != nil {
		return nil, fmt.Errorf("json: %w", err)
	}

	if _, err := g.jsx.Validate(d, entities.RepoConfigSchema); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}
	return data, nil
}
