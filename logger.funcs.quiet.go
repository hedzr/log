//go:build veryquiet
// +build veryquiet

package log

// VeryQuietEnabled identify whether `--tags=veryquiet` has been defined in go building
var VeryQuietEnabled = true

// Tracef prints the text to stdin if logging level is greater than TraceLevel
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Tracef(msg string, args ...interface{}) {
	// logger.Tracef(msg, args...)
}

// Debugf prints the text to stdin if logging level is greater than DebugLevel
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Debugf(msg string, args ...interface{}) {
	// logger.Debugf(msg, args...)
}

// Infof prints the text to stdin if logging level is greater than InfoLevel
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Infof(msg string, args ...interface{}) {
	// logger.Infof(msg, args...)
}

// Warnf prints the text to stderr
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Warnf(msg string, args ...interface{}) {
	// logger.Warnf(msg, args...)
}

// Errorf prints the text to stderr
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Errorf(msg string, args ...interface{}) {
	// logger.Errorf(msg, args...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Fatalf(msg string, args ...interface{}) {
	// if InTesting() {
	//	logger.Panicf(msg, args)
	// }
	// logger.Fatalf(msg, args...)
}

// Panicf is equivalent to Printf() followed by a call to panic().
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Panicf(msg string, args ...interface{}) {
	// logger.Panicf(msg, args...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Printf(msg string, args ...interface{}) {
	// logger.Printf(msg, args...)
}

// Trace prints all args to stdin if logging level is greater than TraceLevel
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Trace(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Trace(args...)
	// }
}

// Debug prints all args to stdin if logging level is greater than DebugLevel
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Debug(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Debug(args...)
	// }
}

// Info prints all args to stdin if logging level is greater than InfoLevel
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Info(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Info(args...)
	// }
}

// Warn prints all args to stderr
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Warn(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Warn(args...)
	// }
}

// Error prints all args to stderr
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Error(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Error(args...)
	// }
}

// Fatal is equivalent to Printf() followed by a call to os.Exit(1).
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Fatal(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	if InTesting() {
	//		l.Panic(args)
	//	}
	//	l.Fatal(args...)
	// }
}

// Panic is equivalent to Printf() followed by a call to panic().
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Panic(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Panic(args...)
	// }
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Print(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Print(args...)
	// }
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
// It would be optimized to discard if `--tags=veryquiet` was been defined.
func Println(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Println(args...)
	// }
}
