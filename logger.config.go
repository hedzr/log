package log

import (
	"github.com/hedzr/log/states"
)

type (
	// LoggerConfig is used for creating a minimal logger with no more dependencies
	LoggerConfig struct {
		Enabled          bool
		Backend          string // zap, sugar, logrus
		Level            string // level
		Format           string // text, json, ...
		Target           string // console, file, console+file
		Directory        string // logdir, for file
		AllToErrorDevice bool   //
		DebugMode        bool   `json:"-" yaml:"-"`
		TraceMode        bool   `json:"-" yaml:"-"`

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

		ExtraSkip       int
		ShortTimestamp  bool   // remove year field for a shorter timestamp stringify
		TimestampFormat string // never used
	}
)

// NewLoggerConfig returns a default LoggerConfig
func NewLoggerConfig(opts ...Opt) *LoggerConfig {
	return NewLoggerConfigWith(true, "sugar", "info", opts...)
}

// NewLoggerConfigWith returns a default LoggerConfig
func NewLoggerConfigWith(enabled bool, backend, level string, opts ...Opt) *LoggerConfig {
	var dm, tm = states.Env().GetDebugMode(), states.Env().GetTraceMode()
	if dm {
		level = "debug"
	}
	if tm {
		level = "trace"
	}

	var l Level
	l, _ = ParseLevel(level)
	states.Env().SetDebugMode(l >= DebugLevel)
	states.Env().SetTraceMode(l >= TraceLevel)

	dm, tm = states.Env().GetDebugMode(), states.Env().GetTraceMode()

	lc := &LoggerConfig{
		Enabled:   enabled,
		Backend:   backend,
		Level:     level,
		Format:    "text",
		Target:    "console",
		Directory: "/var/log",
		DebugMode: dm,
		TraceMode: tm,

		MaxSize:    1024, // megabytes
		MaxBackups: 3,    // 3 backups kept at most
		MaxAge:     7,    // 7 days kept at most
		Compress:   true, // disabled by default

		ExtraSkip:       -1,
		ShortTimestamp:  false, //
		TimestampFormat: "",
	}

	for _, opt := range opts {
		opt(lc)
	}
	return lc
}

// Opt _
type Opt func(lc *LoggerConfig)

// WithTimestamp _
func WithTimestamp(short bool, format ...string) Opt {
	return func(lc *LoggerConfig) {
		lc.ShortTimestamp = short
		for _, fmt := range format {
			lc.TimestampFormat = fmt
		}
	}
}

// WithExtraSkip _
func WithExtraSkip(extraSkip int) Opt {
	return func(lc *LoggerConfig) {
		lc.ExtraSkip = extraSkip
	}
}
