package sysnotify

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func command(name string, arg ...string) error {
	c := exec.Command(name, arg...)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%q: %w: %q", strings.Join(append([]string{name}, arg...), " "), err, out)
	}
	return nil
}
