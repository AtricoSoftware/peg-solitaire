// Generated 2021-01-19 11:42:11 by go-framework v1.2.1-1-ge87b524
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/AtricoSoftware/peg-solitaire/pkg"
)

var fullVersion bool

func CreateVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Shows version",
		Run: func(*cobra.Command, []string) {
			showVersion(fullVersion)
		},
	}
	cmd.PersistentFlags().BoolVarP(&fullVersion, "full", "f", false, "Full program information")
	return cmd
}

func showVersion(fullVersion bool) {
	if fullVersion {
		fmt.Println(pkg.Name)
		fmt.Println(pkg.Description)
	}
	fmt.Println(pkg.Version)
	if fullVersion {
		fmt.Println()
		var details map[string]interface{}
		if err := json.Unmarshal([]byte(pkg.BuildDetails), &details); err == nil && len(details) > 0 {
			fmt.Println("Details")
			fmt.Println("-------")
			displaySection(details, "")
		}
	}
}

func displaySection(section map[string]interface{}, indent string) {
	for k, v := range section {
		fmt.Printf("%s%s:", indent, k)
		switch v.(type) {
		case map[string]interface{}:
			fmt.Println()
			displaySection(v.(map[string]interface{}), indent+"  ")
		default:
			fmt.Printf(" %s\n", v)
		}
	}
}
