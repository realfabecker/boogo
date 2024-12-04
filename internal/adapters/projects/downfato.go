package projects

import (
	"fmt"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/ports"
)

// NewFactory returns a downloader factory definition
func NewFactory() func(logger ports.Logger, t domain.ProjectType) (ports.ProjectDownloader, error) {
	return func(logger ports.Logger, t domain.ProjectType) (ports.ProjectDownloader, error) {
		if t == domain.TypeGithubRepo {
			return NewGithubRepoDownloader(logger), nil
		}
		return nil, fmt.Errorf("%s is not a valid downloader", t)
	}
}
