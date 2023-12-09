package ports

import "github.com/realfabecker/bogo/internal/core/domain"

// ProjectRepository repository interface definition
type ProjectRepository interface {
	Get(name string) (*domain.Project, error)
	List() ([]domain.Project, error)
}

// ProjectDownloader project downloader interface
type ProjectDownloader interface {
	Download(repo *domain.Project, dir string) error
}

// ProjectService service interface definition
type ProjectService interface {
	Install(project string, name string) error
}

// Logger debugger interface
type Logger interface {
	Info(message string)
	Infof(format string, a ...interface{})
	Error(message string)
	Errorf(format string, a ...interface{})
	Debug(message string)
}
