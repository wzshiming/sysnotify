// +build !windows,!darwin,!linux

package sysnotify

func Notify(title string, message string, iconPath string) error {
	return nil
}

func Alert(title string, message string, iconPath string) error {
	return nil
}
