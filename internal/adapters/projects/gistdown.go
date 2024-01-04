package projects

import (
	"github.com/realfabecker/bogo/internal/adapters/github"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/ports"
	"strings"
)

// GithubGistDownloader GitHub struct service definition
type GithubGistDownloader struct {
	api    github.Api
	logger ports.Logger
}

// NewGithubGistDownloader GitHub service construtor
func NewGithubGistDownloader(logger ports.Logger, api github.Api) ports.ProjectDownloader {
	return &GithubGistDownloader{logger: logger, api: api}
}

// Download project install by its struct definition
func (s GithubGistDownloader) Download(repo *domain.Project, dir string) error {
	s.logger.Infof("obtaining gist configuration")
	p := strings.Split(repo.Url, "/")
	gist, err := s.api.GetGist(p[len(p)-1])
	if err != nil {
		return err
	}
	s.logger.Infof("downloading gist content")
	return s.api.Download(repo, gist, dir)
}
