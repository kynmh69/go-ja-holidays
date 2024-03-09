package util

import "log"

func LoggerInitialize() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[App] ")
}
