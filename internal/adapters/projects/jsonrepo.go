package projects

import (
	"encoding/json"
	"fmt"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/ports"
	"os"
	"path/filepath"
)

// JsonProjectRepository memory project repository struct
type JsonProjectRepository struct {
	ports.Logger
}

// NewJsonProjectRepository JsonProjectRepository construtor
func NewJsonProjectRepository(logger ports.Logger) ports.ProjectRepository {
	return &JsonProjectRepository{logger}
}

// Get return a repository by its name
func (m JsonProjectRepository) Get(name string) (*domain.Project, error) {
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
func (m JsonProjectRepository) List() ([]domain.Project, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	d, err := os.ReadFile(filepath.Join(h, ".bogo", "repositories.json"))
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}
	r := make([]domain.Project, 0)
	if err := json.Unmarshal(d, &r); err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}
	return r, nil
}
