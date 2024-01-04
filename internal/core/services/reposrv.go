package services

import (
	"bufio"
	"fmt"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/ports"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type RepositoryService struct {
	repo ports.ProjectRepository
	echo ports.Logger
	fato ports.DownloaderFactory
}

func NewRepositoryService(
	repo ports.ProjectRepository,
	fato ports.DownloaderFactory,
	echo ports.Logger,
) ports.ProjectService {
	return &RepositoryService{repo, echo, fato}
}

func (r RepositoryService) Install(project string, name string) error {
	p, err := r.repo.Get(project)
	if err != nil {
		return err
	}

	if err := r.parseVars(p); err != nil {
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

func (r RepositoryService) parseVars(p *domain.Project) error {
	x := bufio.NewReader(os.Stdin)
	for _, v := range p.Vars {
		if v.Type == domain.VarTypeStdin {
			r.echo.Infof("Forneça valor para: %s", v.Description)
			l, _, _ := x.ReadLine()
			v.Value = string(l)
		} else {
			return fmt.Errorf("%s is not a valid variable input method", v)
		}

	}
	return nil
}

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
