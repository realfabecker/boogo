package services

import (
	"github.com/realfabecker/bogo/internal/core/ports"
	"os"
	"path/filepath"
)

// ConfService conf service struct definition
type ConfService struct {
	down ports.RepoConfigDownloader
	echo ports.Logger
}

// NewConfService conf service constructor
func NewConfService(
	down ports.RepoConfigDownloader,
	echo ports.Logger,
) ports.BogoConfigService {
	return &ConfService{down, echo}
}

// Sync synchronizes project configuration
func (r ConfService) Sync() error {
	r.echo.Info("Downloading repositories config")
	data, err := r.down.Download()
	if err != nil {
		return err
	}
	path, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(path, ".bogo", "repositories.json"), data, 0755)
}
