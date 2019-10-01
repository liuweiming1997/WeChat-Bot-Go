package lib

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/WeChat-Bot-Go/logger"
)

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

// Open calls the OS default program for uri
func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		logger.Error(fmt.Sprintf("don't know how to open things on %s platform", runtime.GOOS))
		return errors.New("Can not open in this system")
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}
