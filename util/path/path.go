package path

import "path/filepath"

// Resolve the full path of target path
func Resolve(currentPath string, targetPath string) string {
	if filepath.IsAbs(targetPath) {
		return targetPath
	}

	return filepath.Join(currentPath, targetPath)
}
