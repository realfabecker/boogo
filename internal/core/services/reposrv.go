package services

import (
	"fmt"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/ports"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// RepositoryService repository service struct definition
type RepositoryService struct {
	repo ports.ProjectRepository
	echo ports.Logger
	fato ports.DownloaderFactory
}

// NewRepositoryService repository service construtor
func NewRepositoryService(
	repo ports.ProjectRepository,
	fato ports.DownloaderFactory,
	echo ports.Logger,
) ports.ProjectService {
	return &RepositoryService{repo, echo, fato}
}

// Install configure a project locally
func (r RepositoryService) Install(project string, name string) error {
	p, err := r.repo.Get(project)
	if err != nil {
		return err
	}

	down, err := r.fato(r.echo, p.Type)
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	if p.Type == domain.TypeGithubRepo {
		dir = filepath.Join(dir, name)
	}

	if err := down.Download(p, dir); err != nil {
		return fmt.Errorf("download: %w", err)
	}

	if p.Scripts != nil && p.Scripts.Install != nil {
		if err := r.runScript(*p.Scripts.Install, dir); err != nil {
			return fmt.Errorf("script: %w", err)
		}
	}
	return nil
}

// runScript run install script for repository
func (r RepositoryService) runScript(script string, dir string) error {
	var cmd *exec.Cmd
	if strings.Contains(runtime.GOOS, "windows") {
		cmd = exec.Command("cmd", "/C", script)
	} else {
		cmd = exec.Command("bash", "-c", script)
	}
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
