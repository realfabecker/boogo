package pjtos

import (
	"os"

	"github.com/realfabecker/bogo/internal/adapters/logger"
	"github.com/realfabecker/bogo/internal/adapters/projects"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/realfabecker/bogo/internal/core/services"
	"github.com/spf13/cobra"
)

// newPojoCmd base project command  definition
func newPojoCmd(p domain.Project) *cobra.Command {
	cmd := &cobra.Command{
		Use:   p.GetUse(),
		Short: p.Description,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			echo := logger.NewConsoleLogger("bogo", os.Stdout)
			serv := services.NewRepositoryService(
				projects.NewYamlProjectRepository(echo),
				projects.NewFactory(),
				echo,
			)
			if len(args) == 0 {
				return serv.Install(p.Name, p.Name)
			}
			return serv.Install(p.Name, args[0])
		},
	}
	return cmd
}

// NewIniDCmd ini command initializer
func NewIniDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "bogo project init",
	}
	echo := logger.NewConsoleLogger("bogo", os.Stdout)
	repo := projects.NewYamlProjectRepository(echo)
	var p []domain.Project
	if p, _ = repo.List(); len(p) == 0 {
		return cmd
	}
	for _, v := range p {
		func(v domain.Project) {
			cmd.AddCommand(newPojoCmd(v))
		}(v)
	}
	return cmd
}
