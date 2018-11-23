package opendir

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// Config opendir
type Config struct {
	// execf is command string
	execf string
	// Path the folder
	Path string
}

// Open the filemanager
// TODO: add test
func (c *Config) Open() error {
	cmdStrArr := strings.Split(
		fmt.Sprintf(c.execf, c.Path), // full cmd string command
		" ")

	cmd := exec.Command(
		cmdStrArr[0],     // put the cmd name
		cmdStrArr[1:]..., // put the rest of cmd args
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("could not open dir %s\n%v", c.Path, string(output)+fmt.Sprint(err))
	}

	return nil
}

// New opendir
func New(path string) (*Config, error) {
	var cfg Config
	cfg.Path = path

	switch os := runtime.GOOS; os {
	case "darwin":
		cfg.execf = "open %s -a finder"
	case "linux":
		cfg.execf = "xdg-open %s"
	case "windows":
		cfg.execf = "explorer %s"
	default:
		// TODO: Support more operating system.
		return nil, fmt.Errorf("os %s is not supported yet", os)
	}

	return &cfg, nil
}
