package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Forge Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Forge v0.1.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
