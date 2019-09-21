package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "v0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "iconizer",
	Short:   "Add icons to filenames by filetypes.",
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// main.main() calls this. It needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
