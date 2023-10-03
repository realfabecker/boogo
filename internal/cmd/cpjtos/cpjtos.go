package cpjtos

import (
	"github.com/spf13/cobra"
)

// New constrói base de comandos do projeto
func New() *cobra.Command {
	CPjtosCmd := &cobra.Command{
		Use:   "pjtos",
		Short: "Base para projetos",
	}
	CPjtosCmd.AddCommand(newNodeTsCmd())
	return CPjtosCmd
}
