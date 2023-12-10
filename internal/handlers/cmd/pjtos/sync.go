package pjtos

import (
	"errors"
	"github.com/realfabecker/bogo/internal/adapters/config"
	"github.com/realfabecker/bogo/internal/adapters/logger"
	"github.com/realfabecker/bogo/internal/adapters/projects"
	"github.com/spf13/cobra"
	"os"
)

// NewSyncCmd initialize a new project
func NewSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sync",
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")

			repo := config.NewJsonConfigRepository()
			conf, err := repo.Get()
			if err != nil {
				return errors.New("unable to get config from repository")
			}

			if url == "" && conf.RepoUrl != "" {
				url = conf.RepoUrl
			} else if url == "" {
				return errors.New("download url is required for repo synchronization")
			}

			echo := logger.NewConsoleLogger("bogo", os.Stdout)
			echo.Info("Syncing config from url source")
			return projects.NewJsonProjectRepository(echo).Sync(url)
		},
	}
	cmd.Flags().String("url", "", "url for download repo config from")
	return cmd
}
