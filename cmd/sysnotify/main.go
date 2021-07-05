package main

import (
	"flag"
	"log"

	"github.com/wzshiming/sysnotify"
)

var (
	title   string
	message string
)

func init() {
	flag.StringVar(&title, "t", "title", "notify title")
	flag.StringVar(&message, "m", "message", "notify message")
}

func main() {
	err := sysnotify.Alert(title, message, "")
	if err != nil {
		log.Println(err)
	}
}
