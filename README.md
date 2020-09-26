# log


<!-- ![Build Status](https://travis-ci.org/hedzr/log.svg?branch=master)](https://travis-ci.org/hedzr/log) -->
![Go](https://github.com/hedzr/log/workflows/Go/badge.svg)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/hedzr/log.svg?label=release)](https://github.com/hedzr/log/releases)
[![Sourcegraph](https://sourcegraph.com/github.com/hedzr/log/-/badge.svg)](https://sourcegraph.com/github.com/hedzr/log?badge)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/hedzr/log)
[![go.dev](https://img.shields.io/badge/go.dev-reference-green)](https://pkg.go.dev/github.com/hedzr/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedzr/log)](https://goreportcard.com/report/github.com/hedzr/log)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fhedzr%2Flog.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fhedzr%2Flog?ref=badge_shield)
<!-- [![codecov](https://codecov.io/gh/hedzr/log/branch/master/graph/badge.svg)](https://codecov.io/gh/hedzr/log) -->
[![Coverage Status](https://coveralls.io/repos/github/hedzr/log/badge.svg)](https://coveralls.io/github/hedzr/log)



## Common Interfaces for logging

Here:

```go
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
```



## Functions

### Package-level functions

Panicf, Fatalf, Errorf, Warnf, Infof, Debugf, Tracef

### Dummy and Standard Logger

See the following codes:

```go
import "github.com/hedzr/log"

var dummy, std log.Logger
func a(){
  // dummy Logger will discard any logging outputs
  dummy = log.NewDummyLogger()
  std = log.NewStdLoggerWith(log.OffLevel) // OffLevel is liks Dummy Logger
  std = log.NewStdLogger()
  
  std.Infof("slsl")
  
  print(std.GetLevel())
}
```

### Utilities

#### Directory Helper

in `exec/dir.go`,

```go
import "github.com/hedzr/log/exec"

func a(){
  print(exec.IsDiretory(exec.GetCurrentDir()))
  print(exec.GetExecutablePath())
  print(exec.GetExecutableDir())
}
```

```go

IsRegularFile
FileExists
EnsureDir
EnsureDirEnh
NormalizeDir

```

```go
_ = exec.ForDirMax(dir, 0, 1, func(depth int, cwd string, fi os.FileInfo) (stop bool, err error) {
	if fi.IsDir() {
		return
	}
      // ... doing something for a file,
	return
})

```



#### Exec Helpers

```go
import "github.com/hedzr/log/exec"

exec.Run()
exec.Sudo()
exec.RunWithOutput()
exec.RunCommand()
exec.IsExitError()
exec.IsEAccess()
```



#### Trace Helpers

```go
import "github.com/hedzr/log/trace"

trace.RegisterOnTraceModeChanges(handler)
trace.IsEnable()

trace.Start()
trace.Stop()
```







## LICENSE

MIT
