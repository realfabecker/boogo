package pjtos

import (
	"fmt"
	"github.com/realfabecker/bogo/internal/adapters/config"
	"github.com/realfabecker/bogo/internal/adapters/github"
	"github.com/realfabecker/bogo/internal/adapters/jsonx"
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
			repo := config.NewJsonConfigRepository()
			conf, err := repo.Get()
			if err != nil {
				return fmt.Errorf("get: %w", err)
			}

			down := github.NewGistRepoConfigDownloader(
				github.NewApi(conf.RepoAuth),
				jsonx.NewValidator(),
			)

			data, err := down.Download(conf.RepoUrl)
			if err != nil {
				return fmt.Errorf("download: %w", err)
			}

			echo := logger.NewConsoleLogger("bogo", os.Stdout)
			rep := projects.NewJsonProjectRepository(echo)
			return rep.Store(data)
		},
	}
	cmd.Flags().String("url", "", "url for download repo config from")
	return cmd
}
