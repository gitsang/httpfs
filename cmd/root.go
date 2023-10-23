package cmd

import (
	"os"

	"github.com/gitsang/httpfs/internal"
	"github.com/spf13/cobra"
)

var (
	listen string
	dir    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httpfs",
	Short: "Simple command-line tool for HTTP file server",
	Long:  `Simple command-line tool for HTTP file server.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Serve(listen, dir)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&listen, "listen", "l", ":8080", "server listen address")
	rootCmd.Flags().StringVarP(&dir, "directory", "d", ".", "the directory of static file to host")
}
