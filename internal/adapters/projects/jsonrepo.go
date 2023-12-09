package projects

import (
	"fmt"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/entities"
	"github.com/realfabecker/bogo/internal/core/ports"
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
	x, ok := entities.Projects[name]
	if !ok {
		return nil, fmt.Errorf("%s is not a valid project", name)
	}
	return &x, nil
}

// List return a list of repositories
func (m JsonProjectRepository) List() ([]domain.Project, error) {
	r := make([]domain.Project, len(entities.Projects))
	for _, v := range entities.Projects {
		r = append(r, v)
	}
	return r, nil
}
