// Copyright (c) 2018 MH Rohman Masyhar
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

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

// ExecString return execf isolated with config path
func (c *Config) ExecString() string {
	return fmt.Sprintf(c.execf, c.Path)
}

// ExecArray get execf full string command splitted in array
func (c *Config) ExecArray() []string {
	return strings.Split(c.ExecString(), " ")
}

// AppName get the app name from execf
func (c *Config) AppName() string {
	if len(c.ExecArray()) > 0 {
		return c.ExecArray()[0]
	}

	return ""
}

// RestArgs get the rest args from execf
func (c *Config) RestArgs() []string {
	if len(c.ExecArray()) > 1 {
		return c.ExecArray()[1:]
	}

	return []string{}
}

// Open the filemanager
// TODO: add test
func (c *Config) Open() error {
	cmd := exec.Command(c.AppName(), c.RestArgs()...)

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
