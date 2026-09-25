package commands

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/Bibintanggg/forge/internal/doctor"
	"github.com/spf13/cobra"
)

type CheckResult struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Version string `json:"version,omitempty"`
}

type CommandCheck struct {
	Name string
	Args []string
}

var tools = []CommandCheck{
	{
		Name: "go",
		Args: []string{"version"},
	},
	{
		Name: "git",
		Args: []string{"--version"},
	},
	{
		Name: "node",
		Args: []string{"--version"},
	},
	{
		Name: "docker",
		Args: []string{"--version"},
	},
}

var jsonOutput bool

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check your development environment",
	Run: func(cmd *cobra.Command, args []string) {
		result := doctor.Run()

		if jsonOutput {
			data, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(data))
			return
		}

		fmt.Println("Forge Doctor")
		fmt.Println()

		fmt.Printf("OS 				%s\n", runtime.GOOS)
		fmt.Printf("Architecture	%s\n", runtime.GOARCH)

		fmt.Println()

		for _, check := range result.Checks {
			if check.Status == "found" {
				fmt.Printf("✓ %-10s %s\n", check.Name, check.Version)
			} else {
				fmt.Printf("✗ %-10s %s\n", check.Name, check.Version)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doctorCmd)

	doctorCmd.Flags().BoolVar(
		&jsonOutput, "json", false, "Output results as JSON",
	)
}
