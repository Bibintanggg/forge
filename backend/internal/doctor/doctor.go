package doctor

import (
	"os/exec"
	"runtime"
	"strings"
)

type CheckResult struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
	Status  string `json:"status"`
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

type Result struct {
	OS           string        `json:"os"`
	Architecture string        `json:"architecture"`
	Checks       []CheckResult `json:"checks"`
}

func Run() Result {
	var checks []CheckResult
	for _, tool := range tools {
		checks = append(checks, checkCommand(tool))
	}

	return Result{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		Checks:       checks,
	}
}

func checkCommand(tool CommandCheck) CheckResult {
	_, err := exec.LookPath(tool.Name)
	if err != nil {
		return CheckResult{
			Name:   tool.Name,
			Status: "not_found",
		}
	}

	output, err := exec.Command(tool.Name, tool.Args...).Output()

	if err != nil {
		return CheckResult{
			Name:   tool.Name,
			Status: "found",
		}
	}

	return CheckResult{
		Name:    tool.Name,
		Status:  "found",
		Version: strings.TrimSpace(string(output)),
	}
}
