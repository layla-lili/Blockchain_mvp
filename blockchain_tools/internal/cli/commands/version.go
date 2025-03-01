package commands

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/layla-lili/blockchain_tools/internal/cli/formatter"
	"github.com/spf13/cobra"
)

// Version information
var (
	Version   = "0.1.0"
	GitCommit = "unknown"
	BuildDate = "unknown"
	versionMu sync.RWMutex
)

func SetVersionInfo(version, commit, buildDate string) {
	versionMu.Lock()
	defer versionMu.Unlock()
	Version = version
	GitCommit = commit
	BuildDate = buildDate
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Long:  `Display version and build information for the blockchain CLI.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			versionInfo := struct {
				Version   string `json:"version"`
				GitCommit string `json:"gitCommit"`
				BuildDate string `json:"buildDate"`
				GoVersion string `json:"goVersion"`
				Platform  string `json:"platform"`
			}{
				Version:   Version,
				GitCommit: GitCommit,
				BuildDate: BuildDate,
				GoVersion: runtime.Version(),
				Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
			}

			format, _ := cmd.Flags().GetString("format")
			fmt := formatter.GetFormatter(format)
			return fmt.Format(cmd.OutOrStdout(), versionInfo)
		},
	}
}
