// Copyright (c) 2018 MH Rohman Masyhar
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rohmanhm/opendir"
	"github.com/rohmanhm/opendir/util/path"
)

// AppVersion of the app
const AppVersion = "0.1.3"

// Usage cli message
const Usage = `
opendir is simple cross platform open directory.

Usage:
  opendir [Flags] [directories...]

Flags:
  version		prints app current version.

Examples:
  - Open current directory. By default, 'opendir' will open the current active directory.
	'opendir'

  - Open sub directory
	'opendir subdirectory'
	'opendir ./subdirectory'

  - Open specified directory path
	'opendir ~/Downloads'

  - Open multiple directories at once
	'opendir subdirectory ~/Downloads'
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, Usage)
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(2)
	}

	flag.Parse()

	// prints the app version
	if flag.Arg(0) == "version" {
		fmt.Printf("opendir version v%s\n", AppVersion)
		os.Exit(0)
	}

	var paths []string
	paths = flag.Args()

	// If no path specified in args
	// Open the current directory
	if len(paths) < 1 {
		paths = append(paths, "./")
	}

	var currentPath, err = filepath.Abs("./")
	if err != nil {
		panic(err)
	}

	for _, path := range paths {
		err := openPath(currentPath, path)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func openPath(currentPath string, targetPath string) error {
	fullPath := path.Resolve(currentPath, targetPath)

	odir, err := opendir.New(fullPath)
	if err != nil {
		return err
	}

	err = odir.Open()
	if err != nil {
		return err
	}

	return nil
}
