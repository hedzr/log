package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// newStdLogger return a stdlib `log` logger
func newStdLogger() Logger {
	return &stdLogger{Level: InfoLevel, skip: 1, fields: make(map[string]interface{})}
}

// newStdLoggerWith return a stdlib `log` logger
func newStdLoggerWith(lvl Level) Logger {
	return &stdLogger{Level: lvl, skip: 1, fields: make(map[string]interface{})}
}

// newStdLoggerWithConfig return a stdlib `log` logger
func newStdLoggerWithConfig(config *LoggerConfig) Logger { //nolint:deadcode,unused
	l, _ := ParseLevel(config.Level)
	return &stdLogger{Level: l, skip: 1, fields: make(map[string]interface{})}
}

type stdLogger struct {
	Level
	skip   int
	fields map[string]interface{}
}

// extraSkipFramesFromLogPackage used for hedzr/log package functions:
//    log.Printf, log.Infof, ...
const extraSkipFramesFromLogPackage = 1
const skipFrames = 2 + extraSkipFramesFromLogPackage

func (s *stdLogger) AddSkip(skip int) Logger { return &stdLogger{Level: s.Level, skip: s.skip + skip} }

func (s *stdLogger) out(args ...interface{}) {
	str := fmt.Sprint(args...)
	_ = log.Output(skipFrames+s.skip, str)
}

func (s *stdLogger) With(key string, val interface{}) Logger {
	s.fields[key] = val
	return s
}

func (s *stdLogger) Trace(args ...interface{}) {
	if s.Level >= TraceLevel {
		s.out(args...)
	}
}

func (s *stdLogger) Debug(args ...interface{}) {
	if s.Level >= DebugLevel {
		s.out(args...)
	}
}

func (s *stdLogger) Info(args ...interface{}) {
	if s.Level >= InfoLevel {
		s.out(args...)
	}
}

func (s *stdLogger) Warn(args ...interface{}) {
	s.out(args...)
}

func (s *stdLogger) Error(args ...interface{}) {
	s.out(args...)
}

func (s *stdLogger) Fatal(args ...interface{}) {
	s.out(args...)
	if InTesting() {
		panic(fmt.Sprint(args...))
	}
	os.Exit(1)
}

func (s *stdLogger) Panic(args ...interface{}) {
	str := fmt.Sprint(args...)
	_ = log.Output(skipFrames+s.skip, str)
	panic(str)
}

func (s *stdLogger) Print(args ...interface{}) {
	s.out(args...)
}

func (s *stdLogger) Println(args ...interface{}) {
	str := fmt.Sprintln(args...)
	_ = log.Output(skipFrames+s.skip, str)
}

func (s *stdLogger) outf(msg string, args ...interface{}) {
	str := fmt.Sprintf(msg, args...)
	_ = log.Output(skipFrames+s.skip, str)
}

func (s *stdLogger) Tracef(msg string, args ...interface{}) {
	if s.Level >= TraceLevel {
		s.outf(msg, args...)
	}
}

func (s *stdLogger) Debugf(msg string, args ...interface{}) {
	if s.Level >= DebugLevel {
		s.outf(msg, args...)
	}
}

func (s *stdLogger) Infof(msg string, args ...interface{}) {
	if s.Level >= InfoLevel {
		s.outf(msg, args...)
	}
}

func (s *stdLogger) Warnf(msg string, args ...interface{}) {
	s.outf(msg, args...)
}

func (s *stdLogger) Errorf(msg string, args ...interface{}) {
	s.outf(msg, args...)
}

func (s *stdLogger) Fatalf(msg string, args ...interface{}) {
	s.outf(msg, args...)
	if InTesting() {
		panic(fmt.Sprintf(msg, args...))
	}
	os.Exit(1)
}

func (s *stdLogger) Panicf(msg string, args ...interface{}) {
	str := fmt.Sprintf(msg, args...)
	_ = log.Output(skipFrames+s.skip, str)
	panic(str)
}

func (s *stdLogger) Printf(msg string, args ...interface{}) {
	s.outf(msg, args...)
}

func (s *stdLogger) SetLevel(lvl Level)      { s.Level = lvl }
func (s *stdLogger) GetLevel() Level         { return s.Level }
func (s *stdLogger) SetOutput(out io.Writer) {}
func (s *stdLogger) Setup()                  {}
