package main

import (
	"fmt"
	"os"

	"github.com/Bibintanggg/forge/cmd/forge/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
