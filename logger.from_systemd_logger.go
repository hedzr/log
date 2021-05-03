package log

import (
	"fmt"
	"io"
	"os"
)

// FromSystemdLogger converts a SystemdLogger to Logger so that you can put it into `log` system via log.SetLogger.
func FromSystemdLogger(sl SystemdLogger) Logger {
	l := &toSystemdLogger{
		lvl: logger.GetLevel(),
		w:   nil,
		sl:  sl,
	}
	return l
}

type toSystemdLogger struct {
	lvl Level
	w   io.Writer
	sl  SystemdLogger
}

func (d *toSystemdLogger) Trace(args ...interface{}) {
	if d.lvl >= TraceLevel {
		_ = d.sl.Info(args...)
	}
}
func (d *toSystemdLogger) Debug(args ...interface{}) {
	if d.lvl >= DebugLevel {
		_ = d.sl.Info(args...)
	}
}
func (d *toSystemdLogger) Info(args ...interface{})  { _ = d.sl.Info(args...) }
func (d *toSystemdLogger) Warn(args ...interface{})  { _ = d.sl.Warning(args...) }
func (d *toSystemdLogger) Error(args ...interface{}) { _ = d.sl.Error(args...) }

func (d *toSystemdLogger) Fatal(args ...interface{}) {
	d.Error(args...)
	if InTesting() {
		panic(fmt.Sprint(args...))
	}
	os.Exit(1)
}

func (d *toSystemdLogger) Panic(args ...interface{}) {
	d.Error(args...)
	panic(fmt.Sprint(args...))
}
func (d *toSystemdLogger) Print(args ...interface{})   { _ = d.sl.Info(args...) }
func (d *toSystemdLogger) Println(args ...interface{}) { _ = d.sl.Info(args...) }
func (d *toSystemdLogger) Tracef(msg string, args ...interface{}) {
	if d.lvl >= TraceLevel {
		_ = d.sl.Infof(msg, args...)
	}
}
func (d *toSystemdLogger) Debugf(msg string, args ...interface{}) {
	if d.lvl >= DebugLevel {
		_ = d.sl.Infof(msg, args...)
	}
}
func (d *toSystemdLogger) Infof(msg string, args ...interface{})  { _ = d.sl.Infof(msg, args...) }
func (d *toSystemdLogger) Warnf(msg string, args ...interface{})  { _ = d.sl.Warningf(msg, args...) }
func (d *toSystemdLogger) Errorf(msg string, args ...interface{}) { _ = d.sl.Errorf(msg, args...) }

func (d *toSystemdLogger) Fatalf(msg string, args ...interface{}) {
	// panic("implement me")
	d.Errorf(msg, args...)
	if InTesting() {
		panic(fmt.Sprintf(msg, args...))
	}
	os.Exit(1)
}

func (d *toSystemdLogger) Panicf(msg string, args ...interface{}) {
	d.Errorf(msg, args...)
	panic(fmt.Sprintf(msg, args...))
}
func (d *toSystemdLogger) Printf(msg string, args ...interface{}) {
	if d.w != nil {
		str := fmt.Sprintf(msg, args...)
		_, _ = d.w.Write([]byte(str))
		return
	}
	d.Infof(msg, args...)
}
func (d *toSystemdLogger) SetLevel(lvl Level)         { d.lvl = lvl }
func (d *toSystemdLogger) GetLevel() Level            { return d.lvl }
func (d *toSystemdLogger) SetOutput(out io.Writer)    { d.w = out }
func (d *toSystemdLogger) GetOutput() (out io.Writer) { return d.w }
func (d *toSystemdLogger) Setup()                     {}
func (d *toSystemdLogger) AddSkip(skip int) Logger    { return d }
