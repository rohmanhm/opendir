package app

import (
	"os"

	"github.com/rohmanhm/go-open-dir/internal/cli/root"
)

// Run cli
func Run() {
	if err := root.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
