//go:build !verbose
// +build !verbose

package log

// VerboseEnabled identify whether `--tags=verbose` has been defined in go building
var VerboseEnabled = false

// VTracef prints the text to stdin if logging level is greater than TraceLevel.
// It would be optimized to discard except `--tags=verbose` was been defined.
func VTracef(msg string, args ...interface{}) {
	// logger.Tracef(msg, args...)
}

// VDebugf prints the text to stdin if logging level is greater than DebugLevel
// It would be optimized to discard except `--tags=verbose` was been defined.
func VDebugf(msg string, args ...interface{}) {
	// logger.Debugf(msg, args...)
}

// VInfof prints the text to stdin if logging level is greater than InfoLevel
// It would be optimized to discard except `--tags=verbose` was been defined.
func VInfof(msg string, args ...interface{}) {
	// logger.Infof(msg, args...)
}

// VWarnf prints the text to stderr
// It would be optimized to discard except `--tags=verbose` was been defined.
func VWarnf(msg string, args ...interface{}) {
	// logger.Warnf(msg, args...)
}

// VErrorf prints the text to stderr
// It would be optimized to discard except `--tags=verbose` was been defined.
func VErrorf(msg string, args ...interface{}) {
	// logger.Errorf(msg, args...)
}

// VFatalf is equivalent to Printf() followed by a call to os.Exit(1).
// It would be optimized to discard except `--tags=verbose` was been defined.
func VFatalf(msg string, args ...interface{}) {
	// if InTesting() {
	//	logger.Panicf(msg, args)
	// }
	// logger.Fatalf(msg, args...)
}

// VPanicf is equivalent to Printf() followed by a call to panic().
// It would be optimized to discard except `--tags=verbose` was been defined.
func VPanicf(msg string, args ...interface{}) {
	// logger.Panicf(msg, args...)
}

// VPrintf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
// It would be optimized to discard except `--tags=verbose` was been defined.
func VPrintf(msg string, args ...interface{}) {
	// logger.Printf(msg, args...)
}

// VTrace prints all args to stdin if logging level is greater than TraceLevel
// It would be optimized to discard except `--tags=verbose` was been defined.
func VTrace(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Trace(args...)
	// }
}

// VDebug prints all args to stdin if logging level is greater than DebugLevel
// It would be optimized to discard except `--tags=verbose` was been defined.
func VDebug(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Debug(args...)
	// }
}

// VInfo prints all args to stdin if logging level is greater than InfoLevel
// It would be optimized to discard except `--tags=verbose` was been defined.
func VInfo(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Info(args...)
	// }
}

// VWarn prints all args to stderr
// It would be optimized to discard except `--tags=verbose` was been defined.
func VWarn(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Warn(args...)
	// }
}

// VError prints all args to stderr
// It would be optimized to discard except `--tags=verbose` was been defined.
func VError(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Error(args...)
	// }
}

// VFatal is equivalent to Printf() followed by a call to os.Exit(1).
// It would be optimized to discard except `--tags=verbose` was been defined.
func VFatal(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	if InTesting() {
	//		l.Panic(args)
	//	}
	//	l.Fatal(args...)
	// }
}

// VPanic is equivalent to Printf() followed by a call to panic().
// It would be optimized to discard except `--tags=verbose` was been defined.
func VPanic(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Panic(args...)
	// }
}

// VPrint calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
// It would be optimized to discard except `--tags=verbose` was been defined.
func VPrint(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Print(args...)
	// }
}

// VPrintln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
// It would be optimized to discard except `--tags=verbose` was been defined.
func VPrintln(args ...interface{}) {
	// if l := AsL(logger); l != nil {
	//	l.Println(args...)
	// }
}
