package opendir

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Config opendir
type Config struct {
	// Path folder to open
	Path string
}

// Open run action to open the filemanager by where executed program OS from.
// TODO: add test
func (c *Config) Open() {
	switch os := runtime.GOOS; os {
	case "darwin":
		c.OpenInDarwin()
	case "linux":
		c.OpenInLinux()
	case "windows":
		c.OpenInWindows()
	default:
		// TODO: Support more operating system.
		fmt.Printf("%s Not supported yet.", os)
	}
}

// OpenInWindows is exec the path from windows OS.
// It's uses `explorer` from windows command prompt.
// TODO: add test
func (c *Config) OpenInWindows() {
	Exec("explorer", c.Path)
}

// OpenInLinux is exec the path from linux OS.
// It's uses `xdg-open` from linux cli.
// TODO: add test
func (c *Config) OpenInLinux() {
	Exec("xdg-open", c.Path)
}

// OpenInDarwin is exec the path from darwin OS.
// It's uses `open` from darwin cli.
// TODO: add test
func (c *Config) OpenInDarwin() {
	Exec("open", c.Path, "-a finder")
}

// Exec is command to execute.
// filemanager is name of the program of each platform support via cli.
// path is the destination of filemanager needs to open at.
// TODO: add test
func Exec(filemanager string, path string, execArgs ...string) {
	cmdArgs := append([]string{path}, execArgs...)
	cmd := exec.Command(filemanager, cmdArgs...)
	err := cmd.Start()

	if err != nil {
		panic(err)
	}
}
