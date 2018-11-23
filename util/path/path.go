// Copyright (c) 2018 MH Rohman Masyhar
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package path

import "path/filepath"

// Resolve the full path of target path
func Resolve(currentPath string, targetPath string) string {
	if filepath.IsAbs(targetPath) {
		return targetPath
	}

	return filepath.Join(currentPath, targetPath)
}
