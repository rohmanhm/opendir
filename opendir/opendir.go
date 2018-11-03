package opendir

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Exec is command to execute.
// filemanager is name of the program of each platform support via cli.
// path is the destination of filemanager needs to open at.
// TODO: add test
func Exec(filemanager string, path string) {
	cmd := exec.Command(filemanager, path)
	err := cmd.Start()

	if err != nil {
		panic(err)
	}
}

// Open run action to open the filemanager by where executed program os from.
// TODO: add test
func Open(path string) {
	switch os := runtime.GOOS; os {
	case "darwin":
		OpenInDarwin(path)
	case "linux":
		OpenInLinux(path)
	case "windows":
		OpenInWindows(path)
	default:
		// TODO: Support more operating system.
		// freebsd, openbsd, plan9, ...
		fmt.Printf("%s Not supported yet.", os)
	}
}

// OpenInWindows is exec the path from windows OS.
// It's uses `explorer` from windows command prompt.
// TODO: add test
func OpenInWindows(path string) {
	Exec("explorer", path)
}

// OpenInLinux is exec the path from linux OS.
// It's uses `nautilus` from linux cli.
// TODO: add test
func OpenInLinux(path string) {
	Exec("nautilus", path)
}

// OpenInDarwin is exec the path from darwin OS.
// It's uses `open` from darwin cli.
// TODO: add test
func OpenInDarwin(path string) {
	Exec("open", path+" -a finder")
}
