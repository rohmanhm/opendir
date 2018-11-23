package path

import "path/filepath"

// GetValueFromArray the targetted target path from given args
func GetValueFromArray(args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	return "./"
}

// Resolve the full path of target path
func Resolve(currentPath string, targetPath string) string {
	if filepath.IsAbs(targetPath) {
		return targetPath
	}

	return filepath.Join(currentPath, targetPath)
}
