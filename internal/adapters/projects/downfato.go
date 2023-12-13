package projects

import (
	"fmt"
	"github.com/realfabecker/bogo/internal/adapters/github"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/ports"
)

// NewFactory returns a downloader factory definition
func NewFactory(config *domain.Config) func(logger ports.Logger, t domain.ProjectType) (ports.ProjectDownloader, error) {
	return func(logger ports.Logger, t domain.ProjectType) (ports.ProjectDownloader, error) {
		if t == domain.TypeGithubRepo {
			return NewGithubRepoDownloader(logger), nil
		}
		if t == domain.TypeGithubGist {
			api := github.NewApi(config.RepoAuth)
			return NewGithubGistDownloader(logger, api), nil
		}
		return nil, fmt.Errorf("%s is not a valid downloader", t)
	}
}
