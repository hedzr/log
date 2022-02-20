package log

import "log"

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logger = newStdLogger()
}

var logger Logger
