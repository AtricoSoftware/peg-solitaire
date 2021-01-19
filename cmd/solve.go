// Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
package cmd

import (
	"github.com/atrico-go/container"
	"github.com/spf13/cobra"

	"github.com/AtricoSoftware/peg-solitaire/api"
)

func CreateSolveCommand(c container.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "solve",
		Short: "Solve board",
		RunE: func(*cobra.Command, []string) error {
			var solveApi api.SolveApi
			c.Make(&solveApi)
			return solveApi.Run()
		},
	}
	return cmd
}
