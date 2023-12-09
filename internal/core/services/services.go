package services

import (
	"errors"
	"github.com/realfabecker/bogo/internal/core/ports"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RepositoryService repository service struct definition
type RepositoryService struct {
	repo ports.ProjectRepository
	down ports.ProjectDownloader
	echo ports.Logger
}

// NewRepositoryService repository service construtor
func NewRepositoryService(
	repo ports.ProjectRepository,
	down ports.ProjectDownloader,
	echo ports.Logger,
) ports.ProjectService {
	return &RepositoryService{repo, down, echo}
}

// Install configure a project locally
func (r RepositoryService) Install(project string, name string) error {
	p, err := r.repo.Get(project)
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	pd := filepath.Join(wd, name)
	if err := r.down.Download(p, pd); err != nil {
		return err
	}
	if p.Scripts != nil && p.Scripts.InstallScript != nil {
		return r.runScript(*p.Scripts.InstallScript, pd)
	}
	return nil
}

// runScript run install script for repository
func (r RepositoryService) runScript(script string, dir string) error {
	r.echo.Infof("running install script: %s", script)
	args := strings.Split(script, " ")
	var cmd *exec.Cmd
	if len(args) > 1 {
		cmd = exec.Command(args[:1][0], args[1:]...)
	} else if len(args) == 1 {
		cmd = exec.Command(args[:1][0])
	} else {
		return errors.New("project install script is not valid")
	}
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
