package utils

import "path/filepath"

// GetTargetPathFromArgs the targetted target path from given args
func GetTargetPathFromArgs(args []string) (targetPath string) {
	if len(args) > 0 {
		targetPath = args[0]
	} else {
		targetPath = "./"
	}

	return
}

// ResolveFullPath the full path of target path
func ResolveFullPath(currentPath string, targetPath string) (fullPath string) {
	if filepath.IsAbs(targetPath) {
		fullPath = targetPath
	} else {
		fullPath = filepath.Join(currentPath, targetPath)
	}

	return
}
