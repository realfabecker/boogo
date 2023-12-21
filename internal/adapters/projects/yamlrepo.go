package projects

import (
	"fmt"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/ports"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// YamlProjectRepository memory project repository struct
type YamlProjectRepository struct {
	ports.Logger
}

// NewYamlProjectRepository YamlProjectRepository construtor
func NewYamlProjectRepository(logger ports.Logger) ports.ProjectRepository {
	return &YamlProjectRepository{logger}
}

// Get return a repository by its name
func (m YamlProjectRepository) Get(name string) (*domain.Project, error) {
	p, err := m.List()
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}
	for _, v := range p {
		if v.Name == name {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("%s is not a valid repository", name)
}

// List return a list of repositories
func (m YamlProjectRepository) List() ([]domain.Project, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	d, err := os.ReadFile(filepath.Join(h, ".bogo", "repositories.yaml"))
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}

	var r domain.RepoConfig
	if err := yaml.Unmarshal(d, &r); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("list: %w", err)
	}
	return r.Projects, nil
}

// Store record byte data repo config
func (m YamlProjectRepository) Store(data []byte) error {
	h, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(h, ".bogo", "repositories.yaml"), data, 0644)
}
