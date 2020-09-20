// Copyright © 2020 Hedzr Yeh.

package log

import (
	"fmt"
	"log"
)

// NewDummyLogger return a dummy logger
func NewDummyLogger() Logger {
	return &dummyLogger{}
}

// NewStdLogger return a stdlib `log` logger
func NewStdLogger() Logger {
	return &stdLogger{Level: InfoLevel}
}

// NewStdLoggerWith return a stdlib `log` logger
func NewStdLoggerWith(lvl Level) Logger {
	return &stdLogger{Level: lvl}
}

type dummyLogger struct{}

func (d *dummyLogger) Tracef(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Debugf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Infof(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Warnf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Errorf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Fatalf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Panicf(msg string, args ...interface{}) {
	panic(fmt.Sprintf(msg, args...))
}

func (d *dummyLogger) Printf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) SetLevel(lvl Level) {
	//panic("implement me")
}

func (d *dummyLogger) GetLevel() Level {
	//panic("implement me")
	return InfoLevel // logex.GetLevel()
}

func (d *dummyLogger) Setup() {
	//panic("implement me")
}

type stdLogger struct {
	Level
}

func (s *stdLogger) Tracef(msg string, args ...interface{}) {
	if s.Level >= TraceLevel {
		log.Printf(msg, args...)
	}
}

func (s *stdLogger) Debugf(msg string, args ...interface{}) {
	if s.Level >= DebugLevel {
		log.Printf(msg, args...)
	}
}

func (s *stdLogger) Infof(msg string, args ...interface{}) {
	if s.Level >= InfoLevel {
		log.Printf(msg, args...)
	}
}

func (s *stdLogger) Warnf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func (s *stdLogger) Errorf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func (s *stdLogger) Fatalf(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}

func (s *stdLogger) Panicf(msg string, args ...interface{}) {
	log.Panicf(msg, args...)
}

func (s *stdLogger) Printf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func (s *stdLogger) SetLevel(lvl Level) {
	s.Level = lvl
}

func (s *stdLogger) GetLevel() Level {
	return s.Level
}

func (s *stdLogger) Setup() {
}

//
//
//
