package cmd

import (
	"runtime"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var (
	version   = "dev"
	buildDate = "1970-01-01T00:00:00Z"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		tbl := table.New("Version", version)
		tbl.AddRow("Go Version", runtime.Version())
		tbl.AddRow("BuildDate", buildDate)
		tbl.AddRow("OS/Arch", runtime.GOOS+"/"+runtime.GOARCH)
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
