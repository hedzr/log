// Copyright © 2020 Hedzr Yeh.

// Package log provide the standard interface of logging for what any go
// libraries want strip off the direct dependency from a known logging
// library.
package log

type (
	// Logger is a minimal logger with no more dependencies
	Logger interface {
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

		// SetLevel sets the logging level
		SetLevel(lvl Level)
		// GetLevel returns the current logging level
		GetLevel() Level

		// Setup will be invoked once an instance created
		Setup()

		// AsFieldLogger() FieldLogger
	}

	// LoggerConfig is used for creating a minimal logger with no more dependencies
	LoggerConfig struct {
		Enabled   bool
		Backend   string // zap, sugar, logrus
		Level     string
		Format    string // text, json, ...
		Target    string // console, file, console+file
		Directory string
		DebugMode bool `json:"-" yaml:"-"`
		TraceMode bool `json:"-" yaml:"-"`

		// the following options are copied from zap rotator

		// MaxSize is the maximum size in megabytes of the log file before it gets
		// rotated. It defaults to 100 megabytes.
		MaxSize int `json:"maxsize" yaml:"maxsize"`

		// MaxAge is the maximum number of days to retain old log files based on the
		// timestamp encoded in their filename.  Note that a day is defined as 24
		// hours and may not exactly correspond to calendar days due to daylight
		// savings, leap seconds, etc. The default is not to remove old log files
		// based on age.
		MaxAge int `json:"maxage" yaml:"maxage"`

		// MaxBackups is the maximum number of old log files to retain.  The default
		// is to retain all old log files (though MaxAge may still cause them to get
		// deleted.)
		MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

		// LocalTime determines if the time used for formatting the timestamps in
		// backup files is the computer's local time.  The default is to use UTC
		// time.
		LocalTime bool `json:"localtime" yaml:"localtime"`

		// Compress determines if the rotated log files should be compressed
		// using gzip. The default is not to perform compression.
		Compress bool `json:"compress" yaml:"compress"`
	}
)

// NewLoggerConfig returns a default LoggerConfig
func NewLoggerConfig() *LoggerConfig {
	return NewLoggerConfigWith(true, "sugar", "info")
}

// NewLoggerConfigWith returns a default LoggerConfig
func NewLoggerConfigWith(enabled bool, backend, level string) *LoggerConfig {
	return &LoggerConfig{
		Enabled:   enabled,
		Backend:   backend,
		Level:     level,
		Format:    "text",
		Target:    "console",
		Directory: "/var/log",
		DebugMode: GetDebugMode(),
		TraceMode: GetTraceMode(),

		MaxSize:    1024, // megabytes
		MaxBackups: 3,    // 3 backups kept at most
		MaxAge:     7,    // 7 days kept at most
		Compress:   true, // disabled by default
	}
}

// SetLevel sets the logging level
func SetLevel(l Level) { logger.SetLevel(l) }

// GetLevel returns the current logging level
func GetLevel() Level { return logger.GetLevel() }

var logger = NewStdLogger()

// SetLogger transfer an instance into log package-level value
func SetLogger(l Logger) { l.SetLevel(logger.GetLevel()); logger = l }

// GetLogger returns the package-level logger globally
func GetLogger() Logger { return logger }

// Tracef prints the text to stdin if logging level is greater than TraceLevel
func Tracef(msg string, args ...interface{}) {
	logger.Printf(msg, args...)
}

// Debugf prints the text to stdin if logging level is greater than DebugLevel
func Debugf(msg string, args ...interface{}) {
	logger.Printf(msg, args...)
}

// Infof prints the text to stdin if logging level is greater than InfoLevel
func Infof(msg string, args ...interface{}) {
	logger.Printf(msg, args...)
}

// Warnf prints the text to stderr
func Warnf(msg string, args ...interface{}) {
	logger.Printf(msg, args...)
}

// Errorf prints the text to stderr
func Errorf(msg string, args ...interface{}) {
	logger.Printf(msg, args...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(msg string, args ...interface{}) {
	logger.Panicf(msg, args...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(msg string, args ...interface{}) {
	logger.Printf(msg, args...)
}
