package cmd

import (
	"github.com/realfabecker/boogo/internal/cmd/cpjtos"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "boogo [command]",
	Long: `
	    	       Boogo
				   
Interface para criação de projetos por base síntese.

De modo a ver o texto de ajuda, você pode executar:

   boogo <command> -h
   boogo <command> <subcommand> -h
`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	rootCmd.AddCommand(cpjtos.New())
}
