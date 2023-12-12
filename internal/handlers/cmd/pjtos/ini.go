package pjtos

import (
	"github.com/realfabecker/bogo/internal/adapters/logger"
	"github.com/realfabecker/bogo/internal/adapters/projects"
	"github.com/realfabecker/bogo/internal/core/services"
	"github.com/spf13/cobra"
	"os"
)

// NewIniCmd initialize a new project
func NewIniCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "init",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			echo := logger.NewConsoleLogger("bogo", os.Stdout)
			serv := services.NewRepositoryService(
				projects.NewJsonProjectRepository(echo),
				projects.NewGithubRepoDownloader(echo),
				echo,
			)
			if len(args) == 1 {
				return serv.Install(args[0], args[0])
			}
			return serv.Install(args[0], args[1])
		},
	}
	return cmd
}
