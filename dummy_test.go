package log

import (
	"fmt"
	"testing"
)

func TestNewLoggerConfig(t *testing.T) {
	log := NewDummyLogger()
	log.Tracef("")
	log.Debugf("")
	log.Infof("")
	log.Warnf("")
	log.Errorf("")
	log.Fatalf("")
	log.Printf("")
	log.SetLevel(InfoLevel)
	_ = log.GetLevel()
	log.Setup()
	tf(log)
	tp(log)

	log = NewStdLogger()
	log.SetLevel(TraceLevel)
	log.Tracef("")
	log.Debugf("")
	log.Infof("")
	log.Warnf("")
	log.Errorf("")
	// log.Fatalf("")

	log = NewStdLoggerWith(TraceLevel)
	log.Printf("")
	log.SetLevel(InfoLevel)
	_ = log.GetLevel()
	log.Setup()
	tf(log)
	tp(log)
}

func tf(logger Logger) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	logger.Fatalf("fatal")
}

func tp(logger Logger) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	logger.Panicf("fatal")
}
