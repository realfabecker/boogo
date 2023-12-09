package projects

import (
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/ports"
	"os"
	"os/exec"
	"path/filepath"
)

// GithubProjectDownloader GitHub struct service definition
type GithubProjectDownloader struct {
	logger ports.Logger
}

// NewGithubProjectDownloader GitHub service construtor
func NewGithubProjectDownloader(logger ports.Logger) ports.ProjectDownloader {
	return &GithubProjectDownloader{logger: logger}
}

// Download project install by its struct definition
func (s GithubProjectDownloader) Download(project *domain.Project, dir string) error {
	if err := s.clone(project.Url, dir); err != nil {
		return err
	}
	return os.RemoveAll(filepath.Join(dir, ".git"))
}

// clone obtains a GitHub repository and check it out locally
func (s GithubProjectDownloader) clone(url string, dir string) error {
	s.logger.Infof("cloning %s into %s", url, dir)
	cmd := exec.Command("git", "clone", url, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
