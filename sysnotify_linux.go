package sysnotify

func Notify(title string, message string, iconPath string) error {
	args := []string{}
	if iconPath != "" {
		args = append(args, "-i", iconPath)
	}
	args = append(args, title, message)
	return notify(args)
}

func Alert(title string, message string, iconPath string) error {
	args := []string{}
	if iconPath != "" {
		args = append(args, "-i", iconPath)
	}
	args = append(args, title, message, "-u", "critical")
	return notify(args)
}

func notify(args []string) error {
	return command("notify-send", args...)
}
