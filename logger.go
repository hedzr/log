// Copyright © 2020 Hedzr Yeh.

// Package log provide the standard interface of logging for what any go
// libraries want strip off the direct dependency from a known logging
// library.
package log

import (
	"io"
	"os"
	"strings"
)

type (

	// SystemdLogger writes to the system log.
	SystemdLogger interface {
		Error(v ...interface{}) error
		Warning(v ...interface{}) error
		Info(v ...interface{}) error

		Errorf(format string, a ...interface{}) error
		Warningf(format string, a ...interface{}) error
		Infof(format string, a ...interface{}) error
	}

	// SL provides a structural logging interface
	SL interface {
		With(key string, val interface{}) Logger
	}

	// L provides a basic logger interface
	L interface {

		// Trace prints all args to stdin if logging level is greater than TraceLevel
		Trace(args ...interface{})
		// Debug prints all args to stdin if logging level is greater than DebugLevel
		Debug(args ...interface{})
		// Info prints all args to stdin if logging level is greater than InfoLevel
		Info(args ...interface{})
		// Warn prints all args to stderr
		Warn(args ...interface{})
		// Error prints all args to stderr
		Error(args ...interface{})
		// Fatal is equivalent to Printf() followed by a call to os.Exit(1).
		Fatal(args ...interface{})
		// Panic is equivalent to Printf() followed by a call to panic().
		Panic(args ...interface{})
		// Print calls Output to print to the standard logger.
		// Arguments are handled in the manner of fmt.Print.
		Print(args ...interface{})
		// Println calls Output to print to the standard logger.
		// Arguments are handled in the manner of fmt.Println.
		Println(args ...interface{})
	}

	// LF provides a L logger interface and format prototypes (such as Debugf...)
	LF interface {
		SL

		// Tracef prints the text to stdin if logging level is greater than TraceLevel
		Tracef(msg string, args ...interface{})
		// Debugf prints the text to stdin if logging level is greater than DebugLevel
		Debugf(msg string, args ...interface{})
		// Infof prints the text to stdin if logging level is greater than InfoLevel
		Infof(msg string, args ...interface{})
		// Warnf prints the text to stderr
		Warnf(msg string, args ...interface{})
		// Errorf prints the text to stderr
		Errorf(msg string, args ...interface{})
		// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
		Fatalf(msg string, args ...interface{})
		// Panicf is equivalent to Printf() followed by a call to panic().
		Panicf(msg string, args ...interface{})
		// Printf calls Output to print to the standard logger.
		// Arguments are handled in the manner of fmt.Printf.
		Printf(msg string, args ...interface{})
	}

	// Logger is a minimal logger with no more dependencies
	Logger interface {
		LF

		// SetLevel sets the logging level
		SetLevel(lvl Level)
		// GetLevel returns the current logging level
		GetLevel() Level
		// SetOutput setup the logging output device
		SetOutput(out io.Writer)
		// GetOutput returns the current logging output device
		GetOutput() (out io.Writer)

		// Setup will be invoked once an instance created
		Setup()

		// AddSkip adds an extra count to skip stack frames
		AddSkip(skip int) Logger

		// AsFieldLogger() FieldLogger
	}

	// LoggerExt is a minimal logger with no more dependencies
	LoggerExt interface {
		L
		Logger
	}

	// BuilderFunc provides a function prototype for creating a hedzr/log & hedzr/logex -compliant creator.
	BuilderFunc func(config *LoggerConfig) (logger Logger)
)

// InTesting detects whether is running under go test mode
func InTesting() bool { return InTestingT(os.Args) }

// InTestingT detects whether is running under go test mode
func InTestingT(args []string) bool {
	if !strings.HasSuffix(args[0], ".test") &&
		!strings.Contains(args[0], "/T/___Test") {

		// [0] = /var/folders/td/2475l44j4n3dcjhqbmf3p5l40000gq/T/go-build328292371/b001/exe/main
		// !strings.Contains(SavedOsArgs[0], "/T/go-build")

		for _, s := range args {
			if s == "-test.v" || s == "-test.run" {
				return true
			}
		}
		return false

	}
	return true
}

// AsL converts a logger to L type (with Info(...), ... prototypes)
func AsL(logger LF) L {
	if l, ok := logger.(L); ok {
		//if l1, ok := l.(Logger); ok {
		//	return l1.AddSkip(1).(L)
		//}
		return l
	}
	return nil
}

// AsLogger converts a logger to LF or Logger type (with Infof(...), ... prototypes)
func AsLogger(logger L) Logger {
	if l, ok := logger.(Logger); ok {
		return l // .AddSkip(1)
	}
	return nil
}

// SetLevel sets the logging level
func SetLevel(l Level) { logger.SetLevel(l) }

// Setup _
func Setup() {
	logger.Setup()
}

// GetLevel returns the current logging level
func GetLevel() Level { return logger.GetLevel() }

// SetOutput setup the logging output device
func SetOutput(w io.Writer) { logger.SetOutput(w) }

// GetOutput return the logging output device
func GetOutput() (w io.Writer) { return logger.GetOutput() }

// SetLogger transfer an instance into log package-level value
func SetLogger(l Logger) { l.SetLevel(logger.GetLevel()); logger = l }

// GetLogger returns the package-level logger globally
func GetLogger() Logger { return logger }

// Skip ignore some extra caller frames
func Skip(skip int) Logger {
	return logger.AddSkip(skip)
	//return logger
}
