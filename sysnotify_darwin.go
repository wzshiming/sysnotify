package sysnotify

import (
	"fmt"
)

func Notify(title string, message string, iconPath string) error {
	return notify("-e", fmt.Sprintf("display notification %q with title %q", message, title))
}

func Alert(title string, message string, iconPath string) error {
	return notify("-e", fmt.Sprintf("display notification %q with title %q sound name %q", message, title, "Ping"))
}

func notify(args ...string) error {
	return command("osascript", args...)
}
