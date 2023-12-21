package pjtos

import (
	"github.com/realfabecker/bogo/internal/adapters/config"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/spf13/cobra"
)

// NewConfCmd configuration interface definition
func NewConfCmd() *cobra.Command {
	var cnf domain.BogoConfig
	cmd := &cobra.Command{
		Use:   "conf",
		Short: "bogo config interface",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo := config.NewJsonConfigRepository()
			return repo.Save(&cnf)
		},
	}
	cmd.Flags().StringVar(&cnf.RepoUrl, "repo-url", "", "repo config download url")
	cmd.Flags().StringVar(&cnf.RepoAuth, "repo-auth", "", "repo authorization download url")
	return cmd
}
