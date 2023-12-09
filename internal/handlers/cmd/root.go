package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

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
	cmd.AddCommand(newIniCmd())
	return cmd
}

// Execute cmd base invocation
func Execute() {
	if err := newBogoCmd().Execute(); err != nil {
		log.Fatalln(err)
	}
}
