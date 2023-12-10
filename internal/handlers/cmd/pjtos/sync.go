package pjtos

import (
	"github.com/realfabecker/bogo/internal/adapters/config"
	"github.com/realfabecker/bogo/internal/adapters/logger"
	"github.com/realfabecker/bogo/internal/core/services"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// NewSyncCmd initialize a new project
func NewSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "config sync",
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			echo := logger.NewConsoleLogger("bogo", os.Stdout)
			serv := services.NewConfService(
				config.NewGistRepoConfigDownloader(url),
				echo,
			)
			return serv.Sync()
		},
	}
	cmd.Flags().String("url", "", "config download url")
	if err := cmd.MarkFlagRequired("url"); err != nil {
		log.Fatalln(err)
	}
	return cmd
}
