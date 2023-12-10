package pjtos

import (
	"github.com/realfabecker/bogo/internal/adapters/config"
	"github.com/realfabecker/bogo/internal/core/domain"
	"github.com/spf13/cobra"
	"log"
)

// NewConfCmd configuration interface definition
func NewConfCmd() *cobra.Command {
	var cnf domain.Config
	cmd := &cobra.Command{
		Use: "conf",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo := config.NewJsonConfigRepository()
			return repo.Save(&cnf)
		},
	}
	cmd.Flags().StringVar(&cnf.RepoUrl, "repo-url", "", "repo config download url")
	if err := cmd.MarkFlagRequired("repo-url"); err != nil {
		log.Fatalln(err)
	}
	return cmd
}
