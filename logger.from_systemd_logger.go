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
		old: logger,
	}
	return l
}

type toSystemdLogger struct {
	lvl Level
	w   io.Writer
	sl  SystemdLogger
	old Logger
}

func (d *toSystemdLogger) With(key string, val interface{}) Logger {
	return d
}

func (d *toSystemdLogger) Trace(args ...interface{}) {
	if d.lvl >= TraceLevel {
		_ = d.sl.Info(args...)
		if d.old != nil {
			AsL(d.old).Trace(args...)
		}
	}
}
func (d *toSystemdLogger) Debug(args ...interface{}) {
	if d.lvl >= DebugLevel {
		_ = d.sl.Info(args...)
		if d.old != nil {
			AsL(d.old).Debug(args...)
		}
	}
}
func (d *toSystemdLogger) Info(args ...interface{}) {
	_ = d.sl.Info(args...)
	if d.old != nil {
		AsL(d.old).Info(args...)
	}
}
func (d *toSystemdLogger) Warn(args ...interface{}) {
	_ = d.sl.Warning(args...)
	if d.old != nil {
		AsL(d.old).Warn(args...)
	}
}
func (d *toSystemdLogger) Error(args ...interface{}) {
	_ = d.sl.Error(args...)
	if d.old != nil {
		AsL(d.old).Error(args...)
	}
}

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
func (d *toSystemdLogger) Print(args ...interface{}) {
	_ = d.sl.Info(args...)
	if d.old != nil {
		AsL(d.old).Info(args...)
	}
}
func (d *toSystemdLogger) Println(args ...interface{}) {
	_ = d.sl.Info(args...)
	if d.old != nil {
		AsL(d.old).Info(args...)
	}
}
func (d *toSystemdLogger) Tracef(msg string, args ...interface{}) {
	if d.lvl >= TraceLevel {
		_ = d.sl.Infof(msg, args...)
		if d.old != nil {
			d.old.Tracef(msg, args...)
		}
	}
}
func (d *toSystemdLogger) Debugf(msg string, args ...interface{}) {
	if d.lvl >= DebugLevel {
		_ = d.sl.Infof(msg, args...)
		if d.old != nil {
			d.old.Debugf(msg, args...)
		}
	}
}
func (d *toSystemdLogger) Infof(msg string, args ...interface{}) {
	_ = d.sl.Infof(msg, args...)
	if d.old != nil {
		d.old.Infof(msg, args...)
	}
}
func (d *toSystemdLogger) Warnf(msg string, args ...interface{}) {
	_ = d.sl.Warningf(msg, args...)
	if d.old != nil {
		d.old.Warnf(msg, args...)
	}
}
func (d *toSystemdLogger) Errorf(msg string, args ...interface{}) {
	_ = d.sl.Errorf(msg, args...)
	if d.old != nil {
		d.old.Errorf(msg, args...)
	}
}

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
