package main

import (
	"os"

	"github.com/layla-lili/blockchain_tools/internal/cli/commands"
)

var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

func main() {
	commands.SetVersionInfo(Version, GitCommit, BuildDate)
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}
