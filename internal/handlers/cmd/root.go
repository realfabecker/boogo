package cmd

import (
	"errors"
	"github.com/realfabecker/bogo/internal/handlers/cmd/pjtos"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

// init app project boostrap
func init() {
	h, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	d, err := os.Open(filepath.Join(h, ".bogo"))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalln(err)
	}
	if d != nil {
		return
	}
	if err := os.Mkdir(filepath.Join(h, ".bogo"), 0755); err != nil {
		log.Fatalln(err)
	}
}

// newBogoCmd bogo cmd constructor
func newBogoCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:           "bogo [command]",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(pjtos.NewIniDCmd())
	return cmd
}

// Execute cmd base invocation
func Execute() {
	if err := newBogoCmd().Execute(); err != nil {
		log.Fatalln(err)
	}
}
