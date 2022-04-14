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
	// log.Fatalf("")
	log.Printf("")
	log.SetLevel(InfoLevel)
	_ = log.GetLevel()
	log.Setup()
	tf(log)
	tp(log)

}

func tf(logger Logger) {
	tf1(logger)
	tf2(logger)
}

func tf1(logger Logger) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	logger.Fatalf("fatal")
}

func tf2(logger Logger) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	logger.(LoggerExt).Fatal("fatal")
}

func tp(logger Logger) {
	tp1(logger)
	tp2(logger)
}

func tp1(logger Logger) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	logger.Panicf("panic")
}

func tp2(logger Logger) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	logger.(LoggerExt).Panic("panic")
}
