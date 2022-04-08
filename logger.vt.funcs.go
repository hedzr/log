//go:build verbose
// +build verbose

package log

// VerboseEnabled identify whether `--tags=verbose` has been defined in go building
var VerboseEnabled = true

// VTracef prints the text to stdin if logging level is greater than TraceLevel
func VTracef(msg string, args ...interface{}) {
	logger.Tracef(msg, args...)
}

// VDebugf prints the text to stdin if logging level is greater than DebugLevel
func VDebugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

// VInfof prints the text to stdin if logging level is greater than InfoLevel
func VInfof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

// VWarnf prints the text to stderr
func VWarnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

// VErrorf prints the text to stderr
func VErrorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

// VFatalf is equivalent to Printf() followed by a call to os.Exit(1).
func VFatalf(msg string, args ...interface{}) {
	if InTesting() {
		logger.Panicf(msg, args)
	}
	logger.Fatalf(msg, args...)
}

// VPanicf is equivalent to Printf() followed by a call to panic().
func VPanicf(msg string, args ...interface{}) {
	logger.Panicf(msg, args...)
}

// VPrintf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func VPrintf(msg string, args ...interface{}) {
	logger.Printf(msg, args...)
}

// VTrace prints all args to stdin if logging level is greater than TraceLevel
func VTrace(args ...interface{}) {
	if l := AsL(logger); l != nil {
		l.Trace(args...)
	}
}

// VDebug prints all args to stdin if logging level is greater than DebugLevel
func VDebug(args ...interface{}) {
	if l := AsL(logger); l != nil {
		l.Debug(args...)
	}
}

// VInfo prints all args to stdin if logging level is greater than InfoLevel
func VInfo(args ...interface{}) {
	if l := AsL(logger); l != nil {
		l.Info(args...)
	}
}

// VWarn prints all args to stderr
func VWarn(args ...interface{}) {
	if l := AsL(logger); l != nil {
		l.Warn(args...)
	}
}

// VError prints all args to stderr
func VError(args ...interface{}) {
	if l := AsL(logger); l != nil {
		l.Error(args...)
	}
}

// VFatal is equivalent to Printf() followed by a call to os.Exit(1).
func VFatal(args ...interface{}) {
	if l := AsL(logger); l != nil {
		if InTesting() {
			l.Panic(args)
		}
		l.Fatal(args...)
	}
}

// VPanic is equivalent to Printf() followed by a call to panic().
func VPanic(args ...interface{}) {
	if l := AsL(logger); l != nil {
		l.Panic(args...)
	}
}

// VPrint calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func VPrint(args ...interface{}) {
	if l := AsL(logger); l != nil {
		l.Print(args...)
	}
}

// VPrintln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func VPrintln(args ...interface{}) {
	if l := AsL(logger); l != nil {
		l.Println(args...)
	}
}
