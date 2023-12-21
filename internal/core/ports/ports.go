package ports

import "github.com/realfabecker/bogo/internal/core/domain"

// ProjectRepository repository interface definition
type ProjectRepository interface {
	Get(name string) (*domain.Project, error)
	List() ([]domain.Project, error)
	Store(data []byte) error
}

// ConfigRepository config repository interface
type ConfigRepository interface {
	Get() (*domain.BogoConfig, error)
	Save(config *domain.BogoConfig) error
}

// ProjectDownloader project downloader interface
type ProjectDownloader interface {
	Download(repo *domain.Project, dir string) error
}

// RepoConfigDownloader repo downloader interface
type RepoConfigDownloader interface {
	Download(url string) ([]byte, error)
}

// BogoConfigService service interface definition
type BogoConfigService interface {
	Sync() error
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

// JsonValidator json validator interface
type JsonValidator interface {
	Validate(data []byte, schema string) (bool, error)
}

// DownloaderFactory download factory function definition
type DownloaderFactory func(logger Logger, t domain.ProjectType) (ProjectDownloader, error)
