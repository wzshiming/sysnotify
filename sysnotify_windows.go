package sysnotify

func Notify(title string, message string, iconPath string) error {
	args := []string{}
	if iconPath != "" {
		args = append(args, "/i:", iconPath)
	}
	args = append(args, "/t:", title, message)
	return notify(args)
}

func Alert(title string, message string, iconPath string) error {
	args := []string{}
	if iconPath != "" {
		args = append(args, "/i:", iconPath)
	}
	args = append(args, "/t:", title, message, "/s", "true", "/p", "2")
	return notify(args)
}

func notify(args []string) error {
	return command("growlnotify", args...)
}
