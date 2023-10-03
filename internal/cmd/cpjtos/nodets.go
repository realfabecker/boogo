package cpjtos

import (
	pjtos "github.com/realfabecker/boogo/internal"
	"github.com/realfabecker/boogo/internal/lib/git"
	"github.com/spf13/cobra"
)

// newNodeTsCmd constrói base comando node ts
func newNodeTsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodets",
		Short: "Base para projetos NodeJS",
		RunE: func(cmd *cobra.Command, args []string) error {
			nome, _ := cmd.Flags().GetString("nome")
			p := pjtos.Projetos[pjtos.NODE_TS]
			return git.Checkout(p.Url, nome)
		},
	}
	cmd.Flags().StringP("nome", "n", "", "Nome do projeto a ser criado")
	cmd.MarkFlagRequired("nome")
	return cmd
}
