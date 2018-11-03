package cmd

import (
	"os"
	"path/filepath"

	"github.com/rohmanhm/go-open-dir/opendir"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "open-dir [DIR]",
	DisableFlagsInUseLine: true,
	Short:                 "Quick open a directory from cli.",
	Run: func(cmd *cobra.Command, args []string) {
		currentPath, err := filepath.Abs("./")

		var targetPath string

		if len(args) > 0 {
			targetPath = args[0]
		} else {
			targetPath = "./"
		}

		if err != nil {
			panic(err)
		}

		var fullPath string
		if filepath.IsAbs(targetPath) {
			fullPath = targetPath
		} else {
			fullPath = filepath.Join(currentPath, targetPath)
		}

		opendir.Open(fullPath)
	},
}

// Execute exec
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
