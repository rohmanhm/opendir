package root

import (
	"fmt"
	"path/filepath"

	opendir "github.com/rohmanhm/go-open-dir"
	"github.com/rohmanhm/go-open-dir/internal/utils"
	"github.com/spf13/cobra"
)

// Command root
type Command = cobra.Command

// Cmd root
var Cmd = &Command{
	Use:                   "open-dir [DIR]",
	DisableFlagsInUseLine: true,
	Short:                 "Quick open a directory from cli.",
	Run: func(cmd *cobra.Command, args []string) {
		var currentPath, err = filepath.Abs("./")
		if err != nil {
			panic(err)
		}

		var targetPath = utils.GetTargetPathFromArgs(args)
		var fullPath = utils.ResolveFullPath(currentPath, targetPath)

		odir, err := opendir.New(fullPath)
		if err != nil {
			fmt.Println(err)
		}

		err = odir.Open()

		if err != nil {
			fmt.Println(err)
		}
	},
}
